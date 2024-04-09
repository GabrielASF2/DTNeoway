package utils

import (
	"strings"
)

func RemoverAcentos(s string) string {
	return strings.ReplaceAll(s, "áàãâäéèêëíìîïóòôõöúùûüç", "aaaaaeeeeiiiiooooouuuuc")
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func RemoverCaracteresEspeciais(s string) string {
	return strings.ReplaceAll(s, "[^a-zA-Z0-9 ]", "")
}
