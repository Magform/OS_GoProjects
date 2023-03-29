package main

import (
	"fmt"
	"math/rand"
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
	veicoli := []Veicolo{
		{tipo: "Berlina"},
		{tipo: "SUV"},
		{tipo: "Station Wagon"},
	}

	// Definizione dei clienti
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

	// Array per memorizzare i veicoli noleggiati
	noleggiati := make(map[string]int)

	// Funzione per noleggiare un veicolo a caso
	noleggia := func(c Cliente, wg *sync.WaitGroup) {
		defer wg.Done()
		v := veicoli[rand.Intn(len(veicoli))]
		noleggiati[v.tipo]++
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
		fmt.Printf("Berline noleggiate: %d\n", noleggiati["Berlina"])
		fmt.Printf("SUV noleggiate: %d\n", noleggiati["SUV"])
		fmt.Printf("Station Wagon noleggiate: %d\n", noleggiati["Station Wagon"])
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

func main() {
	defer timer("main")() //to see esecution time
	for i := 0; i < 100000; i++ {
		code()
	}
}
