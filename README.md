# Cotação Dólar - Sistema Client/Server em Go

Este projeto consiste em dois aplicativos em Go: um servidor HTTP que fornece a cotação atual do dólar e registra no banco de dados SQLite, e um cliente que consome esse endpoint e salva a cotação em um arquivo de texto.

## 📁 Estrutura do Projeto

- `server.go`: Servidor HTTP que consome a API pública de câmbio e salva os dados.
- `client.go`: Cliente que consome o servidor e salva a cotação em um arquivo.
- `Dockerfile`: Define a imagem Docker para execução do servidor.
- `Makefile`: Automatiza tarefas comuns como build, execução e testes.
- `server_test.go`: Testes automatizados básicos.
- `cotacao.txt`: Gerado pelo cliente com a última cotação.
- `cotacoes.db`: Banco SQLite com histórico de cotações (gerado automaticamente pelo servidor).

## 🚀 Como Executar

### Pré-requisitos

- Go 1.21+
- Docker (opcional)
- `make` (opcional, recomendado)

### Usando Makefile

```bash
make build
make run-server
// Em outro terminal
make run-client
```

### Usando Docker

```bash
make docker-build
make docker-run
```

### Executando manualmente

```bash
go run server.go
# Em outro terminal
go run client.go
```

## 🧪 Testes

```bash
make test
```

## 📋 Funcionalidades

- Endpoint /cotacao na porta 8080.
- Consome a API: https://economia.awesomeapi.com.br/json/last/USD-BRL
- Usa context.WithTimeout:
  - 200ms para chamada externa.
  - 10ms para SQLite.
  - 300ms no cliente.
- Logs informam em caso de timeout.
- Cliente salva o valor em cotacao.txt.

## 📦 Exemplo de resposta

```json
{
  "bid": "5.1240"
}
```

## 📌 Observações

- Se qualquer contexto ultrapassar o tempo, é feito log da falha.
- O banco é criado automaticamente.
- Simples, modular e prático.
