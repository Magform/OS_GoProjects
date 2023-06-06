package main

import (
	"fmt"
	"sync"
)

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

func main() {
	s := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
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
