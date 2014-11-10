// Limited HTTP static server in Go.
//
// Installation:
//
//     go get github.com/felixrabe-go/limited-http-static-server
//
// Usage:
//
//     limited-http-static-server -p 8080 .
package main

import (
	"github.com/felixrabe-go/limited-http-static-server/flags"
	"github.com/felixrabe-go/limited-http-static-server/net"
)

func main() {
	port, dir := flags.Parse()
	net.Serve(port, dir)
}
