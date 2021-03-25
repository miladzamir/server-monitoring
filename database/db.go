package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Demo() {

	db, err := sql.Open("mysql", "root:@/server")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO log (name)VALUES ('zamir')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
