package db

import "github.com/jmoiron/sqlx"

type DB struct {
	Conn *sqlx.DB
}
