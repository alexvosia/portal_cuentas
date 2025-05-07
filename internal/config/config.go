package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type ConnDb interface {
	GetConn() *sql.DB
	CloseConn() error
}

type ChecaConn struct {
	Db *sql.DB
}

var checaConn *ChecaConn

func (c ChecaConn) GetConn() *sql.DB {
	//TODO implement me
	panic("implement me")
}

type PicConn struct {
	Db *sql.DB
}

func (p PicConn) GetConn() *sql.DB {
	db, err := sql.Open("sqlserver", "sqlserver://pic_owner:pic_1234@localhost:1433?database=pic")
}

type DuckConn struct {
	Db *sql.DB
}

func GetDB() *DB {
	if db == nil {
		db = &DB{}
	}
	return db
}
