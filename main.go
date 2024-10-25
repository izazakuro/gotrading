package main

import (
	"fmt"
	"gotrading/gotrading/bitflyer"
	"gotrading/gotrading/config"
	"gotrading/gotrading/utils"
	"time"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_JPY")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
