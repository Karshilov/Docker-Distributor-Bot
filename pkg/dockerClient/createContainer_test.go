package dockerClient

import (
	"fmt"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	if err := CreateContainer(1, "", 3142534138); err != nil {
		t.Fatal(fmt.Sprintf("created failed %v", err))
	}
}
