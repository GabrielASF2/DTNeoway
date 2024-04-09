package reader

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// LerArquivoCSV lê um arquivo CSV e retorna um slice de strings.
func LerArquivoCSV(caminhoArquivo string) ([]string, error) {
	arquivo, err := os.Open(caminhoArquivo)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao abrir o arquivo")
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	linhas := make([]string, 0)
	for scanner.Scan() {
		linha := scanner.Text()
		// Tratar separador não convencional
		linha = strings.ReplaceAll(linha, "  ", ",")
		linhas = append(linhas, linha)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "Erro ao ler o arquivo")
	}

	return linhas, nil
}
