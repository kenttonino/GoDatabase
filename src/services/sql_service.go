package services

import (
	"GoDatabase/src/utils"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSQLInstance() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database/sqlite.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateSQLUserTable() error {
	db, err := CreateSQLInstance()

	if err != nil {
		return err
	}

	defer db.Close()

	createTableStatement := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE
		);`

	_, err = db.Exec(createTableStatement)

	if err != nil {
		fmt.Println(utils.TextRed(err.Error()))
		return err
	}

	return nil
}
