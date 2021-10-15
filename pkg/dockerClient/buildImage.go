package dockerClient

import (
	"Docker-Distributor-Bot/pkg/sshTunnel"
	"Docker-Distributor-Bot/utils/config"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func BuildImage(hostId uint) /* types.ImageBuildResponse */ {
	cfg := config.GetConfig()
	fmt.Printf("%s\n", cfg[hostId].User+"@"+cfg[hostId].Host)
	tunnel := sshTunnel.NewSSHTunnel(
		cfg[hostId].User+"@"+cfg[hostId].Host,
		sshTunnel.PrivateKeyFile(cfg[hostId].Key),
		"/var/run/docker.sock",
	)
	go tunnel.Start()
	time.Sleep(time.Millisecond * 100)
	cli, err := client.NewClientWithOpts(client.WithHost("tcp://127.0.0.1:" + strconv.Itoa(tunnel.Local.Port)))
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
