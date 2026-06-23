package config

type Config struct {
	Port    string
	WALPath string
}

func Load() Config {
	return Config{
		Port:    "8081",
		WALPath: "data/wal.log",
	}
}

