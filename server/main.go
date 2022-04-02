package main

import (
	"log"
	"todo/src/db"
	"todo/src/server"
)

func main() {
	rcon := db.ConnectDB()
	if !rcon.Success {
		log.Fatalf("Failed to connect DB %v", rcon.Error.Error())
	}
	con := rcon.Value
	r := server.Router(con)
	r.Run()
}
