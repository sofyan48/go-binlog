package serve

import "context"

type Command interface {
	Exec(ctx context.Context)
}
