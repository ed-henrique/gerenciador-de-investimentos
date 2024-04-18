package internal

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrQuantidadeDeAtivosParaVendaExtrapolada = errors.New("Não é possível vender mais unidades do ativo do que as que estão disponíveis na carteira.")
)

func removePrimeiro(a []Ativo) []Ativo {
	return a[1:]
}

func QuantidadeAtivos(a []Ativo) (quantidade int) {
	for _, ai := range a {
		quantidade += ai.Quantidade
	}

	return
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
	ativos, ok := c[codigo]

	if !ok {
		return errors.New("É necessário que haja algum ativo com o código inserido.")
	}

	var quantidadeTotal int
	for _, a := range ativos {
		quantidadeTotal += a.Quantidade
	}

	if quantidade > quantidadeTotal {
		return ErrQuantidadeDeAtivosParaVendaExtrapolada
	}

	ativosParaVender := quantidade
	for ativosParaVender > 0 {
		if ativosParaVender < c[codigo][0].Quantidade {
			c[codigo][0].Quantidade -= ativosParaVender
			ativosParaVender = 0 
		} else if ativosParaVender == c[codigo][0].Quantidade {
			c[codigo] = removePrimeiro(c[codigo])
			ativosParaVender = 0 
		} else if quantidade > c[codigo][0].Quantidade {
			ativosParaVender -= c[codigo][0].Quantidade
			c[codigo] = removePrimeiro(c[codigo])
		}
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
		sb.WriteString(fmt.Sprintf("| %15s | %21d |\n", cod, QuantidadeAtivos(ativos)))
		sb.WriteString("-------------------------------------------\n")
	}

	sb.WriteString("\n")

	for cod, ativos := range c {
		sb.WriteString(fmt.Sprintf("%s\n", cod))
		sb.WriteString("-------------------------------------------------------------------\n")
		sb.WriteString("|        Quantidade |      Valor Unitario |        Data da Compra |\n")
		sb.WriteString("-------------------------------------------------------------------\n")

		for _, a := range ativos {
			sb.WriteString(fmt.Sprintf("| %17d | %19.2f | %21s |\n", a.Quantidade, a.ValorUnitario, a.DataDaCompra.Format("2006-01-02")))
		sb.WriteString("-------------------------------------------------------------------\n")
		}
	}

	return sb.String()
}

func (c Carteira) ImprimirResumoAtivos() string {
	var (
		sb                 strings.Builder
		quantidadeDeAtivos int
		valorTotal         float64
	)

	for _, ativos := range c {
		for _, a := range ativos {
			quantidadeDeAtivos += a.Quantidade
			valorTotal += a.ValorUnitario * float64(a.Quantidade)
		}
	}

	sb.WriteString("Carteira\n")
	sb.WriteString("--------------------------------------------\n")
	sb.WriteString("|  Quantidade de Ativos |      Valor Total |\n")
	sb.WriteString("--------------------------------------------\n")
	sb.WriteString(fmt.Sprintf("| %21d | %16.2f |\n", quantidadeDeAtivos, valorTotal))
	sb.WriteString("--------------------------------------------\n\n")

	sb.WriteString("Ativos\n")
	sb.WriteString("--------------------------------------------------------------\n")
	sb.WriteString("|          Codigo |      Quantidade Total |      Valor Total |\n")
	sb.WriteString("--------------------------------------------------------------\n")

	for codigo, ativos := range c {
		sb.WriteString(fmt.Sprintf("| %15s | %21d | ", codigo, QuantidadeAtivos(ativos)))

		var valorTotalAtivos float64
		for _, a := range ativos {
			valorTotalAtivos += a.ValorUnitario * float64(a.Quantidade)
		}

		sb.WriteString(fmt.Sprintf("%16.2f |\n", valorTotalAtivos))
		sb.WriteString("--------------------------------------------------------------\n")
	}

	return sb.String()
}
