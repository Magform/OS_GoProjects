package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type semaphore struct {
	instances chan struct{}
	occupied  int64
}

func costructor(n int) *semaphore {
	return &semaphore{
		instances: make(chan struct{}, n),
		occupied:  0,
	}
}

// acquire n resources
func (s *semaphore) P() {
	s.instances <- struct{}{}
	atomic.AddInt64(&s.occupied, 1)
}

// release n resources
func (s *semaphore) V() {
	<-s.instances
	atomic.AddInt64(&s.occupied, -1)
}

func (s *semaphore) Occupied() int64 {
	return s.occupied
}

var pastryChef1_spaces = costructor(2)
var pastryChef2_spaces = costructor(2)

var totalCake = 5

var wg sync.WaitGroup

// cook
func pastryChef1() {
	cakes := 1
	for totalCake >= cakes {
		pastryChef1_spaces.P()
		cakes++
	}
	wg.Done()
}

// icer
func pastryChef2() {
	cakes := 1
	for totalCake >= cakes {
		if pastryChef1_spaces.Occupied() > 0 {
			pastryChef2_spaces.P()
			pastryChef1_spaces.V()
			cakes++
		}
	}
	wg.Done()
}

// decorator
func pastryChef3() {
	cakes := 1
	for totalCake >= cakes {
		if pastryChef2_spaces.Occupied() > 0 {
			pastryChef2_spaces.V()
			cakes++
		}
	}

	wg.Done()
}

func code() {
	go pastryChef1()
	go pastryChef2()
	go pastryChef3()
	wg.Add(3)
	wg.Wait()
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000; i++ {
		code()
	}
}
