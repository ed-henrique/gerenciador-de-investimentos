package main

import (
	"testing"
	"time"
)

func TestAdicionarAtivo(t *testing.T) {
  c := Carteira{}

  got := c.AdicionarAtivo(Ativo{
    Codigo: "STNE",
    Quantidade: 10,
    ValorUnitario: 9.99,
    DataDaCompra: time.Now(),
  })

  var expected error = nil

  if got != expected {
    t.Fatalf("got=%q expected=%q", got, expected)
  }
}

func TestVenderAtivo(t *testing.T) {
  t.Run("Vender uma unidade do ativo (quantidade inicial = 2)", func(t *testing.T) {
    c := Carteira{
      "STNE": []Ativo{
	{ Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.99, DataDaCompra: time.Now() },
	{ Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.98, DataDaCompra: time.Now() },
      },
    }

    got := c.VenderAtivo("STNE", 1)

    var expected error = nil

    if got != expected {
      t.Fatalf("got=%q expected=%q", got, expected)
    }
  })

  t.Run("Vender três unidades do ativo (quantidade inicial = 2)", func(t *testing.T) {
    c := Carteira{
      "STNE": []Ativo{
	{ Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.99, DataDaCompra: time.Now() },
	{ Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.98, DataDaCompra: time.Now() },
      },
    }

    got := c.VenderAtivo("STNE", 3)

    expected := ErrQuantidadeDeAtivosParaVendaExtrapolada

    if got != expected {
      t.Fatalf("got=%q expected=%q", got, expected)
    }
  })
}

func TestImprimirAtivos(t *testing.T) {
  c := Carteira{
    "STNE": []Ativo{
      { Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.99, DataDaCompra: time.Date(2024, time.April, 16, 0, 0, 0, 0, time.Local) },
      { Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.98, DataDaCompra: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.Local) },
    },
  }

  got := c.ImprimirAtivos()

  // Espera-se um resultado como esse:
  //
  // Ativos
  // -------------------------------------------
  // |          Codigo |      Quantidade Total |
  // -------------------------------------------
  // |            STNE |                     2 |
  // -------------------------------------------
  //
  // STNE
  // ----------------------------------------
  // |  Valor Unitario |     Data da Compra |
  // ----------------------------------------
  // |            9.99 |         2024-04-16 |
  // ----------------------------------------
  // |            9.98 |         2024-04-15 |
  // ----------------------------------------
  expected := `Ativos
-------------------------------------------
|          Codigo |      Quantidade Total |
-------------------------------------------
|            STNE |                     2 |
-------------------------------------------

STNE
-------------------------------------------
|  Valor Unitario |        Data da Compra |
-------------------------------------------
|            9.99 |            2024-04-16 |
-------------------------------------------
|            9.98 |            2024-04-15 |
-------------------------------------------
`

  if got != expected {
    t.Fatalf("got=%q expected=%q", got, expected)
  }
}

func TestImprimirResumoAtivos(t *testing.T) {
  c := Carteira{
    "STNE": []Ativo{
      { Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.99, DataDaCompra: time.Date(2024, time.April, 16, 0, 0, 0, 0, time.Local) },
      { Codigo: "STNE", Quantidade: 1, ValorUnitario: 9.98, DataDaCompra: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.Local) },
    },
    "GOOG": []Ativo{
      { Codigo: "GOOG", Quantidade: 1, ValorUnitario: 9.99, DataDaCompra: time.Date(2024, time.April, 16, 0, 0, 0, 0, time.Local) },
      { Codigo: "GOOG", Quantidade: 1, ValorUnitario: 9.98, DataDaCompra: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.Local) },
    },
  }

  got := c.ImprimirResumoAtivos()

  // Espera-se um resultado como esse:
  //
  // Carteira
  // --------------------------------------------
  // |  Quantidade de Ativos |      Valor Total |
  // --------------------------------------------
  // |                     4 |            39.94 |
  // --------------------------------------------
  //
  // Ativos
  // --------------------------------------------------------------
  // |          Codigo |      Quantidade Total |      Valor Total |
  // --------------------------------------------------------------
  // |            STNE |                     2 |            19.97 |
  // --------------------------------------------------------------
  // |            GOOG |                     2 |            19.97 |
  // --------------------------------------------------------------
  expected := `Carteira
--------------------------------------------
|  Quantidade de Ativos |      Valor Total |
--------------------------------------------
|                     4 |            39.94 |
--------------------------------------------

Ativos
--------------------------------------------------------------
|          Codigo |      Quantidade Total |      Valor Total |
--------------------------------------------------------------
|            STNE |                     2 |            19.97 |
--------------------------------------------------------------
|            GOOG |                     2 |            19.97 |
--------------------------------------------------------------
`

  if got != expected {
    t.Fatalf("got=%q expected=%q", got, expected)
  }
}

