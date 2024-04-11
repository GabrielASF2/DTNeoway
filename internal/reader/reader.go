package reader

import (
	"bufio"
	"buyers/domain/buyer"
	"buyers/internal/utils"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/klassmann/cpfcnpj"
	"github.com/pkg/errors"
)

func LerArquivoCSV(caminhoArquivo string) ([]buyer.Buyer, error) {
	arquivo, err := os.Open(caminhoArquivo)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao abrir o arquivo")
	}
	defer arquivo.Close()

	var re = regexp.MustCompile(" +")

	scanner := bufio.NewScanner(arquivo)
	linhas := make([]buyer.Buyer, 0)
	for scanner.Scan() {
		linha := scanner.Text()
		s := re.ReplaceAllString(linha, " ")
		espaco := strings.Split(s, " ")
		if string(espaco[0]) == "CPF" {
			continue
		}
		if len(espaco) != 8 {
			return nil, errors.New("Linha com quantidade de campos diferente de 8")
		}

		comprador := buyer.Buyer{
			CPF: utils.RemoverCaracteresEspeciais(espaco[0]),
		}

		if !ValidarCPF(comprador.CPF) {
			fmt.Println("CPF inválido:", comprador.CPF)
			comprador.CPFValido = false
		} else {
			comprador.CPFValido = true
		}

		if espaco[1] == "1" {
			comprador.Private = true
		}

		if espaco[2] == "1" {
			comprador.Incompleto = true
		}

		if espaco[3] != "NULL" {
			comprador.DataUltimaCompra = espaco[3]
		}

		if espaco[4] != "NULL" {
			comprador.TicketMedio = utils.ConverterStringToFloat64(espaco[4])
		}

		if espaco[5] != "NULL" {
			comprador.TicketUltimaCompra = utils.ConverterStringToFloat64(espaco[5])
		}

		comprador.LojaMaisFrequente = utils.RemoverCaracteresEspeciais(espaco[6])

		if !ValidarCNPJ(comprador.LojaMaisFrequente) {
			fmt.Println("CNPJ inválido:", strings.Join(espaco, ","))
			comprador.LojaMaisFrequenteValido = false
		} else {
			comprador.LojaMaisFrequenteValido = true
		}

		comprador.LojaUltimaCompra = utils.RemoverCaracteresEspeciais(espaco[7])

		if !ValidarCNPJ(comprador.LojaUltimaCompra) {
			fmt.Println("CNPJ inválido:", strings.Join(espaco, ","))
			comprador.LojaUltimaCompraValido = false
		} else {
			comprador.LojaUltimaCompraValido = true
		}

		linhas = append(linhas, comprador)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "Erro ao ler o arquivo")
	}

	return linhas, nil
}

func ValidarCPF(cpf string) bool {
	return cpfcnpj.ValidateCPF(cpf)
}

func ValidarCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}
