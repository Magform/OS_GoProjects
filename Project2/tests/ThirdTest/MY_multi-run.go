package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Vehicle struct {
	model    string
	utilized int
	mutex    sync.Mutex
}

func rent(needRent string, wg *sync.WaitGroup, vehicles *[]*Vehicle) {
	defer wg.Done()
	var v int = rand.Intn(len(*vehicles))

	(*vehicles)[v].mutex.Lock()
	(*vehicles)[v].utilized++
	(*vehicles)[v].mutex.Unlock()

	fmt.Printf("%s has rented the vehicle %s\n", needRent, (*vehicles)[v].model)
}

func code(vehiclesAvailable []*Vehicle, clients []string) {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for _, c := range clients {
		wg.Add(1)
		go rent(c, &wg, &vehiclesAvailable)
	}

	wg.Wait()

	for _, c := range vehiclesAvailable {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		fmt.Printf("%s rented: %d\n", c.model, c.utilized)
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
	vehiclesAvailable := generateVeicles(10000)
	clients := generateClients(1000)
	for i := 0; i < 10000; i++ {
		code(vehiclesAvailable, clients)
	}
}
