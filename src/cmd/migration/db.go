// Package migration
package migration

import (
	"os"
	"time"

	"github.com/sofyan48/go-binlog/src/pkg/mariadb"
	"github.com/sofyan48/go-binlog/src/pkg/util"
	"github.com/spf13/cobra"
)

func Start() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			mariadb.DatabaseMigration(&mariadb.Config{
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
			})
		},
	}

	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")

	return migrateCmd
}
