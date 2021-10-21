package main

import (
	"Docker-Distributor-Bot/pkg/dockerClient"
	"Docker-Distributor-Bot/pkg/simpledb"
	"fmt"
)

func main() {
	if err := simpledb.Prepare(); err != nil {
		panic(err)
	}
	if err := dockerClient.CreateContainer(1, "", 3142534138); err != nil {
		println(fmt.Sprintf("created failed %v", err))
	}
}
