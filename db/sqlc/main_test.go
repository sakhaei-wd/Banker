package db

import (
		"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"	
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"	
)

var testQueries *Queries




func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}