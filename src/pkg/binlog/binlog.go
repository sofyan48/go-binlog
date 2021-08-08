package binlog

import (
	"fmt"
	"runtime/debug"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/sofyan48/go-binlog/src/pkg/mariadb"
)

type binlogHandler struct {
	canal.DummyEventHandler
	BinlogParser
}

func (h *binlogHandler) OnRow(e *canal.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()

	// base value for canal.DeleteAction or canal.InsertAction
	var n = 0
	var k = 1

	if e.Action == canal.UpdateAction {
		n = 1
		k = 2
	}
	for i := n; i < len(e.Rows); i += k {
		key := e.Table.Schema + "." + e.Table.Name
		switch key {
		case User{}.SchemaName() + "." + User{}.TableName():
			fmt.Println("======> ", e.Action)
			user := User{}
			err := h.GetBinLogData(&user, e, i)
			if err != nil {
				fmt.Println("Error:> ", err)
			}

			switch e.Action {
			case canal.UpdateAction:
				oldUser := User{}
				h.GetBinLogData(&oldUser, e, i-1)
				fmt.Printf("User %d name update from %s to %s\n", user.Id, oldUser.Name, user.Name)
			case canal.InsertAction:
				fmt.Printf("User %d is created with name %s\n", user.Id, user.Name)
			case canal.DeleteAction:
				fmt.Printf("User %d is deleted with name %s\n", user.Id, user.Name)
			default:
				fmt.Printf("Unknown action")
			}
		}

	}
	return nil
}

func (h *binlogHandler) String() string {
	return "binlogHandler"
}

func BinlogListener() {
	c, err := mariadb.GetCanal()
	if err == nil {
		coords, err := c.GetMasterPos()
		if err == nil {
			c.SetEventHandler(&binlogHandler{})
			c.RunFrom(coords)
		}
	}
}
