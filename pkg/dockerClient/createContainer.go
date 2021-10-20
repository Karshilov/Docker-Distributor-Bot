package dockerClient

import (
	"context"
	"time"

	"Docker-Distributor-Bot/utils/random"

	"github.com/docker/docker/api/types/container"
)

func createContainer(hostId uint, pubKey string, qqnum int64) error {
	cli, err := getDockerClient(hostId)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	cfg := &container.Config{
		Env: []string{
			"ROOT_PASSWORD=" + random.GetRandomPassword(),
			"AUTHORIZED_KEYS=" + pubKey,
		},
	}
	return nil
}
