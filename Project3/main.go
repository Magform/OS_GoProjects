package main

import (
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

var cookTime = 1 * time.Second
var icerTime = 4 * time.Second
var decoratorTime = 8 * time.Second
var totalCake = 5

var wg sync.WaitGroup

// cook
func pastryChef1() {
	cakes := 1
	for totalCake >= cakes {
		time.Sleep(cookTime)
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
			time.Sleep(icerTime)
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
			time.Sleep(decoratorTime)
			pastryChef2_spaces.V()
			cakes++
		}
	}

	wg.Done()
}

func main() {
	go pastryChef1()
	go pastryChef2()
	go pastryChef3()
	wg.Add(3)
	wg.Wait()
}
