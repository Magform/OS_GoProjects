package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create channels for market data and trading signals
	marketData := make(chan string)
	tradingSignals := make(chan string)

	// Start simulating market data
	go simulateMarketData(marketData)

	// Start selecting and handling trading operations
	go selectPair(marketData, tradingSignals)

	// Start trading cycle for one minute
	timer := time.NewTimer(time.Minute)

	for {
		select {
		case data := <-marketData:
			fmt.Println("Received market data:", data)
		case signal := <-tradingSignals:
			fmt.Println("Received trading signal:", signal)
		case <-timer.C:
			fmt.Println("Trading cycle completed.")
			return
		}
	}
}

func simulateMarketData(marketData chan<- string) {
	rand.Seed(time.Now().UnixNano())

	for {
		// Simulate EUR/USD price between 1.0 and 1.5
		eurUsdPrice := 1.0 + rand.Float64()*0.5
		marketData <- fmt.Sprintf("EUR/USD: %.4f", eurUsdPrice)

		// Simulate GBP/USD price between 1.0 and 1.5
		gbpUsdPrice := 1.0 + rand.Float64()*0.5
		marketData <- fmt.Sprintf("GBP/USD: %.4f", gbpUsdPrice)

		// Simulate JPY/USD price between 0.006 and 0.009
		jpyUsdPrice := 0.006 + rand.Float64()*0.003
		marketData <- fmt.Sprintf("JPY/USD: %.4f", jpyUsdPrice)

		time.Sleep(time.Second)
	}
}

func selectPair(marketData <-chan string, tradingSignals chan<- string) {
	for {
		select {
		case data := <-marketData:
			switch {
			case data[:7] == "EUR/USD" && getPrice(data) > 1.20:
				time.Sleep(4 * time.Second)
				tradingSignals <- "Sell EUR/USD"
			case data[:7] == "GBP/USD" && getPrice(data) < 1.35:
				time.Sleep(3 * time.Second)
				tradingSignals <- "Buy GBP/USD"
			case data[:7] == "JPY/USD" && getPrice(data) < 0.0085:
				time.Sleep(3 * time.Second)
				tradingSignals <- "Buy JPY/USD"
			}
		}
	}
}

func getPrice(data string) float64 {
	var price float64
	fmt.Sscanf(data, "%*s %f", &price)
	return price
}
