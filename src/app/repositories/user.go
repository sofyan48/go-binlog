package repositories

import (
	"context"

	"github.com/sofyan48/go-binlog/src/app/entity"
	"github.com/sofyan48/go-binlog/src/pkg/mariadb"
)

type user struct {
	db mariadb.Adapter
}

func NewUserRepositories(db mariadb.Adapter) UserInterface {
	return &user{
		db: db,
	}
}

// FecthRow data
func (r *user) FetchRow(ctx context.Context, id uint64) (*entity.DBUser, error) {
	p := &entity.DBUser{}
	query := `SELECT * FROM ` + TABLE_USER + ` WHERE  id=?`
	e := r.db.FetchRow(ctx, p, query, id)
	return p, e
}
