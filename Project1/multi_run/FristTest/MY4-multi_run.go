package main

import (
	"fmt"
	"sync"
	"time"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

var mutex sync.Mutex
var repetitions [100000000]int

var n int

func check(caracter string, wg *sync.WaitGroup, n int) {
	if caracter == toSearch {
		repetitions[n] = 1
	} else {
		repetitions[n] = 0
	}
	defer wg.Done()
}

func t(s1 string, wg *sync.WaitGroup) {
	for _, c := range s1 {
		wg.Add(1)
		go check(string(c), wg, n)
		n++
	}
	wg.Done()
}

func code() {
	n = 0

	var wg sync.WaitGroup

	wg.Add(1)
	go t(whereSearch, &wg)

	wg.Wait()

	var repetition int = 0
	var i int = 0

	for i < n {
		repetition = repetition + repetitions[i]
		i++
	}
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetition)
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
