package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/meng621/simpleBank/api"
	db "github.com/meng621/simpleBank/db/sqlc"
	"github.com/meng621/simpleBank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load confit", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}

}
