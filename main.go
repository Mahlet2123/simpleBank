package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load configurations", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can't create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start server:", err)
	}
}