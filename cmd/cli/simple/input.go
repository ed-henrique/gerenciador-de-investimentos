package simple

import (
	"fmt"
	"gerenciador/internal"
	"strconv"
	"time"
)

func AcaoDoUsuario() (acao int) {
	fmt.Scanf("%d", &acao)
	fmt.Println()
	return
}

func AdicionarAtivo(carteira *internal.Carteira) {
	var linha string
	ativo := internal.Ativo{}

	fmt.Printf("-------------------\n| Adicionar Ativo |\n-------------------\n\n")

	fmt.Print("Codigo: ")
	fmt.Scanf("%s", &ativo.Codigo)

	fmt.Print("Quantidade: ")
	_, err := fmt.Scanln(&linha)

	quantidade, err := strconv.ParseInt(linha, 10, 0)
	ativo.Quantidade = int(quantidade)

	if err != nil {
		fmt.Printf("\nErro: A quantidade inserida deve ser um número inteiro\n")
		return
	}

	fmt.Print("Valor Unitario: ")
	_, err = fmt.Scanln(&linha)

	ativo.ValorUnitario, err = strconv.ParseFloat(linha, 64)

	if err != nil {
		fmt.Printf("\nErro: O valor unitário inserido deve ser um número inteiro ou de ponto flutuante\n")
		return
	}

	ativo.DataDaCompra = time.Now()

	err = carteira.AdicionarAtivo(ativo)

	if err != nil {
		fmt.Printf("\nErro: %s\n", err.Error())
	}
}

func VenderAtivo(carteira *internal.Carteira) {
	var (
		linha      string
		codigo     string
		quantidade int64
	)

	fmt.Printf("-----------------\n| Vender Ativo |\n-----------------\n\n")

	fmt.Print("Codigo: ")
	fmt.Scanf("%s", &codigo)

	fmt.Print("Quantidade: ")
	_, err := fmt.Scanln(&linha)

	quantidade, err = strconv.ParseInt(linha, 10, 0)

	if err != nil {
		fmt.Printf("\nErro: A quantidade inserida deve ser um número inteiro\n")
		return
	}

	err = carteira.VenderAtivo(codigo, int(quantidade))

	if err != nil {
		fmt.Printf("\nErro: %s\n", err.Error())
	}
}
