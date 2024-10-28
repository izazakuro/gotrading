package controllers

import (
	"gotrading/gotrading/app/models"
	"gotrading/gotrading/bitflyer"
	"gotrading/gotrading/config"
	"log"
)

func StreamIngestionData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		log.Printf("action=StreamIngestionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated && duration == config.Config.TradeDuration {
				// TODO
			}
		}
	}

}
