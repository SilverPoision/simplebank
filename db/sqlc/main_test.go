package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testqueries *Queries
var testDb *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:silver@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testqueries = New(testDb)

	os.Exit(m.Run())
}
