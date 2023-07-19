package database

import (
	"database/sql"
	"fmt"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
	"gitlab.com/binsabit/billing/config"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	dbConn, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	//db = test
	return dbConn, err
}
