package stream

import (
	"context"
	"fmt"

	"github.com/sofyan48/go-binlog/src/app/repositories"
	"github.com/sofyan48/go-binlog/src/app/usecase/contract"
	"github.com/sofyan48/go-binlog/src/pkg/elastic"
)

type streams struct {
	repo   repositories.UserInterface
	search elastic.Client
}

func NewStream(repo repositories.UserInterface, el elastic.Client) contract.Usecase {
	return &streams{
		repo:   repo,
		search: el,
	}
}

func (c *streams) Exec(ctx context.Context, id int) error {
	fmt.Println(id)
	return nil
}
