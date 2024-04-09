package validator

import (
	"github.com/klassmann/cpfcnpj"
)

func ValidarCPF(cpf string) bool {
	return cpfcnpj.ValidateCPF(cpf)
}

func ValidarCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}
