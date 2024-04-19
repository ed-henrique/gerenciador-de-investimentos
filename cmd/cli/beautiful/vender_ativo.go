package beautiful

import (
	"fmt"
	"gerenciador/internal"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type VenderAtivoTelaModelo struct {
	focusIndex int
	inputs     []textinput.Model
	carteira   *internal.Carteira
	erro       string
}

func VenderAtivoTela(c *internal.Carteira) VenderAtivoTelaModelo {
	a := VenderAtivoTelaModelo{
		carteira: c,
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range a.inputs {
		t = textinput.New()
		t.Cursor.Style = estiloCursor
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Codigo"
			t.Focus()
			t.PromptStyle = estiloFocado
			t.TextStyle = estiloFocado
			t.CharLimit = 64
		case 1:
			t.Placeholder = "Quantidade"
			t.CharLimit = 64
		}

		a.inputs[i] = t
	}

	return a
}

func (a *VenderAtivoTelaModelo) Init() tea.Cmd {
	return textinput.Blink
}

func (a *VenderAtivoTelaModelo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return a, tea.Quit
		case "esc":
			menuTela := MenuTela(a.carteira)
			return RaizTela(a.carteira).MudarTela(&menuTela)
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && a.focusIndex == len(a.inputs) {
				codigo := a.inputs[0].Value()
				quantidade, err := strconv.Atoi(a.inputs[1].Value())

				if err != nil {
					a.erro = "Erro: A quantidade inserida deve ser um número inteiro."
					return a, nil
				}

				err = a.carteira.VenderAtivo(codigo, quantidade)

				if err != nil {
					a.erro = "Erro: " + err.Error()
					return a, nil
				} else {
					a.erro = ""
					menuTela := MenuTela(a.carteira)
					return RaizTela(a.carteira).MudarTela(&menuTela)
				}
			}

			if s == "enter" && a.focusIndex == len(a.inputs) + 1 {
				menuTela := MenuTela(a.carteira)
				return RaizTela(a.carteira).MudarTela(&menuTela)
			}

			if s == "up" || s == "shift+tab" {
				a.focusIndex--
			} else {
				a.focusIndex++
			}

			if a.focusIndex > len(a.inputs) + 1 {
				a.focusIndex = 0
			} else if a.focusIndex < 0 {
				a.focusIndex = len(a.inputs) + 1
			}

			cmds := make([]tea.Cmd, len(a.inputs))
			for i := 0; i <= len(a.inputs)-1; i++ {
				if i == a.focusIndex {
					cmds[i] = a.inputs[i].Focus()
					a.inputs[i].PromptStyle = estiloFocado
					a.inputs[i].TextStyle = estiloFocado
					continue
				}

				a.inputs[i].Blur()
				a.inputs[i].PromptStyle = semEstilo
				a.inputs[i].TextStyle = semEstilo
			}

			return a, tea.Batch(cmds...)
		}
	}

	cmd := a.updateInputs(msg)

	return a, cmd
}

func (a *VenderAtivoTelaModelo) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(a.inputs))

	for i := range a.inputs {
		a.inputs[i], cmds[i] = a.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (a *VenderAtivoTelaModelo) View() string {
	var s strings.Builder

	fmt.Fprintf(&s, "%s\n\n", aplicarTitulo("Vender Ativo"))

	if a.erro != "" {
		fmt.Fprintf(&s, "%s\n\n", aplicarErro(a.erro))
	}

	for i := range a.inputs {
		s.WriteString(a.inputs[i].View())
		if i < len(a.inputs)-1 {
			s.WriteRune('\n')
		}
	}

	botaoVender := &botaoVenderNaoFocado
	botaoVoltar := &botaoVoltarNaoFocado

	if a.focusIndex == len(a.inputs) {
		botaoVender = &botaoVenderFocado
	}
	if a.focusIndex == len(a.inputs) + 1 {
		botaoVoltar = &botaoVoltarFocado
	}

	fmt.Fprintf(&s, "\n\n%s\n%s\n\n", *botaoVender, *botaoVoltar)

	return s.String()
}
