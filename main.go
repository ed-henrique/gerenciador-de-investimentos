package main

import "errors"

type Ativo struct {
	Codigo        string
	Quantidade    int
	ValorUnitario float64
}

type Carteira []Ativo

func (c Carteira) AdicionarAtivo(a Ativo) error {
	if a.Codigo == "" {
		return errors.New("É necessário que o código do ativo tenha ao menos um caractere para adicioná-lo na carteira.")
	}

	if a.ValorUnitario <= 0.0 {
		return errors.New("É necessário que o valor unitário do ativo seja positivo para adicioná-lo na carteira.")
	}

	if a.Quantidade <= 0 {
		return errors.New("É necessário que a quantidade de ativos seja positiva para adicioná-los na carteira.")
	}

	c = append(c, a)

	return nil
}

func main() {}
