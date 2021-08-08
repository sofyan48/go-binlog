// Package mariadb
package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mariaMasterSlave struct {
	db      *sqlx.DB
	dbRead  *sqlx.DB
	cfg     *Config
	cfgRead *Config
}

// NewMariaMasterSlave initialize maria db for write read
func NewMariaMasterSlave(cfgWrite *Config, cfgRead *Config) (Adapter, error) {
	x := mariaMasterSlave{cfg: cfgWrite, cfgRead: cfgRead}

	e := x.initialize()

	return &x, e
}

func (d *mariaMasterSlave) initialize() error {
	dbWrite, err := CreateSession(d.cfg)

	if err != nil {
		return err
	}

	err = dbWrite.Ping()
	if err != nil {
		return err
	}

	dbRead, err := CreateSession(d.cfgRead)
	if err != nil {
		return err
	}

	err = dbRead.Ping()
	if err != nil {
		return err
	}

	d.db = dbWrite
	d.dbRead = dbRead

	return nil
}

// QueryRow select single row database will return  sql.row raw
func (d *mariaMasterSlave) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.selector().QueryRowContext(ctx, query, args...)
}

// QueryRows select multiple rows of database will return  sql.rows raw
func (d *mariaMasterSlave) QueryRows(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.selector().QueryContext(ctx, query, args...)
}

// Fetch select multiple rows of database will cast data to struct passing by parameter
func (d *mariaMasterSlave) Fetch(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return d.selector().SelectContext(ctx, dst, query, args...)
}

// FetchRow fetching one row database will cast data to struct passing by parameter
func (d *mariaMasterSlave) FetchRow(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return d.selector().GetContext(ctx, dst, query, args...)
}

// Exec execute mysql command query
func (d *mariaMasterSlave) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

// BeginTx start new transaction session
func (d *mariaMasterSlave) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}

// Ping check database connectivity
func (d *mariaMasterSlave) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

// HealthCheck checking healthy of database connection
func (d *mariaMasterSlave) HealthCheck() error {
	var err1, err2 error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err1 = d.Ping(context.Background())
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		err2 = d.Ping(context.Background())
		wg.Done()
	}()

	wg.Wait()

	if err1 != nil && err2 != nil {
		return fmt.Errorf("database write error:%s; database read error:%s; ", err1.Error(), err2.Error())
	}

	if err1 != nil {
		return fmt.Errorf("database write error:%s;", err1.Error())

	}

	if err2 != nil {
		return fmt.Errorf("database read error:%s;", err2.Error())

	}

	return nil

}

func (d *mariaMasterSlave) selector() *sqlx.DB {
	if d.dbRead != nil {
		return d.dbRead
	}

	return d.db
}
