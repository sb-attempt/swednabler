package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"net/http"
	"os"
	curat "swednabler/curat/v2/care"
)

const (
	filePath = "./data/database.json"
)

func init() {
	loadConfiguration(filePath)
}

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8082", "caller", log.DefaultCaller)
	// the http server that listens on 8082 and pass the incoming request to the router
	r := curat.NewHttpServer(curat.NewTerminologyService(), logger)
	_ = logger.Log("msg", "HTTP", "addr", "8082")
	_ = logger.Log("err", http.ListenAndServe(":8082", r))
}

func loadConfiguration(filePath string) {
	var logger = log.NewLogfmtLogger(os.Stderr)
	_ = logger.Log("Loading database...")
	configFile, err := os.Open(filePath)
	defer func() {
		cerr := configFile.Close()
		if err == nil {
			err = cerr
		}
	}()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&curat.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
