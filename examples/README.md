# Examples

All commands are supposed to be executed from root of this repository.

## broadcast

Incoming messages are broadcasted to all active clients.

```sh
go run ./examples/broadcast --host localhost --port 8080

# open client with
nc localhost 8080
```

## chat

A TCP-based chat is started. For more check [here](chat/).

```sh
go run ./examples/chat --host localhost --port 8080

# open client with
nc localhost 8080
```

## echo

Incoming messages are echoed.

```sh
go run ./examples/echo --host localhost --port 8080
```

## http simple

Getting started with HTTP/1.1 🚀. Check out the [intro](httpsimple/).

Any request is responded with 200 status code and a simple body.

```sh
go run ./examples/httpsimple --host localhost --port 8080
```
