package wal

type LogEntry struct {
	Operation string
	Key       string
	Value     string
	timestamp int64
}