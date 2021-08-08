package app

import (
	"github.com/sofyan48/go-binlog/src/serve"
	"github.com/spf13/cobra"
)

func Stream() *cobra.Command {

	return &cobra.Command{
		Use:   "run",
		Short: "Run HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {

			serve.NewCommand().Exec()
		},
	}
}
