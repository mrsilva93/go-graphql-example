# Projeto Go-GraphQL

## Descrição
Este projeto implementa uma API GraphQL utilizando Go. Ele gerencia transações, tipos de transações e usuários, com integração a um banco de dados SQLite.

## Estrutura do Projeto

### Arquivos principais
- **`gqlgen.yml`**: Configuração do gerador de código GraphQL.
- **`tools.go`**: Dependências utilizadas para geração de código GraphQL.
- **`data.db`**: Banco de dados SQLite utilizado pela aplicação.

### Diretórios
- **`cmd/server`**: Contém o código principal do servidor.
  - `server.go`: Inicializa o servidor GraphQL e configura os resolvers.
  
- **`graph`**: Contém os arquivos relacionados ao GraphQL.
  - `generated.go`: Código gerado automaticamente pelo gqlgen.
  - `model/`: Modelos utilizados pelo GraphQL.
  - `resolvers/`: Implementações dos resolvers GraphQL.
  - `schema/`: Definições dos schemas GraphQL.

- **`internal/database`**: Contém a lógica de acesso ao banco de dados.
  - `transaction.go`: Gerencia transações.
  - `transaction_type.go`: Gerencia tipos de transações.
  - `users.go`: Gerencia usuários.

## Configuração e Execução

### Pré-requisitos
- Go instalado na máquina.
- SQLite para o banco de dados.

### Passos para executar
1. Instale as dependências:
   ```sh
   go mod tidy
   ```
2. Execute o servidor:
   ```sh
   go run cmd/server/server.go
   ```

### Variáveis de ambiente
- `PORT`: Define a porta do servidor. Padrão: `8080`.

## Funcionalidades

### GraphQL
- **Queries**:
  - Buscar todos os usuários, transações e tipos de transações.
  - Buscar por ID.
- **Mutations**:
  - Criar novos usuários, transações e tipos de transações.

### Banco de Dados
- Utiliza SQLite para persistência de dados.

## Exemplos de Uso

### Query para buscar todos os usuários
```graphql
query {
  users {
    id
    name
    cnpj
  }
}
```

### Mutation para criar um novo usuário
```graphql
mutation {
  createUser(name: "João", cnpj: "123456789") {
    id
    name
    cnpj
  }
}
```

```graphql
  transactions {
    id
    value
    date
    transactionType{
      id
      name
    }
    user {
      id
      name
      cnpj
    }
  }
}


mutation createTransaction {
  createTransaction(input:{transactionTypeId: "3a76c424-e992-403c-98e5-fbf4011c450a", userId: "f0dfbe9b-9d36-4bbe-95d3-d73e7575e6ae",value: 100, date: "2025-05-28",
  }){
    id
    value
    date
    user{
      id
      name
    }
    transactionType{
      name
      description
    }
  }
}


query getTransactionTYpes{
  transactionTypes{
    id
    name
    description
    transactions{
      id
      value
      date
    }
  }
}

query naruto{
  transactions(filter: {  maxValue:100, userId:1, limit:2 }) {
    id
    value
    user {
      name
    }
  }
}


query getUsers {
  users {
  name 
  cnpj
}
```

## Contribuição
Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença
Este projeto está licenciado sob [MIT License](LICENSE).