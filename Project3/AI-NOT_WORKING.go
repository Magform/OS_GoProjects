package main

import (
	"fmt"
	"sync"
	"time"
)

const numCakes = 5
const numCooks = 1
const numDecorators = 1
const numIcers = 1

var (
	cakeCounter    int
	cookSpace      = make(chan struct{}, 2)
	icerSpace      = make(chan struct{}, 2)
	decoratorSpace = make(chan struct{}, 2)
	cookTime       = time.Second
	icerTime       = 4 * time.Second
	decoratorTime  = 8 * time.Second
)

func main() {
	fmt.Println("Starting cake production")

	var wg sync.WaitGroup
	wg.Add(numCooks + numDecorators + numIcers)

	for i := 0; i < numCakes; i++ {
		cookSpace <- struct{}{} // wait for cook space
		wg.Add(1)
		go cookCake(&wg)
	}

	wg.Wait()
	fmt.Println("All cakes produced")
}

func cookCake(wg *sync.WaitGroup) {
	defer wg.Done()

	cakeID := getNextCakeID()
	fmt.Printf("Cooking cake %d\n", cakeID)

	time.Sleep(cookTime)

	fmt.Printf("Finished cooking cake %d\n", cakeID)

	select {
	case icerSpace <- struct{}{}:
		wg.Add(1)
		go iceCake(cakeID, wg)
	case cookSpace <- struct{}{}:
		wg.Add(1)
		go cookCake(wg)
	}
}

func iceCake(cakeID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Icing cake %d\n", cakeID)

	time.Sleep(icerTime)

	fmt.Printf("Finished icing cake %d\n", cakeID)

	select {
	case decoratorSpace <- struct{}{}:
		wg.Add(1)
		go decorateCake(cakeID, wg)
	case icerSpace <- struct{}{}:
		wg.Add(1)
		go iceCake(cakeID, wg)
	}
	<-icerSpace // free up icer space
}

func decorateCake(cakeID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Decorating cake %d\n", cakeID)

	time.Sleep(decoratorTime)

	fmt.Printf("Finished decorating cake %d\n", cakeID)

	select {
	case <-cookSpace:
		cakeCounter++
		if cakeCounter == numCakes {
			wg.Done()
			return
		}
		wg.Add(1)
		go cookCake(wg)
	case decoratorSpace <- struct{}{}:
		wg.Add(1)
		go decorateCake(cakeID, wg)
	}
	<-decoratorSpace // free up decorator space
}

func getNextCakeID() int {
	return cakeCounter + 1
}
