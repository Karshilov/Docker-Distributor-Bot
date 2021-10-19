package dockerClient

import (
	"Docker-Distributor-Bot/pkg/sshTunnel"
	"Docker-Distributor-Bot/utils/config"
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func BuildImage(hostId uint, tar io.ReadCloser) /* types.ImageBuildResponse */ {
	cfg := config.GetConfig()
	tunnel := sshTunnel.NewSSHTunnel(
		cfg[hostId].User+"@"+cfg[hostId].Host,
		sshTunnel.PrivateKeyFile(cfg[hostId].Key),
		"/var/run/docker.sock",
	)
	ch := make(chan bool)
	go tunnel.Start(ch)
	<-ch
	cli, err := client.NewClientWithOpts(
		client.WithHost("tcp://127.0.0.1:"+strconv.Itoa(tunnel.Local.Port)),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
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
		panic(err)
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		lastLine := scanner.Text()
		fmt.Println(lastLine)
	}
}
