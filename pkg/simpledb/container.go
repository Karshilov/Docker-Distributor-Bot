package simpledb

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type ContainerInfo struct {
	Id            string
	Cid           int
	UserID        int64
	HostID        string
	Name          string
	Port          string
	InitialPasswd string
}

func GetLatestContainerId(hostId uint, qqnum int64) (int, error) {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT cid FROM container WHERE hostID = ? AND userID = ?", hostId, qqnum)
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

func UpdateLatestContainer(c ContainerInfo) error {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return err
	}
	defer db.Close()
	prep, err := db.Prepare(`
		INSERT INTO 
		container(id, cid, name, port, hostID, initialPasswd, userId) 
		values(?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = prep.Exec(
		c.Id,
		c.Cid,
		c.Name,
		c.Port,
		c.HostID,
		c.InitialPasswd,
		c.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}

func CheckOwner(containerID string, qqnum int64) error {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query("SELECT userID FROM container WHERE id = ?", containerID)
	if err != nil {
		return err
	}
	if rows.Next() {
		var userID int64
		err = rows.Scan(&userID)
		if err != nil {
			return err
		}
		if userID != qqnum {
			return errors.New("you are not the owner")
		}
	}
	return nil
}

func DeleteContainer(containerID string, qqnum int64) error {
	db, err := sql.Open("sqlite3", "./dockerInfo.db")
	if err != nil {
		return err
	}
	defer db.Close()
	prep, err := db.Prepare("DELETE FROM container WHERE id = ?")
	if err != nil {
		panic(err)
	}

	_, err = prep.Exec(containerID)
	if err != nil {
		panic(err)
	}
	return nil
}
