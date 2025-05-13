package config

import (
	"database/sql"
	"log"

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

func NewMSSQLConn() *sql.DB {
	// Configurar conexión a la base de datos
	db, err := sql.Open("sqlserver", "sqlserver://pic_owner:pic_1234@localhost:1433?database=pic")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//defer func(db *sql.DB) {
//	err := db.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//}(db)
