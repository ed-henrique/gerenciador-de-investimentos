package beautiful

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	COR_TITULO = "#04F06A"
	COR_ERRO   = "#E74B56"
	COR_FOCO   = "#758BFD"
	COR_FUNDO  = "#494949"
	COR_ESCURA = "#031A8C"

	BOTAO_ADICIONAR = "Adicionar"
	BOTAO_VENDER    = "Vender"
	BOTAO_VOLTAR    = "Voltar"
)

var (
	semEstilo       = lipgloss.NewStyle()
	semEstiloTabela = table.DefaultStyles()
	estiloTitulo    = lipgloss.NewStyle().Foreground(lipgloss.Color(COR_TITULO))
	estiloErro      = lipgloss.NewStyle().Foreground(lipgloss.Color(COR_ERRO))
	estiloFocado    = lipgloss.NewStyle().Foreground(lipgloss.Color(COR_FOCO))
	estiloCursor    = estiloFocado.Copy()
	estiloTabela    = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))

	botaoAdicionarFocado    = botao(aplicarFoco(BOTAO_ADICIONAR))
	botaoAdicionarNaoFocado = botao(BOTAO_ADICIONAR)
	botaoVenderFocado       = botao(aplicarFoco(BOTAO_VENDER))
	botaoVenderNaoFocado    = botao(BOTAO_VENDER)
	botaoVoltarFocado       = botao(aplicarFoco(BOTAO_VOLTAR))
	botaoVoltarNaoFocado    = botao(BOTAO_VOLTAR)
)

func aplicarFocoTabela(tabela *table.Model) {
	estilo := semEstiloTabela

	estilo.Header = estilo.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(COR_FUNDO)).
		BorderBottom(true).
		Bold(true)
	estilo.Selected = estilo.Selected.
		Foreground(lipgloss.Color(COR_ESCURA)).
		Background(lipgloss.Color(COR_FOCO)).
		Bold(true)

	tabela.SetStyles(estilo)
	tabela.Focus()
}

func removerFocoTabela(tabela *table.Model) {
	estilo := semEstiloTabela

	estilo.Header = estilo.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(COR_FUNDO)).
		BorderBottom(true).
		Bold(true)

	tabela.SetStyles(estilo)
	tabela.Blur()
}

func aplicarTitulo(s string) string {
	return estiloTitulo.Bold(true).Render(s)
}

func aplicarErro(s string) string {
	return estiloErro.Bold(true).Render(s)
}

func removerEstilo(s string) string {
	return semEstilo.Render(s)
}

func aplicarFoco(s string) string {
	return estiloFocado.Bold(true).Render(s)
}

func botao(s string) string {
	return fmt.Sprintf(removerEstilo("[ %s ]"), s)
}
