package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Client structure with the
type Client struct {
	name string
}

// Vehicle  structure with the
type Vehicle struct {
	model string
}

func code() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the available vehicles
	vehicles := generateVeicles(10000)

	clients := generateClients(1000)

	// Map to store the rented vehicles
	rented := make(map[string]int)

	// Mutex to protect access to the rented map
	var mu sync.Mutex

	// Function to rent a random vehicle
	rent := func(c Client, wg *sync.WaitGroup) {
		defer wg.Done()
		v := vehicles[rand.Intn(len(vehicles))]
		mu.Lock()
		rented[v.model]++
		mu.Unlock()
		fmt.Printf("%s has rented the vehicle %s\n", c.name, v.model)
	}

	// Synchronizer for groups of goroutines
	var wg sync.WaitGroup

	// Rent vehicles for all clients
	for _, c := range clients {
		wg.Add(1)
		go rent(c, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Function to print the number of rented Berline, SUV, and Station Wagons
	print := func() {
		for _, c := range vehicles {
			fmt.Printf("%s noleggiate: %d\n", c.model, rented[c.model])
		}
	}

	// Print the number of rented vehicles
	print()
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// function to generate clients
func generateClients(k int) []Client {
	names := []string{"Mario", "Luigi", "Peach", "Bowser", "Yoshi", "Toad", "Wario", "Waluigi", "Donkey Kong", "Daisy"}
	rand.Seed(time.Now().UnixNano())

	clients := make([]Client, k)
	for i := 0; i < k; i++ {
		name := names[rand.Intn(len(names))]
		clients[i] = Client{name: name}
	}

	return clients
}

func generateVeicles(k int) []Vehicle {
	models := []string{"Berlina", "Suv", "StationWagon"}
	rand.Seed(time.Now().UnixNano())

	vehicles := make([]Vehicle, k)
	nextNumber := make(map[string]int)

	for i := 0; i < k; i++ {
		vehicle := models[rand.Intn(len(models))]
		name := vehicle
		if nextNumber[vehicle] > 0 {
			name += strconv.Itoa(nextNumber[vehicle])
		}
		nextNumber[vehicle]++

		vehicles[i] = Vehicle{model: name}
	}

	return vehicles
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000; i++ {
		code()
	}
}
