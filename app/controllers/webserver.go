package controllers

import (
	"fmt"
	"gotrading/gotrading/app/models"
	"gotrading/gotrading/config"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("app/views/googlechart.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	limit := 100
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)
	err := templates.ExecuteTemplate(w, "googlechart.html", df.Candles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
