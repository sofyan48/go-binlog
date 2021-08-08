package router

import (
	"context"
	"fmt"

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

	// repositories
	userRepo := repositories.NewUserRepositories(db)

	// routes
	switch table {
	case "user":
		streamact := stream.NewStream(userRepo)
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
