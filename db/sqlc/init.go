package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/hacker_noon?sslmode=disable"
)

var DbQueries *Queries
var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	DbQueries = New(Db)
}