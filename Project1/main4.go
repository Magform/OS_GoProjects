package main

import (
	"fmt"
	"sync"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

var mutex sync.Mutex
var repetitions [100000000]int //this need to be higher than string lenght

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

func main() {
	n = 0

	var wg sync.WaitGroup

	wg.Add(1)
	go t(whereSearch, &wg)

	wg.Wait()

	var repetition int
	var i int

	for i < n {
		repetition += repetitions[i]
		i++
	}
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetition)
}
