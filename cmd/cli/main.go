package main

import (
	"flag"
	"fmt"
	"gerenciador/cmd/cli/beautiful"
	"gerenciador/cmd/cli/simple"
	"gerenciador/internal"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var bonito bool
	flag.BoolVar(&bonito, "b", false, "ativa o modo bonito")
	flag.Parse()

	carteira := &internal.Carteira{}

	if bonito {
		if _, err := tea.NewProgram(beautiful.RaizTela(carteira), tea.WithAltScreen()).Run(); err != nil {
			fmt.Printf("Não foi possível iniciar o programa: %s\n", err)
			os.Exit(1)
		}
	} else {
	loop:
		for {
			simple.Menu()
			acao := simple.AcaoDoUsuario()

			switch acao {
			case 1:
				simple.AdicionarAtivo(carteira)
			case 2:
				simple.VenderAtivo(carteira)
			case 3:
				simple.ListarAtivos(carteira)
			case 4:
				simple.GerarRelatorio(carteira)
			case 5:
				fmt.Println("Até logo")
				break loop
			default:
				continue
			}

			fmt.Println()
		}
	}
}
