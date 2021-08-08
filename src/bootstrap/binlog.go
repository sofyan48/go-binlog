package bootstrap

import (
	"log"

	"github.com/sofyan48/go-binlog/src/pkg/binlog"
)

func GetBinlog() binlog.Contract {
	canal := GetCanal()

	cn, err := canal.GetCanal()

	if err != nil {
		log.Fatal("Error 1 :>", err)
	}
	return binlog.NewBinlog(cn)
}
