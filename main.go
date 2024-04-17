package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

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

	c = append(c, a)

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
