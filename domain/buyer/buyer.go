package buyer

type Buyer struct {
	CPF                string  `db:"cpf"`
	Nome               string  `db:"nome"`
	Status             string  `db:"status"`
	DataUltimaCompra   string  `db:"data_ultima_compra"`
	TicketMedio        float64 `db:"ticket_medio"`
	TicketUltimaCompra float64 `db:"ticket_ultima_compra"`
	LojaMaisFrequente  string  `db:"loja_mais_frequente"`
	LojaUltimaCompra   string  `db:"loja_ultima_compra"`
}
