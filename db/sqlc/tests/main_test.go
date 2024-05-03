package db

import (
	"testing"
	"database/sql"
	"os"
	"log"
)

var testQueries *Queries;


funct TestMain(m *testing.M) {
	conn, err := sql.Open(
		os.Getenv("dbDriver"),
		os.Gertenv("dbSource")
	);
	if err != nil {
		log.Fatal("Cannot connect to db:", err);
	}
	testQueries = New(conn);
	os.Exit(m.Run());
}
