package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Silverpoision/simple_bank/util"
	_ "github.com/lib/pq"
)

var testqueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("error parsing config", err)
		return
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testqueries = New(testDb)

	os.Exit(m.Run())
}
