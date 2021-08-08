package parser

import "github.com/go-mysql-org/go-mysql/canal"

type Contract interface {
	GetBinLogData(element interface{}, e *canal.RowsEvent, n int) error
}
