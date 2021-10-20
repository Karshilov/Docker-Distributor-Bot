package dockerClient

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
)

func BuildImage(hostId uint, tar io.ReadCloser) error {
	cli, err := getDockerClient(hostId)
	if err != nil {
		return err
	}
	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{"nvidia-cuda"},
		Remove:     true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		lastLine := scanner.Text()
		fmt.Println(lastLine)
	}
	return nil
}
