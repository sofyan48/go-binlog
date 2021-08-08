package mariadb

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
)

func GetCanal() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", "127.0.0.1", 3306)
	cfg.User = "root"
	cfg.Password = "root"
	cfg.Flavor = "mysql"

	cfg.Dump.ExecutionPath = ""

	return canal.NewCanal(cfg)
}
