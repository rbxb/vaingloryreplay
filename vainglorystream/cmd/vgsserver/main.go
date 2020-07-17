package main

import (
	"net/http"
	"flag"
	"github.com/rbxb/vaingloryreplay/vainglorystream/cmd/vgsserver/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":8080", "The port to listen on. (:8080)")
}

func main() {
	flag.Parse()
	srvr := server.NewServer()
	if err := http.ListenAndServe(port, srvr); err != nil {
		panic(err)
	}
}