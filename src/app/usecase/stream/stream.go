package stream

import (
	"github.com/sofyan48/go-binlog/src/app/usecase/contract"
)

type streams struct {
}

func NewStream() contract.LogProcessor {
	return &streams{}
}

func (c *streams) Exec() error {
	return nil
}
