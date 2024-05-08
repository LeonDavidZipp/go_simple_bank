package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/LeonDavidZipp/go_simple_bank/api"
	"github.com/LeonDavidZipp/go_simple_bank/util"
	db "github.com/LeonDavidZipp/go_simple_bank/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(
		config.DBDriver,
		config.DBSource,
	)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
