package beautiful

import (
	"gerenciador/internal"
	tea "github.com/charmbracelet/bubbletea"
)

type RaizTelaModelo struct {
	modelo   tea.Model
}

func RaizTela(c *internal.Carteira) RaizTelaModelo {
	modeloInicial := MenuTela(c)

	return RaizTelaModelo{
		modelo: &modeloInicial,
	}
}

func (r RaizTelaModelo) Init() tea.Cmd {
	return r.modelo.Init() 
}

func (r RaizTelaModelo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r.modelo.Update(msg)
}

func (r RaizTelaModelo) View() string {
	return r.modelo.View()
}

func (r RaizTelaModelo) MudarTela(modelo tea.Model) (tea.Model, tea.Cmd) {
	r.modelo = modelo
	return r.modelo, r.modelo.Init()
}
