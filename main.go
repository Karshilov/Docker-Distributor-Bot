package main

import (
	"Docker-Distributor-Bot/pkg/dockerClient"
	"os"
	"path"

	"github.com/docker/docker/pkg/archive"
)

func main() {
	pwd, _ := os.Getwd()
	println(path.Join(pwd, "utils/nvidia-cuda"))
	tar, err := archive.TarWithOptions(path.Join("utils/nvidia-cuda"), &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	dockerClient.BuildImage(0, tar)
}
