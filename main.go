package main

import (
	"erply/infra/database"
	"erply/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	var err error
	var db *sqlx.DB

	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	if db, err = database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	r := middlewares.NewRouter(db)
	log.Fatal(r.Run(":9000"))
}
