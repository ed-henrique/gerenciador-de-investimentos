# Terminal de Gerenciamento de Investimentos

Ferramenta CLI de gerenciamento de investimentos.

## Como Executar

Execute o comando abaixo:

```bash
docker run -it --rm eduhenfm/terminal-de-gerenciamento-de-investimentos:latest
```

> [!IMPORTANT]
> É **extremamente recomendado** que você utilize o [Docker](https://www.docker.com/) para executar a aplicação, ao invés de executá-la pelo código-fonte.

## Roadmap

### Funcionalidades

- [x] Adicionar ativo
- [x] Vender ativo
- [x] Listar ativos
- [x] Consultar relatório de ativos
- [x] Mostrar menu de escolhas para o usuário

### Extra

- [ ] Persistir dados entre sessões em arquivo `.csv`
- [x] Conteinerizar a aplicação

## Tecnologias Utilizadas

- [Go](https://go.dev/) (1.22.2)

## Práticas Utilizadas

- TDD

> 1. Criar teste falho
> 2. Fazer o teste passar
> 3. Refatorar

> [!NOTE]
> Para executar os testes, execute o comando `go test ./tests`.
