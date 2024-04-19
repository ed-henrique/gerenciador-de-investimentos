package beautiful

import (
	"fmt"
	"gerenciador/internal"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GerarRelatorioTelaModelo struct {
	tabela       table.Model
	tabelaAtivos table.Model
	carteira     *internal.Carteira
}

func GerarRelatorioTela(c *internal.Carteira) GerarRelatorioTelaModelo {
	var (
		linhasAtivos       []table.Row
		quantidadeDeAtivos int
		valorTotal         float64
	)

	for codigo, ativos := range *c {
		linha := table.Row{codigo, fmt.Sprint(internal.QuantidadeAtivos(ativos))}

		var valorTotalAtivos float64
		for _, a := range ativos {
			quantidadeDeAtivos += a.Quantidade
			valorTotal += a.ValorUnitario * float64(a.Quantidade)
			valorTotalAtivos += a.ValorUnitario * float64(a.Quantidade)
		}

		linha = append(linha, fmt.Sprintf("%.2f", valorTotalAtivos))
		linhasAtivos = append(linhasAtivos, linha)
	}

	colunas := []table.Column{
		{Title: "Quantidade de Ativos", Width: 20},
		{Title: "Valor Total", Width: 11},
	}
	linhas := table.Row{fmt.Sprint(quantidadeDeAtivos), fmt.Sprintf("%.2f", valorTotal)}

	tabela := table.New(
		table.WithColumns(colunas),
		table.WithRows([]table.Row{linhas}),
		table.WithHeight(1),
		table.WithFocused(false),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		AlignHorizontal(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(COR_FUNDO)).
		BorderBottom(true).
		Bold(true)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color(COR_FOCO)).
		AlignHorizontal(lipgloss.Center).
		Bold(false)
	tabela.SetStyles(s)

	colunasAtivos := []table.Column{
		{Title: "Codigo", Width: 10},
		{Title: "Quantidade Total", Width: 16},
		{Title: "Valor Total", Width: 11},
	}

	tabelaAtivos := table.New(
		table.WithColumns(colunasAtivos),
		table.WithRows(linhasAtivos),
		table.WithHeight(10),
		table.WithFocused(true),
	)

	sAtivos := table.DefaultStyles()
	sAtivos.Header = sAtivos.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(COR_FUNDO)).
		BorderBottom(true).
		Bold(true)
	sAtivos.Selected = sAtivos.Selected.
		Foreground(lipgloss.Color(COR_ESCURA)).
		Background(lipgloss.Color(COR_FOCO)).
		Bold(true)
	tabelaAtivos.SetStyles(sAtivos)

	a := GerarRelatorioTelaModelo{
		tabela:       tabela,
		tabelaAtivos: tabelaAtivos,
		carteira:     c,
	}

	return a
}

func (i *GerarRelatorioTelaModelo) Init() tea.Cmd {
	return nil
}

func (i *GerarRelatorioTelaModelo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return i, tea.Quit
		case "enter", "esc", "q":
			menuTela := MenuTela(i.carteira)
			return RaizTela(i.carteira).MudarTela(&menuTela)
		}
	}

	var (
		cmd tea.Cmd
		cmdAtivos tea.Cmd
	)

	i.tabela, cmd = i.tabela.Update(msg)
	i.tabelaAtivos, cmdAtivos = i.tabelaAtivos.Update(msg)

	return i, tea.Batch(cmd, cmdAtivos)
}

func (i *GerarRelatorioTelaModelo) View() string {
	var s strings.Builder

	fmt.Fprintf(&s, "%s\n\n", aplicarTitulo("Relatório de Ativos"))
	fmt.Fprintf(&s, "%s\n\n", estiloTabela.Render(i.tabela.View()))
	fmt.Fprintf(&s, "%s\n\n", estiloTabela.Render(i.tabelaAtivos.View()))
	fmt.Fprintf(&s, "%s\n", lipgloss.NewStyle().Foreground(lipgloss.Color(COR_FUNDO)).Render("Aperte 'q', 'enter' ou 'esc' para voltar ao menu"))

	return s.String()
}
