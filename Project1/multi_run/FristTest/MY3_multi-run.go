package main

import (
	"fmt"
	"sync"
	"time"
)

const toSearch string = "c"
const whereSearch string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"

var mutex1 sync.Mutex
var mutex2 sync.Mutex
var mutex3 sync.Mutex

var repetitions1 int = 0
var repetitions2 int = 0
var repetitions3 int = 0

func check1(caracter string, wg *sync.WaitGroup) {
	if caracter == toSearch {
		mutex1.Lock()
		repetitions1++
		mutex1.Unlock()
	}
	defer wg.Done()
}

func check2(caracter string, wg *sync.WaitGroup) {
	if caracter == toSearch {
		mutex2.Lock()
		repetitions2++
		mutex2.Unlock()
	}
	defer wg.Done()
}

func check3(caracter string, wg *sync.WaitGroup) {
	if caracter == toSearch {
		mutex3.Lock()
		repetitions3++
		mutex3.Unlock()
	}
	defer wg.Done()
}

func t1(s1 string, wg *sync.WaitGroup) {
	for _, c := range s1 {
		wg.Add(1)
		go check1(string(c), wg)
	}
	wg.Done()
}

func t2(s1 string, wg *sync.WaitGroup) {
	for _, c := range s1 {
		wg.Add(1)
		go check2(string(c), wg)
	}
	wg.Done()
}

func t3(s1 string, wg *sync.WaitGroup) {
	for _, c := range s1 {
		wg.Add(1)
		go check3(string(c), wg)
	}
	wg.Done()
}

func code() {
	repetitions1 = 0
	repetitions2 = 0
	repetitions3 = 0
	var wg sync.WaitGroup

	var len int = len(whereSearch)
	var s1 = whereSearch[0 : len/3]
	var s2 = whereSearch[len/3 : (len-len/3)/2]
	var s3 = whereSearch[(len-len/3)/2 : len-(len-len/3)/2]

	wg.Add(3)
	go t1(s1, &wg)
	go t2(s2, &wg)
	go t3(s3, &wg)

	wg.Wait()
	var repetitions = repetitions1 + repetitions2 + repetitions3
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetitions)
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 1000000; i++ {
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
