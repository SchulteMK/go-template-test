package main

import (
	"github.com/SchulteMK/go-template-test/server"
	"log"
	"net/http"
)

func main() {
	ser, err := server.New("/", "./templates")
	if err != nil {
		log.Fatalf("could not create server: %v\n", err)
	}
	h := &http.Server{Addr: ":8080", Handler: ser}
	if err := h.ListenAndServe(); err != nil {
		log.Fatalf("error running server: %v\n", err)
	}
}
