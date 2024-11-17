package main

import (
	"github.com/Nyae44/ToDo-API-service/internal/db"
	"github.com/Nyae44/ToDo-API-service/internal/todo"
	"github.com/Nyae44/ToDo-API-service/internal/transport"
	"log"
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
