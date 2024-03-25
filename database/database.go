package database

import (
	"database/sql"
	"sync"
	"todo/core"
)

var (
	dbOnce sync.Once
	db     *sql.DB
)

func InitDB() {
	var err error
	db, err = core.ConnectDb()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}

func GetDB() *sql.DB {
	dbOnce.Do(InitDB)
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
