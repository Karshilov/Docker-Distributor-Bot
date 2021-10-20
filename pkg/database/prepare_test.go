package simpledb

import "testing"

func TestPrepare(t *testing.T) {
	if prepare() != nil {
		t.Fatal("prepare() failed")
	}
}
