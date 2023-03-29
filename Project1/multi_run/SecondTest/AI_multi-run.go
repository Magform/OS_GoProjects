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

func countOccurrences(s string, c byte, wg *sync.WaitGroup, count chan int) {
	defer wg.Done()
	var occurrences int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			occurrences++
		}
	}
	count <- occurrences
}

func code() {
	s := RandStringBytes(10000)
	c := byte('c')
	var wg sync.WaitGroup
	count := make(chan int)
	for i := 0; i < len(s); i++ {
		wg.Add(1)
		go countOccurrences(string(s[i]), c, &wg, count)
	}
	go func() {
		wg.Wait()
		close(count)
	}()
	var totalOccurrences int
	for occurrences := range count {
		totalOccurrences += occurrences
	}
	fmt.Printf("Il carattere %c compare %d volte nella stringa.\n", c, totalOccurrences)
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
