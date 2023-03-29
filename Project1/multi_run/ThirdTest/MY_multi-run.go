package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

const toSearch string = "c"

func check(caracter string, repetitions chan int, wg *sync.WaitGroup) {
	if caracter == toSearch {
		rep := <-repetitions
		rep++
		repetitions <- rep
	}
	defer wg.Done()
}

func code() {
	var whereSearch string = RandStringBytes(100000)
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
	for i := 0; i < 1000; i++ {
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
