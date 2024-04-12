package database

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"buyers/domain/buyer"
)

func ConectarDB() (*sql.DB, error) {
	conexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db", "5432", "postgres", "postgres", "buyers")

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao conectar ao banco de dados")
	}

	return db, nil
}
func PersistirCompradores(db *sql.DB, compradores []buyer.Buyer, canalCompradores chan bool, wg *sync.WaitGroup) error {
	defer func() {
		<-canalCompradores
		wg.Done()
	}()

	values := []string{}
	for _, comprador := range compradores {
		if comprador.DataUltimaCompra == "" {
			values = append(values, fmt.Sprintf("('%s', %t, %t, %t, null, %f, %f, '%s', %t, '%s', %t)",
				comprador.CPF, comprador.CPFValido, comprador.Private, comprador.Incompleto, comprador.TicketMedio, comprador.TicketUltimaCompra, comprador.LojaMaisFrequente, comprador.LojaMaisFrequenteValido, comprador.LojaUltimaCompra, comprador.LojaMaisFrequenteValido))
		} else {
			values = append(values, fmt.Sprintf("('%s', %t, %t, %t, '%s', %f, %f, '%s', %t, '%s', %t)",
				comprador.CPF, comprador.CPFValido, comprador.Private, comprador.Incompleto, comprador.DataUltimaCompra, comprador.TicketMedio, comprador.TicketUltimaCompra, comprador.LojaMaisFrequente, comprador.LojaMaisFrequenteValido, comprador.LojaUltimaCompra, comprador.LojaMaisFrequenteValido))
		}
	}

	query := fmt.Sprintf(`INSERT INTO compradores (cpf, cpf_valid, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_mais_frequente_valid, loja_ultima_compra, loja_ultima_compra_valid) VALUES %s`, strings.Join(values, ","))

	_, err := db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "Erro ao iniciar a transação")
	}
	return nil
}
