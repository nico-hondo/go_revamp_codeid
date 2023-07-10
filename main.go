package main

import (
	"context"
	"log"
	"os"

	"codeid.revamptwo/config"
	"codeid.revamptwo/repositories"
	"codeid.revamptwo/server"

	_ "github.com/lib/pq"
)

// _ "github.com/lib/pq"

func main() {
	log.Println("starting db_revamp")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database...")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	context := context.Background()
	queries := repositories.New(dbHandler)

	InsertData(context, queries)

	DeleteData(context, queries)
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}

	return "db_revamp"
}
