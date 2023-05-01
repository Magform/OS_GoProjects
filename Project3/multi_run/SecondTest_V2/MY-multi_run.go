package main

import (
	"fmt"
	"sync"
	"time"
)

type semaphore struct {
	instances chan struct{}
	occupied  int
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
	s.occupied++
}

// release n resources
func (s *semaphore) V() {
	<-s.instances
	s.occupied--
}

func (s *semaphore) Occupied() int {
	return s.occupied
}

const torte = 1000

var Pasticciere1_spazi = costructor(2)
var Pasticciere2_spazi = costructor(2)

var torte1 = torte
var torte2 = torte
var torte3 = torte

var wg sync.WaitGroup

// cucinatore
func Pasticciere1() {
	for torte1 > 0 {
		Pasticciere1_spazi.P()
		torte1--
	}
	wg.Done()
}

// guarnitore
func Pasticciere2() {
	for torte2 > 0 {
		if Pasticciere1_spazi.Occupied() > 0 {
			Pasticciere2_spazi.P()
			Pasticciere1_spazi.V()
			torte2--
		}
	}
	wg.Done()
}

// decoratore
func Pasticciere3() {
	for torte3 > 0 {
		if Pasticciere2_spazi.Occupied() > 0 {
			Pasticciere2_spazi.V()
			torte3--
		}
	}

	wg.Done()
}

func code() {
	go Pasticciere1()
	go Pasticciere2()
	go Pasticciere3()
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
		torte1 = torte
		torte2 = torte
		torte3 = torte
		code()
	}
}
