package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func RemoverCaracteresEspeciais(s string) string {
	var re = regexp.MustCompile("[^a-zA-Z0-9 ]")
	return re.ReplaceAllString(s, "")
}

func ConverterStringToFloat64(valor string) float64 {
	valor = strings.ReplaceAll(valor, ",", ".")
	valorConvertido, _ := strconv.ParseFloat(valor, 64)
	return valorConvertido
}
