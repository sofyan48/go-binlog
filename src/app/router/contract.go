package router

import "context"

//

type Router interface {
	LogRouter(ctx context.Context)
}
