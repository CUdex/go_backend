package main

import (
	"net/http"
	"tag-controller/logger"
	"tag-controller/aws_api"
	"tag-controller/prom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	
)

func main() {
	const port = "4040"
    mux := http.NewServeMux()
    mux.HandleFunc("/ins", aws_api.TagGetHandler)
	mux.HandleFunc("/tag", aws_api.TagAddHandler)

	//prometheus exporter 생성
	prom.InitPrometheus()
	mux.Handle("/metrics", promhttp.Handler())

	logger.Info("start Server2")
    http.ListenAndServe(":" + port, mux)
}
