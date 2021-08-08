package serve

import (
	"context"

	"github.com/sofyan48/go-binlog/src/app/router"
)

type commandRouter struct {
	route router.Router
}

func NewCommand() Command {
	return &commandRouter{
		route: router.NewRouter(),
	}
}

func (c *commandRouter) Exec(ctx context.Context) {
	c.route.LogRouter(ctx)
}
