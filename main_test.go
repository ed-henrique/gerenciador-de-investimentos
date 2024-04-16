package main

import (
	"testing"
)

func TestAdicionarAtivo(t *testing.T) {
  got := AdicionarAtivo(Ativo{
    Codigo: "STNE",
    Quantidade: 10,
    ValorUnitario: 9.99,
  })

  var expected error = nil

  if got != expected {
    t.Fatalf("got=%q expected=%q", got, expected)
  }
}

