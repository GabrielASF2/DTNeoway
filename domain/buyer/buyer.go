package buyer

type Buyer struct {
	CPF                     string  `db:"cpf"`
	CPFValido               bool    `db:"cpf_valido"`
	Private                 bool    `db:"private"`
	Incompleto              bool    `db:"incompleto"`
	DataUltimaCompra        string  `db:"data_ultima_compra"`
	TicketMedio             float64 `db:"ticket_medio"`
	TicketUltimaCompra      float64 `db:"ticket_ultima_compra"`
	LojaMaisFrequente       string  `db:"loja_mais_frequente"`
	LojaMaisFrequenteValido bool    `db:"loja_mais_frequente_valido"`
	LojaUltimaCompra        string  `db:"loja_ultima_compra"`
	LojaUltimaCompraValido  bool    `db:"loja_ultima_compra_valido"`
}
