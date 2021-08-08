package binlog

import "github.com/go-mysql-org/go-mysql/canal"

type Contract interface {
	Exec(event canal.EventHandler)
}
