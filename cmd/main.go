package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wboayue/ibapi"
)

func main() {
	client, err := ibapi.Connect("localhost", 4002, 100)
	if err != nil {
		log.Printf("error connecting: %v", err)
		return
	}

	defer client.Close()

	fmt.Printf("server version: %v\n", client.ServerVersion)
	fmt.Printf("server time: %v\n", client.ServerTime)

	go client.ProcessMessages()

	ctx := context.Background()

	// realTimeBars(ctx, client)
	contractDetails(ctx, client)
	//	tickByTickTrades(ctx, client)
	// tickByTickSpreads(ctx, client)
}

func realTimeBars(ctx context.Context, client *ibapi.IbClient) {
	contract := ibapi.Contract{
		LocalSymbol: "ESH2",
		// LocalSymbol:  "6EF2",
		SecurityType: "FUT",
		Currency:     "USD",
		Exchange:     "GLOBEX",
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	bars, err := client.RealTimeBars(ctx, contract, "TRADES", false)
	if err != nil {
		log.Printf("error connecting: %v", err)
		return
	}

	for bar := range bars {
		fmt.Println(bar)
	}
}

func tickByTickTrades(ctx context.Context, client *ibapi.IbClient) {
	contract := ibapi.Contract{
		LocalSymbol: "ESH2",
		// LocalSymbol:  "CLG2",
		SecurityType: "FUT",
		Currency:     "USD",
		Exchange:     "GLOBEX",
		// Exchange: "NYMEX",
	}

	trades, err := client.TickByTickTrades(ctx, contract)
	if err != nil {
		log.Printf("error connecting: %v", err)
		return
	}

	for trade := range trades {
		fmt.Printf("trade: %+v\n", trade)
	}
}

func contractDetails(ctx context.Context, client *ibapi.IbClient) {
	contract := ibapi.Contract{
		Symbol:                       "ES",
		SecurityType:                 "FUT",
		Currency:                     "USD",
		LastTradeDateOrContractMonth: "2022",
	}

	contracts, err := client.ContractDetails(ctx, contract)
	if err != nil {
		log.Printf("error connecting: %v", err)
		return
	}

	for i, contract := range contracts {
		fmt.Printf("%d - %+v\n", i, contract)
	}
}

func tickByTickSpreads(ctx context.Context, client *ibapi.IbClient) {
	contract := ibapi.Contract{
		LocalSymbol: "ESH2",
		// LocalSymbol:  "CLG2",
		SecurityType: "FUT",
		Currency:     "USD",
		Exchange:     "GLOBEX",
		// Exchange: "NYMEX",
	}

	spreads, err := client.TickByTickBidAsk(ctx, contract)
	if err != nil {
		log.Printf("error connecting: %v", err)
		return
	}

	for spread := range spreads {
		fmt.Printf("bid/ask: %+v\n", spread)
	}
}
