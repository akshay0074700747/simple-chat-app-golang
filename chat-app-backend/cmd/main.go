package main

import (
	"log"
	"os"

	"github.com/akshay0074700747/chat-app/database"
	"github.com/akshay0074700747/chat-app/database/provider"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db_addr := os.Getenv("DATABASE_ADDR")

	db, err := database.FireUp(db_addr)

	if err != nil {
		log.Fatal(err.Error())
	}

	engine := provider.Provide(db)

	engine.Start()

}
