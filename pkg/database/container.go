package simpledb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getLatestContainerId(hostId uint, qqnum int64) (int, error) {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT cid FROM container WHERE hostID = ? AND userID = ?")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	ret := 0
	for rows.Next() {
		var cid int
		err = rows.Scan(&cid)
		if err != nil {
			return 0, err
		}
		if cid > ret {
			ret = cid
		}
	}
	return ret, nil
}
