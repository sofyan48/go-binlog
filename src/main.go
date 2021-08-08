package main

import (
	"github.com/joho/godotenv"
	"github.com/sofyan48/go-binlog/src/cmd"
)

func main() {
	godotenv.Load()
	cmd.Start()
}
