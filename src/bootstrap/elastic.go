package bootstrap

import (
	"log"
	"os"
	"strings"

	"github.com/sofyan48/go-binlog/src/pkg/elastic"
)

// NewSessionElastic ...
func NewSessionElastic() elastic.Client {
	addr := strings.Split(os.Getenv("ELASTICSEARCH_URL"), ",")
	config := &elastic.Configuration{
		Address:  addr,
		Username: os.Getenv("ELASTICSEARCH_USERNAME"),
		Password: os.Getenv("ELASTICSEARCH_PASSWORD"),
	}
	cfgClient := elastic.Config(config)
	client, err := elastic.NewClient(cfgClient)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
