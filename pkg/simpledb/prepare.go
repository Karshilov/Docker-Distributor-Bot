package simpledb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Prepare() error {
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
			cid INTEGER NOT NULL,
			name VARCHAR(255) NOT NULL,
			port VARCHAR(255) NOT NULL,
			hostID VARCHAR(255) NOT NULL,
			initialPasswd VARCHAR(255) NOT NULL,
			userID INTEGER NOT NULL
		);
	`
	_, err = db.Exec(table)
	if err != nil {
		return err
	}
	return nil
}
