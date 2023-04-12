package main

import (
	"fmt"
	"sync"
	"time"
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

func code() {
	var wg sync.WaitGroup
	repetitions = 0
	for _, c := range whereSearch {
		wg.Add(1)
		go check(string(c), &wg)
	}

	wg.Wait()
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetitions)
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
