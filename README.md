# DTNeoway

Serviço Desenvolvido para o Desafio Técnico da Neoway

# Como executar

Necessário ter o Docker instalado

### Clone o repositório e execute o seguinte comando:

```
docker-compose up
```

### Para conectar no Banco de Dados

Use as configurações:

```
URL: localhost:54321
DB: buyers
User: postgres
Password: postgres
Table: compradores
```

# Estratégia:

A estratégia se baseia nos seguintes princípios:

### Leitura em blocos: 
O arquivo de entrada é lido em blocos de 1000 registros por vez, utilizando a função reader.LerArquivoCSV. Isso evita a leitura do arquivo inteiro na memória de uma só vez, otimizando o uso de memória e reduzindo o risco de falhas em conjuntos de dados muito grandes.

### Armazenamento em cache: 
Os registros de cada bloco são armazenados em um cache temporário compradoresCache. Isso permite a manipulação eficiente dos dados antes da persistência no banco de dados.

### Persistência em lotes: 
Quando o cache atinge 1000 registros, o conteúdo é enviado para o banco de dados de uma só vez, utilizando a função database.PersistirCompradores. Essa estratégia de persistência em lotes reduz o número total de escritas no banco de dados, diminuindo o overhead e melhorando o desempenho geral.

### Gerenciamento de goroutines: 
A persistência em lotes é realizada em goroutines separadas, utilizando canais para sincronizar o processo. Isso permite que a leitura e a persistência de dados sejam realizadas simultaneamente, otimizando ainda mais o tempo de processamento.

### Espera por goroutines: 
A função wg.Wait() garante que todas as goroutines de persistência sejam finalizadas antes de encerrar o programa principal. Isso evita a perda de dados e garante a integridade da transação.



# Estrutura relacional: 

```
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

```

