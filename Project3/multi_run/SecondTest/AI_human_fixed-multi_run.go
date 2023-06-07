package main

import (
	"fmt"
	"sync"
	"time"
)

const numCakes = 1000
const numCooks = 1
const numDecorators = 1
const numIcers = 1

var (
	cakeCounter   int
	cookedSpace   = make(chan struct{}, 2)
	iceredSpace   = make(chan struct{}, 2)
	cooks         = make(chan struct{}, numCooks)
	decorators    = make(chan struct{}, numDecorators)
	icers         = make(chan struct{}, numIcers)
	cookTime      = (1 * time.Second) / 100
	icerTime      = (4 * time.Second) / 100
	decoratorTime = (8 * time.Second) / 100
)

func code() {
	fmt.Println("Starting cake production")
	var wg sync.WaitGroup

	for i := 0; i < numCakes; i++ {
		cakeID := getNextCakeID()
		wg.Add(1)
		go cookCake(cakeID, &wg)
	}

	wg.Wait()
	fmt.Println("All cakes produced")
}

func cookCake(cakeID int, wg *sync.WaitGroup) {
	cooks <- struct{}{}

	fmt.Printf("Cooking cake %d\n", cakeID)

	time.Sleep(cookTime)

	fmt.Printf("Finished cooking cake %d\n", cakeID)

	cookedSpace <- struct{}{}
	wg.Add(1)
	go iceCake(cakeID, wg)

	<-cooks
	defer wg.Done()
}

func iceCake(cakeID int, wg *sync.WaitGroup) {
	icers <- struct{}{}
	fmt.Printf("Icing cake %d\n", cakeID)

	time.Sleep(icerTime)

	fmt.Printf("Finished icing cake %d\n", cakeID)

	iceredSpace <- struct{}{}
	wg.Add(1)
	go decorateCake(cakeID, wg)

	<-cookedSpace // free up icer space
	<-icers
	defer wg.Done()
}

func decorateCake(cakeID int, wg *sync.WaitGroup) {
	decorators <- struct{}{}
	fmt.Printf("Decorating cake %d\n", cakeID)

	time.Sleep(decoratorTime)

	fmt.Printf("Finished decorating cake %d\n", cakeID)

	<-iceredSpace // free up decorator space
	<-decorators
	defer wg.Done()
}

func getNextCakeID() int {
	cakeCounter++
	return cakeCounter
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10; i++ {
		cakeCounter = 0
		code()
	}
}
