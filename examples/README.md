# Examples

## broadcastecho

A TCP server is started that runs a read loop and broadcasts all incomming
messages.

```sh
go run ./examples/echo --host localhost --port 8080

# open clients with
nc localhost 8080
```

## chat

ChatProtocol

## echo

A TCP server is started that runs a read loop and echos all incomming messages.

```sh
go run ./examples/echo --host localhost --port 8080
```

## http

HTTP Protocol
