package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sofyan48/go-binlog/src/serve"
	"github.com/spf13/cobra"
)

func Stream() *cobra.Command {

	return &cobra.Command{
		Use:   "run",
		Short: "Run HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-quit
				cancel()
			}()
			serve.NewCommand().Exec(ctx)
		},
	}
}
