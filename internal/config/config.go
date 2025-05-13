package config

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"
	_ "github.com/joho/godotenv"
	_ "github.com/marcboeker/go-duckdb"
)

type ConnDb interface {
	GetConn() *sql.DB
	CloseConn() error
}

type DbConn struct {
	Db ConnDb
}

var checaConn *ConnDb
var picConn *ConnDb
var duckConn *ConnDb

func (c DbConn) GetConn(entorno string) *sql.DB {
	switch entorno {
	case "checa":
		if checaConn == nil {
			db, err := sql.Open("godror", "user=checa_owner password=checa_1234 connectString=localhost:1521/xe")
			if err != nil {
				log.Fatal(err)
			}
			checaConn.Db = db
		}
		return checaConn.Db
	case "pic":
		if picConn == nil {
			db, err := sql.Open("sqlserver", "sqlserver://pic_owner:pic_1234@localhost:1433?database=pic")
			if err != nil {
				log.Fatal(err)
			}
			picConn.Db = db
		}
		return picConn.Db
	case "duck":
		if duckConn == nil {
			db, err := sql.Open("duckdb", "file=./local_db/duckdb.db")
			if err != nil {
				log.Fatal(err)
			}
			duckConn.Db = db
		}
	}

}
func (c DbConn) CloseConn(entorno string) error {
	switch entorno {
	case "checa":
		if checaConn != nil {
			err := checaConn.Db.Close()
			if err != nil {
				return err
			}
			checaConn = nil
		}
	case "pic":
		if picConn != nil {
			err := picConn.Db.Close()
			if err != nil {
				return err
			}
			picConn = nil
		}
	case "duck":
		if duckConn != nil {
			err := duckConn.Db.Close()
			if err != nil {
				return err
			}
			duckConn = nil
		}
	}
	return nil
}

func getConfigDotENV() error {
	// Load environment variables from .env file
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}

	return nil
}
