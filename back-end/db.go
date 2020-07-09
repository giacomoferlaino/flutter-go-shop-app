package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // loads sqlite DB driver
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	initStatement := `
	CREATE TABLE IF NOT EXISTS product (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, title TEXT, description TEXT, price REAL, imageUrl TEXT, isFavorite INTEGER);
	`

	_, err = db.Exec(initStatement)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
