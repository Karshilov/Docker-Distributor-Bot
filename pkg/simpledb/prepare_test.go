package simpledb

import "testing"

func TestPrepare(t *testing.T) {
	if Prepare() != nil {
		t.Fatal("prepare() failed")
	}
}
