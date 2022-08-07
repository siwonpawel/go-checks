package main

import (
	"log"
	"net/http"

	_ "github.com/siwonpawel/go-stdnum/stdnum/validation/pl"
)

const addr = ":80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting http server on port %v...\n", addr)
	if err := http.ListenAndServe(addr, app.handlers()); err != nil {
		panic(err)
	}
}
