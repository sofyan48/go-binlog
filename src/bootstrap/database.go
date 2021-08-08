// Package bootstrap
package bootstrap

import (
	"log"
	"os"
	"time"

	"github.com/sofyan48/go-binlog/src/pkg/mariadb"
	"github.com/sofyan48/go-binlog/src/pkg/util"
)

// RegistryMariaDB initialize maria or mysql session
func RegistryMariaDB() mariadb.Adapter {
	db, err := mariadb.NewMariaMasterSlave(
		&mariadb.Config{
			Host:         os.Getenv("DB_HOST"),
			Name:         os.Getenv("DB_NAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			Port:         util.StringToInt(os.Getenv("DB_PORT")),
			User:         os.Getenv("DB_USER"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET"),
		},

		&mariadb.Config{
			Host:         os.Getenv("DB_HOST_READ"),
			Name:         os.Getenv("DB_NAME_READ"),
			Password:     os.Getenv("DB_PASSWORD_READ"),
			Port:         util.StringToInt(os.Getenv("DB_PORT_READ")),
			User:         os.Getenv("DB_USER_READ"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT_READ"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN_READ")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN_READ")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME_READ"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET_READ"),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// RegistryMariaMasterSlave initialize maria or mysql session
func RegistryMariaMasterSlave() mariadb.Adapter {
	db, err := mariadb.NewMariaMasterSlave(
		&mariadb.Config{
			Host:         os.Getenv("DB_HOST"),
			Name:         os.Getenv("DB_NAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			Port:         util.StringToInt(os.Getenv("DB_PORT")),
			User:         os.Getenv("DB_USER"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET"),
		},

		&mariadb.Config{
			Host:         os.Getenv("DB_HOST_READ"),
			Name:         os.Getenv("DB_NAME_READ"),
			Password:     os.Getenv("DB_PASSWORD_READ"),
			Port:         util.StringToInt(os.Getenv("DB_PORT_READ")),
			User:         os.Getenv("DB_USER_READ"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT_READ"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN_READ")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN_READ")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME_READ"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET_READ"),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
