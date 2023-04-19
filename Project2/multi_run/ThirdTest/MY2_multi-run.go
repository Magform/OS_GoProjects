package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

func noleggia(needRent Cliente, wg *sync.WaitGroup, veicoli *[]*Veicolo) {
	defer wg.Done()
	var v int = rand.Intn(len(*veicoli))

	atomic.AddInt32(&(*veicoli)[v].utilized, 1)

	fmt.Printf("%s ha noleggiato il veicolo %s\n", needRent.nome, (*veicoli)[v].tipo)
}

func code() {
	rand.Seed(time.Now().UnixNano())

	veicoliDisponibili := generateVeicles(10000)

	clienti := generateClients(10)

	var wg sync.WaitGroup

	for _, c := range clienti {
		wg.Add(1)
		go noleggia(c, &wg, &veicoliDisponibili)
	}

	wg.Wait()

	for _, c := range veicoliDisponibili {

		fmt.Printf("%s noleggiate: %d\n", c.tipo, c.utilized)
	}
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

func generateVeicles(k int) []*Veicolo {
	tipi := []string{"Berlina", "Suv", "StationWagon"}
	rand.Seed(time.Now().UnixNano())

	veicoli := make([]*Veicolo, k)
	nextNumber := make(map[string]int)

	for i := 0; i < k; i++ {
		veicolo := tipi[rand.Intn(len(tipi))]

		// Genera un nuovo nome unico per il veicolo
		name := veicolo
		if nextNumber[veicolo] > 0 {
			name += strconv.Itoa(nextNumber[veicolo])
		}
		nextNumber[veicolo]++

		veicoli[i] = &Veicolo{tipo: name,
			utilized: 0}
	}

	return veicoli
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 1000; i++ {
		code()
	}
}
