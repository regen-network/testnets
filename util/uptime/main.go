package main

import (
	"fmt"
	"log"
	"uptime/db"
)

func main() {

	uri := db.ReadDBConfig()
	session, err := db.Connect(uri)
	if err != nil {
		log.Fatalf("ERR_DB_CONN: %s", err)
	}

	fmt.Println("DB_CONN_SUCCESS")
	defer session.Terminate()
}
