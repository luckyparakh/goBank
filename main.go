package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/luckyparakh/goBank/api"
	db "github.com/luckyparakh/goBank/db/sqlc"
	"github.com/luckyparakh/goBank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err!=nil{
		log.Fatal("cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	server := api.NewServer(db.NewStore(conn))
	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
