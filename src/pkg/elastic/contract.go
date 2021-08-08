package elastic

import "context"

// Configuration ...
type Configuration struct {
	Address  []string
	Username string
	Password string
}

// Client ..
type Client interface {
	Version() (map[string]interface{}, error)
	Mapping(ctx context.Context, index string, field interface{}) (map[string]interface{}, error)
	Search(ctx context.Context, index string, query map[string]interface{}) (interface{}, error)
	Insert(ctx context.Context, index, id string, field interface{}) (map[string]interface{}, error)
	Delete(ctx context.Context, index, id string) (map[string]interface{}, error)
}
