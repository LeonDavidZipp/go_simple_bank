package main

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/LeonDavidZipp/go_simple_bank/api"
	db "github.com/LeonDavidZipp/go_simple_bank/db/sqlc"
)

func main() {
	conn, err := sql.Open(
		os.Getenv("dbDriver"),
		os.Getenv("dbSource"),
	)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(os.Getenv("serverAddress"))
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
