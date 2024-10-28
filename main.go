package main

import (
	"fmt"
	"gotrading/gotrading/app/models"
	"gotrading/gotrading/config"
	"gotrading/gotrading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
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
	fmt.Println(models.DbConnection)

}
