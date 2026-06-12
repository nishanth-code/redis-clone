package main

import (
	"log"

	"github.com/nishanth-code/redis-clone/internal/server"
	"github.com/nishanth-code/redis-clone/internal/store"
)

func main() {

	store := store.NewStore()

	tcpServer :=
		server.NewTCPServer(
			"8081",
			store,
		)

	log.Fatal(
		tcpServer.Start(),
	)
}