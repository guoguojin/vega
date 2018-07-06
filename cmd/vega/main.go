// cmd/vega/main.go
package main

import (
	"vega/api/endpoints/rest"
	"vega/api/endpoints/sse"
	"vega/blockchain"
	"vega/core"
	"vega/datastore"
	"vega/proto"
	"vega/api"
)

const sseChannelSize = 2 << 16
const storeChannelSize = 2 << 16
const marketName = "BTC/DEC18"

func main() {

	config := core.DefaultConfig()

	// Storage Service provides read stores for consumer VEGA API
	// Uses in memory storage (maps/slices etc), configurable in future
	storage := &datastore.MemoryStoreProvider{}
	storage.Init([]string{marketName})

	// Initialise concrete consumer services
	orderService := api.NewOrderService()
	tradeService := api.NewTradeService()
	orderService.Init(storage.OrderStore())
	tradeService.Init(storage.TradeStore())

	// REST server
	restServer := rest.NewRestServer(orderService, tradeService)

	sseOrderChan := make(chan msg.Order, sseChannelSize)
	sseTradeChan := make(chan msg.Trade, sseChannelSize)
	sseServer := sse.NewServer(sseOrderChan, sseTradeChan)

	vega := core.New(config)
	vega.CreateMarket(marketName)

	go restServer.Start()
	go sseServer.Start()
	blockchain.Start(vega)

}
