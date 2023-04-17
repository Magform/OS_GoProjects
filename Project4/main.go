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
	EUR_USD <- rand.Float64()*0.5 + 1
}

func simulateGBP() {
	GBP_USD <- rand.Float64()*0.5 + 1
}

func simulateJPY() {
	JPY_USD <- rand.Float64()*0.003 + 0.006
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
}

func buyGBP() {
	time.Sleep(4 * time.Second)
}

func buyJPY() {
	time.Sleep(4 * time.Second)
}

func selectPair() {
	select {
	case EUR := <-EUR_USD:
		if EUR > 1.2 {
			go sellEUR()
		}
	case GBP := <-GBP_USD:
		if GBP < 1.35 {
			go buyGBP()
		}
	case JPY := <-JPY_USD:
		if JPY < 0.0085 {
			go buyJPY()
		}
	}
	selectPair()
}

func main() {
	defer timer("main")() //to see esecution time
	rand.Seed(time.Now().UnixNano())
	go simulateMarketData()
	go selectPair()
	time.Sleep(60 * time.Second)

}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
