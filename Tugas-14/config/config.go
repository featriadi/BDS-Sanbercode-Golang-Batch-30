//config/config.go, digunakan untuk melakukan konfigurasi MySQL.

package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	database string = "db_tugas14"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
