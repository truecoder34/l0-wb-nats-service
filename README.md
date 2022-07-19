# l0-wb-nats-service





```
> go mod init github.com/truecoder34/l0-wb-nats-service
> go get github.com/nats-io/nats-streaming-server
> go get github.com/nats-io/go-nats-streaming
```


## NATS USAGE <br/>
NATS Straming vased on NATS (gnatsd) and provides an extra capability of having a persist logs to be used for event streaming systems.
https://github.com/nats-io/stan.go/blob/main/examples/

1. Run Server <br/>
```go run nats-streaming-server.go```

2. Run Sender <br/>
```go run nats-streaming-server.go```

3. Run Consumer <br/>
```go run nats-streaming-server.go```




