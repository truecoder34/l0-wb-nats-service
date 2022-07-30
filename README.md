# l0-wb-nats-service

## API METHODS:
<b>[GET]</b> - get all transactions from DB <br/>
```
curl --location --request GET 'http://localhost:8080/transactions'
```
<b>[GET]</b> - get transaction by ID from DB <br/>
```
curl --location --request GET 'http://localhost:8080/transaction/63473674-95dc-4d81-9e88-41eb932b863f'
```

<b>[GET]</b> - get all transactions from cache <br/>
```
curl --location --request GET 'http://localhost:8080/transactions-cache'
```

<b>[GET]</b> - get transaction by ID from from cache <br/> 
```
curl --location --request GET 'http://localhost:8080/transaction-cache/99e16d0d-2bbd-4c6c-8667-1e0e1b7db52f'
```



### USED PACKAGES <br/>
```
> go mod init github.com/truecoder34/l0-wb-nats-service
> go get github.com/nats-io/nats-streaming-server
> go get github.com/nats-io/go-nats-streaming
> go get gorm.io/gorm
> go get github.com/satori/go.uuid
> go get github.com/joho/godotenv
> go get "gorm.io/driver/postgres"
> go get "github.com/gorilla/mux"
> go get "github.com/gin-gonic/gin" 
> go get "github.com/gorilla/mux"
```

### NATS USAGE <br/>
NATS Straming vased on NATS (gnatsd) and provides an extra capability of having a persist logs to be used for event streaming systems.
https://github.com/nats-io/stan.go/blob/main/examples/

1. Run Server <br/>
```go run nats-streaming-server.go```

2. Run Sender <br/>
```go run nats-streaming-server.go```

3. Run Consumer <br/>
```go run nats-streaming-server.go```
