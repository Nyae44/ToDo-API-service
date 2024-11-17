package main

import (
	"log"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)

type TodoItem struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

func main() {
	d, err := db.New("postgres", "example", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}

	svc := todo.NewService(d)
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}
