package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// struttura Cliente con il campo "nome"
type Cliente struct {
	nome string
}

// struttura Veicolo con il campo "tipo"
type Veicolo struct {
	tipo string
}

func code() {
	// Inizializzazione del generatore di numeri casuali
	rand.Seed(time.Now().UnixNano())

	// Definizione dei veicoli disponibili
	veicoli := generateVeicles(10000)
	// Definizione dei clienti
	clienti := generateClients(10)
	// Array per memorizzare i veicoli noleggiati
	noleggiati := make(map[string]int)

	// Mutex per proteggere l'accesso alla mappa noleggiati
	var mu sync.Mutex

	// Funzione per noleggiare un veicolo a caso
	noleggia := func(c Cliente, wg *sync.WaitGroup) {
		defer wg.Done()
		v := veicoli[rand.Intn(len(veicoli))]
		mu.Lock()
		noleggiati[v.tipo]++
		mu.Unlock()
		fmt.Printf("%s ha noleggiato il veicolo %s\n", c.nome, v.tipo)
	}

	// Sincronizzatore di gruppi di goroutine
	var wg sync.WaitGroup

	// Noleggio dei veicoli per tutti i clienti
	for _, c := range clienti {
		wg.Add(1)
		go noleggia(c, &wg)
	}

	// Attende il completamento di tutte le goroutine
	wg.Wait()

	// Funzione per stampare il numero di Berline, SUV e Station Wagon noleggiati
	stampa := func() {
		for _, c := range veicoli {
			fmt.Printf("%s noleggiate: %d\n", c.tipo, noleggiati[c.tipo])
		}
	}

	// Stampa del numero di veicoli noleggiati
	stampa()
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

func generateVeicles(k int) []Veicolo {
	tipi := []string{"Berlina", "Suv", "StationWagon"}
	rand.Seed(time.Now().UnixNano())

	veicoli := make([]Veicolo, k)
	nextNumber := make(map[string]int)

	for i := 0; i < k; i++ {
		veicolo := tipi[rand.Intn(len(tipi))]

		// Genera un nuovo nome unico per il veicolo
		name := veicolo
		if nextNumber[veicolo] > 0 {
			name += strconv.Itoa(nextNumber[veicolo])
		}
		nextNumber[veicolo]++

		veicoli[i] = Veicolo{tipo: name}
	}

	return veicoli
}

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 1000; i++ {
		code()
	}
}
