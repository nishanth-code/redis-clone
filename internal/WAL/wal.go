package wal

import (
	"fmt"
	"os"
	"sync"
)

type WAL struct {
	file *os.File
	path string
	mu   sync.Mutex
}

func New(path string) (*WAL, error) {

	file, err := os.OpenFile(
		path,
		os.O_APPEND|
			os.O_CREATE|
			os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: file,
		path: path,
	}, nil
}

func (w *WAL) Append(
	entry LogEntry,
) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	_,err := w.file.WriteString(entry.Operation+"|"+entry.Key+"|"+entry.Value+"\n")
	fmt.Println("Appended to WAL:" + entry.Operation + "|" + entry.Key + "|" + entry.Value + "\n")
	if err != nil {
		return err
	}

	if err := w.file.Sync(); err != nil {
		return err
	}
	return nil


}