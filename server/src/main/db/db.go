package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Connection *sqlx.DB
}

func NewDatabase(dsn string) (*Database, error) {
	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{Connection: conn}, nil
}
