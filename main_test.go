package main

import (
	"testing"
)

func TestAdicionarAtivo(t *testing.T) {
  c := Carteira{}

  got := c.AdicionarAtivo(Ativo{
    Codigo: "STNE",
    Quantidade: 10,
    ValorUnitario: 9.99,
  })

  var expected error = nil

  if got != expected {
    t.Fatalf("got=%q expected=%q", got, expected)
  }
}

