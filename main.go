package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/sakhaei-wd/banker/api"
	db "github.com/sakhaei-wd/banker/db/sqlc"
)

//we will refactor the code to load all of these configurations from environment variables or a setting file
const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"
	serverAddress = "0.0.0.0:8080" 
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
    if err != nil {
        log.Fatal("cannot start server:", err)
    }
}
