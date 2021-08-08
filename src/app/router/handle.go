package router

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/sofyan48/go-binlog/src/app/usecase/contract"
	"github.com/sofyan48/go-binlog/src/bootstrap"
	"github.com/sofyan48/go-binlog/src/pkg/binlog"
)

type router struct {
}

// NewRouter initialize new router wil return Router Interface
func NewRouter() Router {
	return &router{}
}

func NewBinlogCommand(log binlog.Contract) contract.LogProcessor {
	return &binlogCommand{
		binLog: log,
	}
}

func (rtr *router) LogRouter() {
	binLog := bootstrap.GetBinlog()
	binlogStart := NewBinlogCommand(binLog)
	binlogStart.Exec()
}

type binlogCommand struct {
	binLog binlog.Contract
	canal.DummyEventHandler
}

func (l *binlogCommand) Exec() error {
	l.binLog.Exec(&binlogCommand{})
	return nil
}

func (l *binlogCommand) OnRow(e *canal.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()
	var n = 0
	var k = 1

	if e.Action == canal.UpdateAction {
		n = 1
		k = 2
	}

	for i := n; i < len(e.Rows); i += k {
		l.logCMDRouter(ctx, cancel, e.Table.Name, e.Action, e, i)
	}

	return nil
}

func (l *binlogCommand) String() string {
	return "binlogHandler"
}
