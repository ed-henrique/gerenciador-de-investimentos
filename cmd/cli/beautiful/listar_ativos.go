package beautiful

import (
	"fmt"
	"gerenciador/internal"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ImprimirAtivosTelaModelo struct {
	tabela       table.Model
	tabelaAtivos table.Model
	carteira     *internal.Carteira
}

func ImprimirAtivosTela(c *internal.Carteira) ImprimirAtivosTelaModelo {
	var (
		linhas          []table.Row
		linhasAtivos    []table.Row
	)

	for codigo, ativos := range *c {
		linhas = append(linhas, table.Row{codigo, fmt.Sprint(internal.QuantidadeAtivos(ativos))})
	}

	colunas := []table.Column{
		{Title: "Codigo", Width: 10},
		{Title: "Quantidade Total", Width: 16},
	}

	colunasAtivos := []table.Column{
		{Title: "Codigo", Width: 10},
		{Title: "Quantidade", Width: 10},
		{Title: "Valor Unitario", Width: 14},
		{Title: "Data da Compra", Width: 10},
	}

	tabela := table.New(
		table.WithColumns(colunas),
		table.WithRows(linhas),
		table.WithHeight(10),
		table.WithFocused(true),
	)

	aplicarFocoTabela(&tabela)

	for _, ativos := range *c {
		for _, a := range ativos {
			linhasAtivos = append(linhasAtivos, table.Row{
				a.Codigo,
				fmt.Sprint(a.Quantidade),
				fmt.Sprint(a.ValorUnitario),
				a.DataDaCompra.Format("2006-01-02"),
			})
		}
	}

	tabelaAtivos := table.New(
		table.WithColumns(colunasAtivos),
		table.WithRows(linhasAtivos),
		table.WithHeight(10),
		table.WithFocused(false),
	)

	removerFocoTabela(&tabelaAtivos)

	a := ImprimirAtivosTelaModelo{
		tabela:          tabela,
		tabelaAtivos:   tabelaAtivos,
		carteira:        c,
	}

	return a
}

func (i *ImprimirAtivosTelaModelo) Init() tea.Cmd {
	return nil
}

func (i *ImprimirAtivosTelaModelo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab":
			if i.tabela.Focused() {
				removerFocoTabela(&i.tabela)
				aplicarFocoTabela(&i.tabelaAtivos)
			} else {
				aplicarFocoTabela(&i.tabela)
				removerFocoTabela(&i.tabelaAtivos)
			}
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

func (i *ImprimirAtivosTelaModelo) View() string {
	var s strings.Builder

	fmt.Fprintf(&s, "%s\n\n", aplicarTitulo("Lista de Ativos"))
	fmt.Fprintf(&s, "%s\n\n", estiloTabela.Render(i.tabela.View()))
	fmt.Fprintf(&s, "%s\n\n", estiloTabela.Render(i.tabelaAtivos.View()))
	fmt.Fprintf(&s, "%s\n", lipgloss.NewStyle().Foreground(lipgloss.Color(COR_FUNDO)).Render("Aperte 'tab' ou 'shift+tab' para mudar de tabela"))
	fmt.Fprintf(&s, "%s\n", lipgloss.NewStyle().Foreground(lipgloss.Color(COR_FUNDO)).Render("Aperte 'q', 'enter' ou 'esc' para voltar ao menu"))

	return s.String()
}
