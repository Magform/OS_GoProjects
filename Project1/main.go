package main

import (
	"fmt"
	"sync"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

func check(caracter string, repetitions chan int, wg *sync.WaitGroup) {
	if caracter == toSearch {
		rep := <-repetitions
		rep++
		repetitions <- rep
	}
	defer wg.Done()
}

func main() {
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
