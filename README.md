# From TCP to HTTP

I created this repository to learn more about the HTTP protocol and protocols in
general.

> A **protocol** is a system of rules that define how data is exchanged within
> or between computers. Communications between devices require that the devices
> agree on the format of the data that is being exchanged. The set of rules that
> defines a format is called a protocol.

A sever has been implemented that accepts TCP connections and passes them on to
any custom protocol that statisfies the protocol interface.

## Examples

Check out the [examples/](examples/) directory for implementations of the
protocol interface.

## Resources

- [RFC 7231 HTTP/1.1](https://datatracker.ietf.org/doc/html/rfc7231)
- [MDN Web Docs HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Overview)
- [Golang net/http](https://cs.opensource.google/go/go/+/master:src/net/http/request.go;l=113)
