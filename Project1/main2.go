package main

import (
	"fmt"
	"sync"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

var mutex sync.Mutex
var repetitions int = 0

func check(caracter string, wg *sync.WaitGroup) {
	if caracter == toSearch {
		mutex.Lock()
		repetitions++
		mutex.Unlock()
	}
	defer wg.Done()
}

func main() {
	repetitions = 0
	var wg sync.WaitGroup

	for _, c := range whereSearch {
		wg.Add(1)
		go check(string(c), &wg)
	}

	wg.Wait()
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetitions)
}
