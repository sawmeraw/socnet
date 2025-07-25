package main

import (
	"log"

	"github.com/sawmeraw/gogo/internal/db"
	"github.com/sawmeraw/gogo/internal/env"
	"github.com/sawmeraw/gogo/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
