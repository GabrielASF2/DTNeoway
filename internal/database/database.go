package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"buyers/domain/buyer"
)

func ConectarDB() (*sql.DB, error) {
	conexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "buyers")

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao conectar ao banco de dados")
	}

	return db, nil
}

func PersistirCompradores(db *sql.DB, compradores []buyer.Buyer) error {

	query := `INSERT INTO compradores (cpf, nome, status, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "Erro ao preparar a query")
	}
	defer stmt.Close()

	for _, comprador := range compradores {
		_, err := stmt.Exec(comprador.CPF, comprador.Nome, comprador.Status, comprador.DataUltimaCompra, comprador.TicketMedio, comprador.TicketUltimaCompra, comprador.LojaMaisFrequente, comprador.LojaUltimaCompra)
		if err != nil {
			return errors.Wrap(err, "Erro ao executar a query")
		}
	}

	return nil
}
