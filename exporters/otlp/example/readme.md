# Example of retries with OTLP exporter

## Run the server
The server will only respond with an Abort code.  This should be retried a number of times by the client.
It is working correctly if you see `Received request` 6 times.

```bash
go run ./server
Starting Server localhost:9000
Received request
...
```

## Run the client
The client is an example of a client with default configuration.  If working to spec it will retry 5 times.

```bash
go run ./client
2021/03/03 15:46:37 rpc error: code = Aborted desc = Error Message

# Messages from the server
go run ./server
Starting Server localhost:9000
Received request
```

## GRPC retry workaround
The golang grpc client has experimental support for retries.  To enable you have to set the environment variable `GRPC_GO_RETRY=on`

```shell
GRPC_GO_RETRY=on go run ./client
2021/03/03 15:46:37 rpc error: code = Aborted desc = Error Message

# Messages from the server
go run ./server
Starting Server localhost:9000
Received request
Received request
Received request
Received request
Received request
```