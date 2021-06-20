package api

import (
	"github.com/atominkiss/YamlOpenMetricsConverter/internal"
	"html/template"
	"log"
	"net/http"
)

// Парсим шаблон
var tpl = template.Must(template.ParseFiles("./web/simpleFront.html"))

// GetMetricsHandler pushing data to frontend template
func GetMetricsHandler(w http.ResponseWriter, _ *http.Request) {
	// Устанавливаем header как указано в документации OpenMetrics
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Заполняем шаблон данными
	if err := tpl.Execute(w, internal.YamlParser()); err != nil {
		log.Fatal("Проблемы с шаблоном")
	}
}
