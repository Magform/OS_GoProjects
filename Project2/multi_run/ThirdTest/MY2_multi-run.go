package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Vehicle struct {
	model    string
	utilized int32
}

func noleggia(needRent string, wg *sync.WaitGroup, vehicle *[]*Vehicle) {
	defer wg.Done()
	var v int = rand.Intn(len(*vehicle))

	atomic.AddInt32(&(*vehicle)[v].utilized, 1)

	fmt.Printf("%s has rented the vehicle %s\n", needRent, (*vehicle)[v].model)
}

func code() {
	rand.Seed(time.Now().UnixNano())

	vehiclesAvailable := generateVeicles(10000)

	clients := generateClients(1000)

	var wg sync.WaitGroup

	for _, c := range clients {
		wg.Add(1)
		go noleggia(c, &wg, &vehiclesAvailable)
	}

	wg.Wait()

	for _, c := range vehiclesAvailable {

		fmt.Printf("%s noleggiate: %d\n", c.model, c.utilized)
	}
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// function to generate clients
func generateClients(k int) []string {
	names := []string{"Mario", "Luigi", "Peach", "Bowser", "Yoshi", "Toad", "Wario", "Waluigi", "Donkey Kong", "Daisy"}
	rand.Seed(time.Now().UnixNano())

	clients := make([]string, k)
	for i := 0; i < k; i++ {
		clients[i] = names[rand.Intn(len(names))]
	}

	return clients
}

func generateVeicles(k int) []*Vehicle {
	models := []string{"Berlina", "Suv", "StationWagon"}
	rand.Seed(time.Now().UnixNano())

	vehicles := make([]*Vehicle, k)
	nextNumber := make(map[string]int)

	for i := 0; i < k; i++ {
		vehicle := models[rand.Intn(len(models))]

		name := vehicle
		if nextNumber[vehicle] > 0 {
			name += strconv.Itoa(nextNumber[vehicle])
		}
		nextNumber[vehicle]++

		vehicles[i] = &Vehicle{model: name, utilized: 0}
	}

	return vehicles
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000; i++ {
		code()
	}
}
