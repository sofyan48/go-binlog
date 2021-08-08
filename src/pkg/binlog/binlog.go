package binlog

import (
	"github.com/go-mysql-org/go-mysql/canal"
)

type binLog struct {
	cnl *canal.Canal
}

func NewBinlog(cnl *canal.Canal) Contract {
	return &binLog{
		cnl: cnl,
	}
}

func (m *binLog) Exec(event canal.EventHandler) {
	coords, err := m.cnl.GetMasterPos()
	if err == nil {
		m.cnl.SetEventHandler(event)
		m.cnl.RunFrom(coords)
	}
}
