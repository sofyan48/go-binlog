package main

import (
	"fmt"
	"time"

	"github.com/sofyan48/go-binlog/src/pkg/binlog"
)

func main() {
	go binlog.BinlogListener()
	time.Sleep(2 * time.Minute)
	fmt.Print("Thanks")
}
