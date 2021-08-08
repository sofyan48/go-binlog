package bootstrap

import (
	"os"

	"github.com/sofyan48/go-binlog/src/pkg/canal"
)

func GetCanal() canal.Contract {
	cfg := &canal.Config{
		Flavor:   "mysql",
		Host:     os.Getenv("BINLOG_DB_HOST"),
		Password: os.Getenv("BINLOG_DB_PASSWORD"),
		Port:     os.Getenv("BINLOG_DB_PORT"),
		User:     os.Getenv("BINLOG_DB_USER"),
	}
	return canal.NewCanal(cfg)
}
