package main

import (
	"fmt"
	"math/rand"
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

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the available vehicles
	vehicles := []Vehicle{
		{model: "Berlina"},
		{model: "SUV"},
		{model: "Station Wagon"},
	}

	// Define the clients
	clients := []Client{
		{name: "Mario"},
		{name: "Luigi"},
		{name: "Peach"},
		{name: "Bowser"},
		{name: "Yoshi"},
		{name: "Toad"},
		{name: "Wario"},
		{name: "Waluigi"},
		{name: "Donkey Kong"},
		{name: "Daisy"},
	}

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
		fmt.Printf("Rented Berline: %d\n", rented["Berlina"])
		fmt.Printf("Rented SUVs: %d\n", rented["SUV"])
		fmt.Printf("Rented Station Wagons: %d\n", rented["Station Wagon"])
	}

	// Print the number of rented vehicles
	print()
}
