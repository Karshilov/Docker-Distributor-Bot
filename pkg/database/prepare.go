package simpledb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func prepare() error {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return err
	}
	defer db.Close()
	table := `
		CREATE TABLE IF NOT EXISTS user (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(128) NULL
		);
    `
	_, err = db.Exec(table)
	if err != nil {
		return err
	}
	table = `
		CREATE TABLE IF NOT EXISTS container (
			id VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			port VARCHAR(255) NOT NULL,
			serverID VARCHAR(255) NOT NULL,
			initialPasswd VARCHAR(255) NOT NULL,
			userID VARCHAR(255) NOT NULL
		);
	`
	_, err = db.Exec(table)
	if err != nil {
		return err
	}
	return nil
}
