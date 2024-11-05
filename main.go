package main

import (
	"fmt"
	"gotrading/gotrading/app/models"
	"time"
)

func main() {
	//utils.LoggingSetting(config.Config.LogFile)
	/*
		apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

		order := &bitflyer.Order{
			ProductCode:    config.Config.ProductCode,
			ChildOrderType: "MARKET", // "LIMIT"
			Side:           "BUY",    // or SELL
			Size:           0.01,
			//Price : (Price to BUY or SELL)
			MinuteToExpires: 1,
			TimeInForce:     "GTC",
		}
		res, _ := apiClient.SendOrder(order)
		fmt.Println(res.ChildOdrderAcceptanceID)

		orderID := "XXXXXXXXXXXXXXXXXXXXXXXXXX"
		params := map[string]string{
			"product_code":              config.Config.ProductCode,
			"child_order_acceptance_id": orderID,
		}
		r, _ := apiClient.ListOrder(params)
		fmt.Println(r)
	*/
	//fmt.Println(models.DbConnection)

	s := models.NewSingalEvents()
	df, _ := models.GetAllCandle("BTC_JPY", time.Minute, 10)
	c1 := df.Candles[0]
	c2 := df.Candles[3]
	s.Buy("BTC_JPY", c1.Time.Local().UTC(), c1.Close, 1.0, true)
	s.Sell("BTC_JPY", c2.Time.Local().UTC(), c2.Close, 1.0, true)
	fmt.Println(models.GetSignalEventsByCount(1))
	fmt.Println(models.GetSignalEventsAfterTime(c1.Time))
	fmt.Println(s.CollectAfter(time.Now().Local().UTC()))
	fmt.Println(s.CollectAfter(c1.Time))

}
