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

func removeLast(a []Ativo) []Ativo {
	return a[:len(a)-1]
}

type Ativo struct {
	Codigo        string
	Quantidade    int
	ValorUnitario float64
	DataDaCompra  time.Time
}

type Carteira map[string][]Ativo

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

	c[a.Codigo] = append(c[a.Codigo], a)

	return nil
}

func (c Carteira) VenderAtivo(codigo string, quantidade int) error {
	ativos, ok := c[codigo]; 

	if !ok {
		return errors.New("É necessário que haja algum ativo com o código inserido.")
	}

	if quantidade > len(ativos) {
		return ErrQuantidadeDeAtivosParaVendaExtrapolada
	}

	for i := 0; i < quantidade; i++ {
		c[codigo] = removeLast(c[codigo])
	}

	return nil
}

func (c Carteira) ImprimirAtivos() string {
	var sb strings.Builder

	sb.WriteString("Ativos\n")
	sb.WriteString("-------------------------------------------\n")
	sb.WriteString("|          Codigo |      Quantidade Total |\n")
	sb.WriteString("-------------------------------------------\n")

	for cod, ativos := range c {
		sb.WriteString(fmt.Sprintf("| %15s | %21d |\n", cod, len(ativos)))
		sb.WriteString("-------------------------------------------\n")
	}

	sb.WriteString("\n")

	for cod, ativos := range c {
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
