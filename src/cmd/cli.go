package cmd

import (
	"log"

	"github.com/sofyan48/go-binlog/src/cmd/app"
	"github.com/sofyan48/go-binlog/src/cmd/migration"
	"github.com/spf13/cobra"
)

func Start() {

	rootCmd := &cobra.Command{}

	cmd := []*cobra.Command{
		app.Stream(),
		migration.Start(),
	}
	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error:> ", err)
	}
}
