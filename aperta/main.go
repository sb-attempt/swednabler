package main

import (
	"github.com/go-kit/log"
	"net/http"
	"os"
	"swednabler/aperta/v2/open"
)

// The main method
func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8081", "caller", log.DefaultCaller)
	// the http server that listens on 8081 and pass the incoming request to the router
	r := open.NewHttpServer(open.NewService(), open.NewJwtService(), logger)
	_ = logger.Log("msg", "HTTP", "addr", "8081")
	_ = logger.Log("err", http.ListenAndServe(":8081", r))
}
