package main

import (
	"fmt"
	config "golang/platform/geth"
	"golang/platform/geth/client"
)

func main() {
	config.LoadConfig()
	infuraClient := client.NewClient(
		client.GetInfuraURL(
			config.GetConfig("NETWORK"),
			config.GetConfig("INFURA_API_KEY"),
		),
	)

	fmt.Println(infuraClient.BalanceAt(
		config.GetConfig("ACCOUNT_ADDRESS"),
	))
	fmt.Println(infuraClient.PendingNonceAt(
		config.GetConfig("ACCOUNT_ADDRESS"),
	))
}
