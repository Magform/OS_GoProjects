package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Cliente struct {
	nome string
}
type Veicolo struct {
	tipo     string
	utilized int32
}

func noleggia(needRent Cliente, wg *sync.WaitGroup, veicoli *[]Veicolo) {
	defer wg.Done()
	var v int = rand.Intn(len(*veicoli))

	atomic.AddInt32(&(*veicoli)[v].utilized, 1)

	fmt.Printf("%s ha noleggiato il veicolo %s\n", needRent.nome, (*veicoli)[v].tipo)
}

func main() {
	defer timer("main")() //to see esecution time
	rand.Seed(time.Now().UnixNano())

	veicoliDisponibili := []Veicolo{
		{tipo: "Berlina",
			utilized: 0},
		{tipo: "SUV",
			utilized: 0},
		{tipo: "Station Wagon",
			utilized: 0},
	}

	clienti := []Cliente{
		{nome: "Mario"},
		{nome: "Luigi"},
		{nome: "Peach"},
		{nome: "Bowser"},
		{nome: "Yoshi"},
		{nome: "Toad"},
		{nome: "Wario"},
		{nome: "Waluigi"},
		{nome: "Donkey Kong"},
		{nome: "Daisy"},
	}

	var wg sync.WaitGroup

	for _, c := range clienti {
		wg.Add(1)
		go noleggia(c, &wg, &veicoliDisponibili)
	}

	wg.Wait()

	fmt.Printf("Berline noleggiate: %d\n", veicoliDisponibili[0].utilized)
	fmt.Printf("SUV noleggiate: %d\n", veicoliDisponibili[1].utilized)
	fmt.Printf("Station Wagon noleggiate: %d\n", veicoliDisponibili[2].utilized)
}

// To see esecution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
