package db

import (
	"testing"
	"database/sql"
	"os"
)

var testQueries *Queries;


funct TestMain(m *testing.M) {
	conn, err := sql.Open(
		os.Getenv("dbDriver"),
		os.Gertenv("dbSource")
	);
}