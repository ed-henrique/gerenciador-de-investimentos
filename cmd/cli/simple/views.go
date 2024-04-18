package simple

import (
	"fmt"
	"gerenciador/internal"
	"strings"
)

func Menu() {
  var sb strings.Builder

  sb.WriteString("Terminal de Gerenciamento de Investimentos\n\n")
  sb.WriteString("1 - Adicionar Ativo\n")
  sb.WriteString("2 - Vender Ativo\n")
  sb.WriteString("3 - Listar Ativos\n")
  sb.WriteString("4 - Gerar Relatorio\n")
  sb.WriteString("5 - Sair\n")
  sb.WriteString("\nDigite a ação que você quer tomar: ")

  fmt.Print(sb.String())
}

func ListarAtivos(carteira *internal.Carteira) {
  fmt.Print(carteira.ImprimirAtivos())
}

func GerarRelatorio(carteira *internal.Carteira) {
  fmt.Print(carteira.ImprimirResumoAtivos())
}
