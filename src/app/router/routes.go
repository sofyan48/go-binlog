package router

import (
	"context"
	"fmt"
	"time"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/sofyan48/go-binlog/src/app/entity"
	"github.com/sofyan48/go-binlog/src/app/repositories"
	"github.com/sofyan48/go-binlog/src/app/usecase/stream"
	"github.com/sofyan48/go-binlog/src/bootstrap"
	"github.com/sofyan48/go-binlog/src/pkg/parser"
)

func (l *binlogCommand) logCMDRouter(ctx context.Context, cnc context.CancelFunc, table, action string, row *canal.RowsEvent, i int) {
	parseFunc := parser.NewParserLog()

	// bootstrap
	db := bootstrap.RegistryMariaMasterSlave()
	elastic := bootstrap.NewSessionElastic()

	// repositories
	userRepo := repositories.NewUserRepositories(db)
	time.Sleep(5 * time.Second)
	// routes
	switch table {
	case "user":
		streamact := stream.NewStream(userRepo, elastic)
		userData := entity.User{}
		parseFunc.GetBinLogData(&userData, row, i)
		switch action {
		case canal.InsertAction:
			streamact.Exec(ctx, userData.Id)
		case canal.UpdateAction:
			fmt.Println("Update Action")
		case canal.DeleteAction:
			fmt.Println("Delete Action")

		}
	}
	cnc()
}
