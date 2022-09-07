package open

import (
	"github.com/go-kit/log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHTTP(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8080", "caller", log.DefaultCaller)
	s := NewService()
	r := NewHttpServer(s, NewJwtService(), logger)
	srv := httptest.NewServer(r)

	for _, testcase := range []struct {
		method, url, body string
		want              int
	}{
		{"POST", srv.URL + "/v1/token", `{"username":"user10","password":"password10"}`, http.StatusOK},
		{"POST", srv.URL + "/v1/token/validate", `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0"}`, http.StatusOK},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s %s: want %d have %d", testcase.method, testcase.url, testcase.body, testcase.want, resp.StatusCode)
		}

	}
}
