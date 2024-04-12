CREATE TABLE IF NOT EXISTS compradores (
  cpf VARCHAR(14) PRIMARY KEY,
  cpf_valid BOOLEAN NOT NULL,
  private BOOLEAN NOT NULL,
  incompleto BOOLEAN NOT NULL,
  data_ultima_compra DATE,
  ticket_medio DECIMAL(10,2),
  ticket_ultima_compra DECIMAL(10,2),
  loja_mais_frequente VARCHAR(255),
  loja_mais_frequente_valid BOOLEAN NOT NULL,
  loja_ultima_compra VARCHAR(255),
  loja_ultima_compra_valid BOOLEAN NOT NULL
);
