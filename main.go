package main

import (
	"net/http"
	"tag-controller/logger"
	"tag-controller/aws_api"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/ins", aws_api.InstanceHandler)

	logger.Info("start Server2")
    http.ListenAndServe(":4040", mux)
}
