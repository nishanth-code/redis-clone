package wal

import (
	"bufio"
	"os"
	"strings"
	"github.com/nishanth-code/redis-clone/internal/store"
)

func (w *WAL) Replay(store *store.Store) error {
	file, err := os.Open(w.path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()

		parts := strings.Split(line,"|")

		if len(parts) < 3 {
			continue
		}	
		switch parts[0] {
		case "SET":
			store.Set(parts[1], parts[2])
		case "DEL":
			store.Delete(parts[1])
		}
	}
	return scanner.Err()
}