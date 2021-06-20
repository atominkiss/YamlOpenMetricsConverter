package main

import (
	"github.com/atominkiss/YamlOpenMetricsConverter/api"
	"log"
	"net/http"
)

func main() {
	// Запускаем сервер и вешаем хендлер который слушает по "http://localhost:8080/metrics"
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", api.GetMetricsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
