package simpledb

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetLatestContainerId(t *testing.T) {
	if _, err := GetLatestContainerId(1, 1656632517); err != nil {
		t.Fatal("get failed")
	}
}
