package main

import (
	"buyers/domain/buyer"
	"fmt"
	"os"

	"buyers/internal/database"
	"buyers/internal/reader"
	"buyers/internal/utils"
	"buyers/internal/validator"
)

func main() {
	caminhoArquivo := "base_teste.txt"

	linhas, err := reader.LerArquivoCSV(caminhoArquivo)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		os.Exit(1)
	}

	db, err := database.ConectarDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		os.Exit(1)
	}

	// Processar linhas
	compradores := make([]buyer.Buyer, 0)
	for _, linha := range linhas {
		// Ignorar header
		if string(linha[0]) == "CPF" {
			continue
		}

		comprador := buyer.Buyer{
			CPF: utils.RemoverCaracteresEspeciais(utils.ToLower(utils.RemoverAcentos(string(linha[0])))),
			// ...
		}

		if !validator.ValidarCPF(comprador.CPF) {
			comprador.CPF = "CPF inválido"
		}
		// ...

		compradores = append(compradores, comprador)
	}

	// Persistir os dados no banco de dados
	err = database.PersistirCompradores(db, compradores)
	if err != nil {
		fmt.Println("Erro ao persistir os dados:", err)
		os.Exit(1)
	}

	fmt.Println("Dados persistidos com sucesso!")
}
