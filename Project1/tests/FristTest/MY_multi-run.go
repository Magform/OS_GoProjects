package main

import (
	"fmt"
	"sync"
	"time"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

func check(character string, repetitions chan int, wg *sync.WaitGroup) {
	if character == toSearch {
		rep := <-repetitions
		rep++
		repetitions <- rep
	}
	defer wg.Done()
}

func code() {
	results := make(chan int, 1)
	results <- 0

	var wg sync.WaitGroup

	for _, c := range whereSearch {
		wg.Add(1)
		go check(string(c), results, &wg)
	}

	wg.Wait()
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", <-results)
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000000; i++ {
		code()
	}
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
