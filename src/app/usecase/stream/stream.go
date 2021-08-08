package stream

import (
	"context"
	"fmt"

	"github.com/sofyan48/go-binlog/src/app/repositories"
	"github.com/sofyan48/go-binlog/src/app/usecase/contract"
)

type streams struct {
	repo repositories.UserInterface
}

func NewStream(repo repositories.UserInterface) contract.Usecase {
	return &streams{
		repo: repo,
	}
}

func (c *streams) Exec(ctx context.Context, id int) error {
	fmt.Println(id)
	return nil
}
