package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type Log struct {
	id         int
	remoteAddr string
	localAddr  string
	pingAt     string
	live       int
}

func init() {
	db, err = sql.Open("mysql", "root:@/server")
	if err != nil {
		panic(err.Error())
	}
}

func insert(query string) {
	//defer db.Close()

	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
func selectQ(query string) []Log {
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	logs := make([]Log, 0)

	defer rows.Close()

	for rows.Next() {
		log := new(Log)
		if err := rows.Scan(&log.id, &log.remoteAddr, &log.localAddr, &log.pingAt, &log.live); err != nil {
			panic(err)
		}
		logs = append(logs, *log)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	defer rows.Close()
	return logs
}
