package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Vehicle struct {
	model    string
	utilized int32
}

func noleggia(needRent string, wg *sync.WaitGroup, vehicle *[]Vehicle) {
	defer wg.Done()
	var v int = rand.Intn(len(*vehicle))

	atomic.AddInt32(&(*vehicle)[v].utilized, 1)

	fmt.Printf("%s has rented the vehicle %s\n", needRent, (*vehicle)[v].model)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	vehiclesAvailable := []Vehicle{
		{model: "Berlina", utilized: 0},
		{model: "SUV", utilized: 0},
		{model: "Station Wagon", utilized: 0},
	}

	clients := []string{
		"Mario",
		"Luigi",
		"Peach",
		"Bowser",
		"Yoshi",
		"Toad",
		"Wario",
		"Waluigi",
		"Donkey Kong",
		"Daisy",
	}

	var wg sync.WaitGroup

	for _, c := range clients {
		wg.Add(1)
		go noleggia(c, &wg, &vehiclesAvailable)
	}

	wg.Wait()

	fmt.Printf("Rented Berline: %d\n", vehiclesAvailable[0].utilized)
	fmt.Printf("Rented SUVs: %d\n", vehiclesAvailable[1].utilized)
	fmt.Printf("Rented Station Wagons: %d\n", vehiclesAvailable[2].utilized)
}
