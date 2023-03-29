package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const toSearch string = "c"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

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
	var whereSearch string = RandStringBytes(10000)
	for _, c := range whereSearch {
		wg.Add(1)
		go check(string(c), &wg)
	}

	wg.Wait()
	fmt.Printf("Il carattere cercato compare %d volte nella stringa.\n", repetitions)
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000; i++ {
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
