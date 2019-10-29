package main

import (
	"fmt"
	"log"
	"uptime/db"
	"uptime/src"
	"os"
)

func main() {

	//Read database configuration from config.toml
	uri := db.ReadDBConfig()

	//Connect Mongo database
	session, err := db.Connect(uri)
	if err != nil {
		log.Fatalf("ERR_DB_CONN: %s", err)
	}

	fmt.Println("DB_CONN_SUCCESS")
	defer session.Terminate()

	handler := src.New(session)

	if len(os.Args) > 1 && os.Args[1] == "from" {
		handler.CalUptime()
	}
}
