package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sakhaei-wd/banker/util"
)


var testQueries *Queries
var testDB *sql.DB



func TestMain(m *testing.M) {
	
	//../.. mean go to the parent folder.
	config, err := util.LoadConfig("../..")
	
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}