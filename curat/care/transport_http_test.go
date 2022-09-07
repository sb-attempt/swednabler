package curat

import (
	"github.com/go-kit/log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHTTPPost(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8080", "caller", log.DefaultCaller)
	s := NewTerminologyService()
	r := NewHttpServer(s, logger)
	srv := httptest.NewServer(r)

	for _, testcase := range []struct {
		method, url, body string
		want              int
	}{
		{"POST", srv.URL + "/v1/term/simplify", `{"id": 0}`, http.StatusUnauthorized},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s %s: want %d have %d", testcase.method, testcase.url, testcase.body, testcase.want, resp.StatusCode)
		}

	}
}

func TestHTTPGet(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8080", "caller", log.DefaultCaller)
	s := NewTerminologyService()
	r := NewHttpServer(s, logger)
	srv := httptest.NewServer(r)

	for _, testcase := range []struct {
		method, url string
		want        int
	}{
		{"GET", srv.URL + "/v1/term/list", http.StatusUnauthorized},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(""))
		resp, _ := http.DefaultClient.Do(req)
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s : want %d have %d", testcase.method, testcase.url, testcase.want, resp.StatusCode)
		}

	}
}
