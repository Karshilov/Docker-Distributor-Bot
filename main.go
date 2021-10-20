package main

import (
	"Docker-Distributor-Bot/pkg/dockerClient"
	"path"

	"github.com/docker/docker/pkg/archive"
)

func main() {
	tar, err := archive.TarWithOptions(path.Join("utils/nvidia-cuda"), &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	err = dockerClient.BuildImage(0, tar)
	if err != nil {
		panic(err)
	}
}
