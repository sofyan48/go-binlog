package repositories

import (
	"context"

	"github.com/sofyan48/go-binlog/src/app/entity"
)

const (
	TABLE_USER = "user"
)

type UserInterface interface {
	FetchRow(ctx context.Context, id uint64) (*entity.DBUser, error)
}
