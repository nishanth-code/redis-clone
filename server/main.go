package main

import (
	"log"

	"github.com/nishanth-code/redis-clone/internal/WAL"
	"github.com/nishanth-code/redis-clone/internal/config"
	"github.com/nishanth-code/redis-clone/internal/server"
	"github.com/nishanth-code/redis-clone/internal/store"
)

func main() {

	store := store.NewStore()
	cfg := config.Load()

    wal, err := wal.New(cfg.WALPath)
	

	if err != nil {
		log.Fatal(err)
	}

	err = wal.Replay(store)

	if err != nil {
		log.Fatal(err)
	}

	tcpServer :=
		server.NewTCPServer(
			cfg.Port,
			store,
			wal,
		)

	log.Fatal(
		tcpServer.Start(),
	)
}