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

func code() {
	rand.Seed(time.Now().UnixNano())

	veicoliDisponibili := []Veicolo{
		{tipo: "Berlina",
			utilized: 0},
		{tipo: "SUV",
			utilized: 0},
		{tipo: "Station Wagon",
			utilized: 0},
	}

	clienti := generateClients(1000)

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

// function to generate clients
func generateClients(k int) []Cliente {
	nomi := []string{"Mario", "Luigi", "Peach", "Bowser", "Yoshi", "Toad", "Wario", "Waluigi", "Donkey Kong", "Daisy"}
	rand.Seed(time.Now().UnixNano())

	clienti := make([]Cliente, k)
	for i := 0; i < k; i++ {
		nome := nomi[rand.Intn(len(nomi))]
		clienti[i] = Cliente{nome: nome}
	}

	return clienti
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 10000; i++ {
		code()
	}
}
