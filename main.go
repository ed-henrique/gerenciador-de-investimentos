package main

import "errors"

// Essa variável global está simulando um banco de dados
// que armazenaria as informações relativas aos ativos
var ativos = []Ativo{}

type Ativo struct {
	Codigo        string
	Quantidade    int
	ValorUnitario float64
}

func AdicionarAtivo(a Ativo) error {
	if a.Codigo == "" {
		return errors.New("É necessário que o código do ativo tenha ao menos um caractere para adicioná-lo na carteira.")
	}

	if a.ValorUnitario <= 0.0 {
		return errors.New("É necessário que o valor unitário do ativo seja positivo para adicioná-lo na carteira.")
	}

	if a.Quantidade <= 0 {
		return errors.New("É necessário que a quantidade de ativos seja positiva para adicioná-los na carteira.")
	}

	ativos = append(ativos, a)

	return nil
}

func main() {}
