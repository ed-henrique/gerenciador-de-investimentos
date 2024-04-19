package beautiful

import (
	"fmt"
	"gerenciador/internal"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Adicionar         key.Binding
	Vender            key.Binding
	ImprimirAtivos    key.Binding
	ImprimirRelatorio key.Binding
	Ajuda             key.Binding
	Sair              key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Ajuda, k.Sair}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Adicionar, k.Vender, k.ImprimirAtivos, k.ImprimirRelatorio},
		{k.Ajuda, k.Sair},
	}
}

var keys = keyMap{
	Adicionar: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "adicionar ativo"),
	),
	Vender: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "vender ativo"),
	),
	ImprimirAtivos: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "listar ativos"),
	),
	ImprimirRelatorio: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "gerar relatório"),
	),
	Ajuda: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "mostrar menu de ajuda"),
	),
	Sair: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "sair"),
	),
}

type MenuTelaModelo struct {
	teclas   keyMap
	ajuda    help.Model
	carteira *internal.Carteira
	saindo   bool
}

func MenuTela(c *internal.Carteira) MenuTelaModelo {
	return MenuTelaModelo{
		teclas:   keys,
		ajuda:    help.New(),
		carteira: c,
	}
}

func (m *MenuTelaModelo) Init() tea.Cmd {
	return nil
}

func (m *MenuTelaModelo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ajuda.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.teclas.Adicionar):
			adicionarAtivoTela := AdicionarAtivoTela(m.carteira)
			return RaizTela(m.carteira).MudarTela(&adicionarAtivoTela)
		case key.Matches(msg, m.teclas.ImprimirAtivos):
			imprimirAtivosTela := ImprimirAtivosTela(m.carteira)
			return RaizTela(m.carteira).MudarTela(&imprimirAtivosTela)
		case key.Matches(msg, m.teclas.Vender):
			venderAtivoTela := VenderAtivoTela(m.carteira)
			return RaizTela(m.carteira).MudarTela(&venderAtivoTela)
		case key.Matches(msg, m.teclas.ImprimirRelatorio):
			gerarRelatorioTela := GerarRelatorioTela(m.carteira)
			return RaizTela(m.carteira).MudarTela(&gerarRelatorioTela)
		case key.Matches(msg, m.teclas.Ajuda):
			m.ajuda.ShowAll = !m.ajuda.ShowAll
		case key.Matches(msg, m.teclas.Sair):
			m.saindo = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *MenuTelaModelo) View() string {
	if m.saindo {
		return "Até logo!\n"
	}

	var s strings.Builder

	fmt.Fprintf(&s, "%s\n\n", aplicarTitulo("Menu"))
	fmt.Fprintf(&s, "%s\n", removerEstilo("Escolha uma operação (Aperte '?' para ter ajuda)"))

	helpView := m.ajuda.View(m.teclas)
	height := 3 - strings.Count(helpView, "\n")

	fmt.Fprintf(&s, "\n%s%s", strings.Repeat("\n", height), helpView)

	return s.String()
}
