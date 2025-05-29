# Variáveis
APP_NAME = go-graphql
SERVER_FILE = cmd/server/server.go
PORT ?= 8080

# Comandos
.PHONY: all
all: build

.PHONY: build
build:
    @echo "Compilando o projeto..."
    go build -o $(APP_NAME) $(SERVER_FILE)

.PHONY: run
run:
    @echo "Executando o servidor na porta $(PORT)..."
    PORT=$(PORT) go run $(SERVER_FILE)

.PHONY: tidy
tidy:
    @echo "Instalando dependências..."
    go mod tidy

.PHONY: clean
clean:
    @echo "Limpando arquivos gerados..."
    rm -f $(APP_NAME)

.PHONY: test
test:
    @echo "Executando testes..."
    go test ./...

.PHONY: gqlgen
gqlgen:
    @echo "Gerando código GraphQL..."
    go run github.com/99designs/gqlgen generate

.PHONY: help
help:
    @echo "Comandos disponíveis:"
    @echo "  make build   - Compila o projeto"
    @echo "  make run     - Executa o servidor"
    @echo "  make tidy    - Instala dependências"
    @echo "  make clean   - Remove arquivos gerados"
    @echo "  make test    - Executa testes"
    @echo "  make gqlgen  - Gera código GraphQL"