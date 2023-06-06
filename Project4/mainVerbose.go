package main

import (
	"fmt"
	"math/rand"
	"time"
)

var EUR_USD = make(chan float64, 1)
var GBP_USD = make(chan float64, 1)
var JPY_USD = make(chan float64, 1)

func simulateEUR() {
	var EUR_USD_value float64 = rand.Float64()*0.5 + 1
	EUR_USD <- EUR_USD_value
	fmt.Println("EUR/USD: ", EUR_USD_value)
}

func simulateGBP() {
	var GBP_USD_value float64 = rand.Float64()*0.5 + 1
	GBP_USD <- GBP_USD_value
	fmt.Println("GBP/USD: ", GBP_USD_value)
}

func simulateJPY() {
	var JPY_USD_value float64 = rand.Float64()*0.003 + 0.006
	JPY_USD <- JPY_USD_value
	fmt.Println("JPY/USD: ", JPY_USD_value)
}

func simulateMarketData() {
	go simulateEUR()
	go simulateGBP()
	go simulateJPY()
	time.Sleep(1 * time.Second)
	simulateMarketData()
}

func sellEUR() {
	time.Sleep(4 * time.Second)
	fmt.Println("EUR/USD sold correctly")
}

func buyGBP() {
	time.Sleep(4 * time.Second)
	fmt.Println("GBP/USD sold correctly")
}

func buyJPY() {
	time.Sleep(4 * time.Second)
	fmt.Println("JPY/USD sold correctly")
}

func selectPair() {
	select {
	case EUR := <-EUR_USD:
		if EUR > 1.2 {
			fmt.Println("Selling EUR/USD")
			go sellEUR()
		}
	case GBP := <-GBP_USD:
		if GBP < 1.35 {
			fmt.Println("Buying GBP/USD")
			go buyGBP()
		}
	case JPY := <-JPY_USD:
		if JPY < 0.0085 {
			fmt.Println("Buying JPY/USD")
			go buyJPY()
		}
	}
	selectPair()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go simulateMarketData()
	go selectPair()
	time.Sleep(60 * time.Second)
}
