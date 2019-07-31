package util

import (
	"database/sql"
	"time"
)

func DBConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "people"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(1* time.Second)

	return db
}
