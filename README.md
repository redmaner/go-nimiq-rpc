# go-nimiq-rpc

[![Report card](https://goreportcard.com/badge/github.com/redmaner/go-nimiq-rpc)](https://goreportcard.com/report/github.com/redmaner/go-nimiq-rpc)
[![](https://godoc.org/github.com/redmaner/go-nimiq-rpc?status.svg)](https://godoc.org/github.com/redmaner/go-nimiq-rpc)

A Nimiq RPC client library in GO. This client library uses the [JSON-RPC protocol](https://www.jsonrpc.org/specification) and implements the [Nimiq RPC specification](https://github.com/nimiq/core-js/wiki/JSON-RPC-API#remotejs-client).

### What is Nimiq?
A blockchain technology inspired by Bitcoin but designed to run in your browser. It is money by nature but capable to do much more.
See [the Nimiq website](https://nimiq.com) for more information.

### How to use this library?
<code>
// Initialise a new client<br>
nimiqClient := NewClient("address.to.nimiqnode.com")
<br><br>
// Do an RPC call. For example retrieve the balance of a Nimiq account:<br>
balance, err := nimiqClient.GetBalance("NQ52 V4BF 52J3 0PM6 BG4M 9QY1 RUYS UAL6 CJD2")<br>
if err != nil {<br>
// Do your error handling here<br>
}<br>
<br>
// Do something with the response. In this case we print the balance<Br>
fmt.Printf("Balance: %v\n", balance)<br></code>

For more examples, see the example folder of this repository.

### Documentation
See the full documentation on [GoDoc](https://godoc.org/github.com/redmaner/go-nimiq-rpc)

### Contributing
Questions or issues can be filled in the issue tracker.
