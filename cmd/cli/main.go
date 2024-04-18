package main

import (
	"fmt"
	"gerenciador/cmd/cli/simple"
	"gerenciador/internal"
)

func main() {
  carteira := &internal.Carteira{}

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
