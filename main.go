package main

import (
	"os"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/leondavidzipp/go_simple_bank/api"
	db "github.com/leondavidzipp/go_simple_bank/db/sqlc"
)

func main() {
	conn, err := sql.Open(
		os.Getenv("dbDriver"),
		os.Getenv("dbSource"),
	)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	var store *Store = db.NewStore(conn)
	var server *Server = api.NewServer(store)

	err = server.start(os.GetEnv("serverAddress"))
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
