# Desafio - Clean Architecture
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.
Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

## Instruções de como rodar o desafio:
Primeiramente devemos clonar o repositorio

```bash
git clone git@github.com:mbrocco/desafio-clean-arq.git
```

## Acessar o diretório do projeto clonado:
```bash
cd desafio-clean-arq
```

## Rodar o comando para subir o docker e instalar as dependências necessárias para rodar o mysql e o RabbitMQ:
```bash
docker-compose up --build -d
```

## Executar o projeto GO:
```bash
go mod tidy
cd cmd/ordersystem
go run main.go wire_gen.go
```

## Informações de logs para o acompanhamento do projeto:
```bash
# Starting web server on port :8000
# Starting gRPC server on port 50051
# Starting GraphQL server on port 8080
```

## Como utilizar as apis:
Foram desenvolvidas duas APIs e foram utilizados três tipos de protocolos(REST,GRAPHQL e GRPC)

### REST

Acessar o diretório api para rodar via REST as apis

#### API - Criar uma nova ordem
```bash
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "id":"iiii",
    "price": 100.5,
    "tax": 0.5
}
```

#### API - Listar as ordens criadas
```bash
GET http://localhost:8000/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

### GRAPHQL
Acessar o navegador e acionar a rota http://localhost:8080

#### API - Criar uma nova ordem
```graphql
mutation createOrder {
  createOrder(input: {id:"xxx", Price: 200.0, Tax: 3.0}) {
  	id
  	Price
		Tax
  	FinalPrice
  }
}
```

#### API - Listar as ordens criadas
```graphql
query listOrders {
  listOrders {
    id
    Price
		Tax
    FinalPrice
  }
}
```

### GRPC
Utilizar o evans: https://github.com/ktr0731/evans, e execute os comandos a seguir:

```bash
evans -r repl
```

```bash
package pb
```

```bash
service OrderService
```

#### API - Criar uma nova ordem
```bash
call CreateOrder
```

#### API - Listar as ordens criadas
```bash
call ListOrders
```








