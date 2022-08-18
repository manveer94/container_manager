package main

import (
	"com/manveer/manager/database"
	"com/manveer/manager/server"
	"log"
	"os"
	"strconv"
)

func main() {
	port := 8080
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("invalid port provided")
		}
	}
	database.Init()
	server.Run(port)
}
