package dockerClient

import (
	"Docker-Distributor-Bot/pkg/sshTunnel"
	"Docker-Distributor-Bot/utils/config"
	"strconv"

	"github.com/docker/docker/client"
)

func getDockerClient(hostId uint) (*client.Client, error) {
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
		return cli, err
	}
	return cli, nil
}
