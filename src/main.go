package main

import (
	"fmt"
	"time"

	"github.com/JackShadow/go-binlog-example/src/pkg/binlog"
)

func main() {
	go binlog.BinlogListener()
	time.Sleep(2 * time.Minute)
	fmt.Print("Thx for watching, goodbuy")
}
