// Package mariadb
package mariadb

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CreateSession create new session maria db
func CreateSession(cfg *Config) (*sqlx.DB, error) {
	if len(strings.Trim(cfg.Charset, "")) == 0 {
		cfg.Charset = "UTF8"
	}
	connStr := fmt.Sprintf(connStringTemplate,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Timeout,
		cfg.Charset,
	)

	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	return db, nil
}
