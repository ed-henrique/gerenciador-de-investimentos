package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrQuantidadeDeAtivosParaVendaExtrapolada = errors.New("Não é possível vender mais unidades do ativo do que as que estão disponíveis na carteira.")
)

func remove(c Carteira, i int) (Carteira, error) {
	if i < 0 || i >= len(c) {
		return nil, errors.New("Não há como remover um elemento com índice negativo ou com índice maior ou igual ao tamanho do vetor.")
	}

	// Troca o elemento a ser removido com o último elemento do vetor
	// e devolve o vetor sem o último elemento
	c[i] = c[len(c)-1]
	return c[:len(c)-1], nil
}

type Ativo struct {
	Codigo        string
	Quantidade    int
	ValorUnitario float64
	DataDaCompra  time.Time
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

	if a.DataDaCompra.IsZero() {
		return errors.New("É necessário que haja uma data de compra para os ativos para adicioná-los na carteira.")
	}

	c = append(c, a)

	return nil
}

func (c Carteira) VenderAtivo(codigo string, quantidade int) error {
	var quantidadeTotal int

	for _, a := range c {
		if a.Codigo == codigo {
			quantidadeTotal += a.Quantidade
		}
	}

	if quantidade > quantidadeTotal {
		return ErrQuantidadeDeAtivosParaVendaExtrapolada
	}

	var i int
	for lenAtivos := len(c); lenAtivos != len(c) - quantidade; i++ {
		if c[i].Codigo == codigo {
			remove(c, i)
			lenAtivos--
		}
	}

	return nil
}

func (c Carteira) ImprimirAtivos() string {
	var sb strings.Builder
	ativosPorCodigo := map[string][]Ativo{}

	for _, a := range c {
		ativosPorCodigo[a.Codigo] = append(ativosPorCodigo[a.Codigo], a)
	}

	sb.WriteString("Ativos\n")
	sb.WriteString("-------------------------------------------\n")
	sb.WriteString("|          Codigo |      Quantidade Total |\n")
	sb.WriteString("-------------------------------------------\n")

	for cod, ativos := range ativosPorCodigo {
		sb.WriteString(fmt.Sprintf("| %15s | %21d |\n", cod, len(ativos)))
		sb.WriteString("-------------------------------------------\n")
	}

	sb.WriteString("\n")

	for cod, ativos := range ativosPorCodigo {
		sb.WriteString(fmt.Sprintf("%s\n", cod))
		sb.WriteString("-------------------------------------------\n")
		sb.WriteString("|  Valor Unitario |        Data da Compra |\n")
		sb.WriteString("-------------------------------------------\n")

		for _, a := range ativos {
			sb.WriteString(fmt.Sprintf("| %15.2f | %21s |\n", a.ValorUnitario, a.DataDaCompra.Format("2006-01-02")))
			sb.WriteString("-------------------------------------------\n")
		}
	}

	return sb.String()
}

func main() {}
