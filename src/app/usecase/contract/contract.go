package contract

import "context"

type LogProcessor interface {
	Exec(ctx context.Context) error
}

type Usecase interface {
	Exec(ctx context.Context, id int) error
}
