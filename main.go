package main

import (
	"fmt"
	"gotrading/gotrading/bitflyer"
	"gotrading/gotrading/config"
	"gotrading/gotrading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
