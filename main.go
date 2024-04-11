package main

import (
	"buyers/domain/buyer"
	"buyers/internal/database"
	"buyers/internal/reader"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	inicio := time.Now()
	caminhoArquivo := "base_teste.txt"

	compradores, err := reader.LerArquivoCSV(caminhoArquivo)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		os.Exit(1)
	}

	db, err := database.ConectarDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		os.Exit(1)
	}

	canalCompradores := make(chan bool, 100)
	var wg sync.WaitGroup
	compradoresCache := []buyer.Buyer{}

	for _, comprador := range compradores {
		compradoresCache = append(compradoresCache, comprador)
		if len(compradoresCache) == 1000 {
			compradoresBatch := append([]buyer.Buyer{}, compradoresCache...)
			wg.Add(1)
			canalCompradores <- true
			go database.PersistirCompradores(db, compradoresBatch, canalCompradores, &wg)
			compradoresCache = []buyer.Buyer{}
		}
	}

	wg.Add(1)
	canalCompradores <- true
	go database.PersistirCompradores(db, compradoresCache, canalCompradores, &wg)
	wg.Wait()

	fmt.Println("Dados persistidos com sucesso!")

	fim := time.Now()

	tempoTotal := fim.Sub(inicio)

	fmt.Println("Tempo total:", tempoTotal)
}
