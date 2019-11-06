package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/regen-network/testnets/util/uptime/db"
	"github.com/regen-network/testnets/util/uptime/src"
)

func main() {

	fmt.Println("Starting...")

	var (
		startBlock int
		endBlock   int
	)

	//Read the start, end block flags passed from cmd
	flag.IntVar(&startBlock, "start", -1, "start flag: Start Block Number")
	flag.IntVar(&endBlock, "end", -1, "end flag: End Block Number")

	flag.Parse()

	if startBlock < 0 || endBlock < 1 {
		panic("--start and/or --end block flags are missing. Use --start, --end to input the range of blocknumbers")
	}

	//Read database configuration from config.toml
	uri := db.ReadDBConfig()

	//Connect Mongo database
	session, err := db.Connect(uri)

	if err != nil {
		log.Fatalf("ERR_DB_CONN: %s", err)
	}

	fmt.Println("DB connection established successfully")

	//Close the session safely after the operations are done
	defer session.Terminate()

	handler := src.New(session)

	handler.CalculateUptime(int64(startBlock), int64(endBlock))
}
