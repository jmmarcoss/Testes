# Testes - Engenharia da Computação 2024.2

Este repositório contém implementações e testes relacionados a atividades passadas ao longo da disciplina de `Teste de Software` no IFPB - Campina Grande.

## Estrutura do Projeto

```
testes/
├── models/
│   ├── biblioteca/
│   │   ├── biblioteca.go          # Implementação da classe Biblioteca
│   │   └── biblioteca_test.go     # Testes unitários para a classe Biblioteca
│   ├── calculadora/
│   │   ├── calculadora.go         # Implementação da funcionalidade de cálculo
│   │   └── calculadora_test.go    # Testes unitários para a funcionalidade de cálculo
│   ├── livro/
│   │   ├── livro.go               # Implementação da classe Livro
│   │   └── livro_test.go          # Testes unitários para a classe Livro
│   ├── triangulo/
│   │   ├── triangle.go            # Implementação da classe Triangulo
│   │   └── triangle_test.go       # Testes unitários para a classe Triangulo
│   └── oficina/
│       ├── mecanico.go            # Implementação da classe Mecanico
│       ├── mecanico_test.go       # Testes unitários para a classe Mecanico
│       ├── servico.go             # Implementação da classe Servico
│       ├── servico_test.go        # Testes unitários para a classe Servico
│       ├── veiculo.go             # Implementação da classe Veiculo
│       └── veiculo_test.go        # Testes unitários para a classe Veiculo
├── main.go                        # Ponto de entrada principal do projeto
├── go.mod                         # Arquivo de gerenciamento de dependências do Go
└── README.md                      # Documentação do projeto
```

## Pré-requisitos

1. **Golang instalado**: Certifique-se de ter o [Go](https://golang.org/) instalado em sua máquina.
    - Verifique a instalação com:
      ```bash
      go version
      ```

2. **Configuração do ambiente**:
    - Inicie o módulo do Go se ainda não o fez:
      ```bash
      go mod init testes
      ```
    - Instale dependências adicionais caso necessário.

## Como Rodar os Testes

Os testes unitários estão disponíveis para as funcionalidades em cada módulo.

1. **Navegue para o diretório do projeto**:
   ```bash
   cd /caminho/para/testes
    ```
2. **Execute os testes para todos os pacotes:**
    ```bash
   go test ./...
   ```
3. **Execute os testes para um pacote específico:**
- Exemplo: Para rodar os testes do módulo `biblioteca`:
     ```bash
    go test ./models/biblioteca
    ```
4. **Execute os testes com saída detalhada:**
    ```bash
    go test ./... -v
    ```

## Exemplo de Comando para Teste

Para rodar os testes do método `PatrimonioHistorico` na `Biblioteca`:

```bash
go test ./models/biblioteca -run TestPatrimonioHistorico
```
