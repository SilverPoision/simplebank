package main

import (
	"database/sql"
	"log"

	"github.com/Silverpoision/simple_bank/api"
	db "github.com/Silverpoision/simple_bank/db/sqlc"
	"github.com/Silverpoision/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannont start server:", err)
	}
}
