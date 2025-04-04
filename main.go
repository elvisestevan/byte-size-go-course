package main

import (
	"byte-size-go-course/internal/todo"
	"byte-size-go-course/internal/transport"
	"log"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
