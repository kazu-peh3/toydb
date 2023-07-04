package main

import (
	"flag"
	"log"
	"toydb/db"
)

func server() {
	_, err := db.NewToyDb()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	serverMode = flag.Bool("server", false, "boot the db server")
)

func main() {
	flag.Parse()

	if *serverMode {
		server()
		return
	}
}
