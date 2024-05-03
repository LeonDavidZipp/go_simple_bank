package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(
		os.Getenv("dbDriver"),
		os.Getenv("dbSource"),
	)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
