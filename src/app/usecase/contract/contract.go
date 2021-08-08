package contract

import "context"

type Usecase interface {
	Exec(ctx context.Context, id int) error
}
