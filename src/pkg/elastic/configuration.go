package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
)

// Config ...
func Config(cfg *Configuration) elasticsearch.Config {
	return elasticsearch.Config{
		Addresses: cfg.Address,
		Username:  cfg.Username,
		Password:  cfg.Password,
	}
}
