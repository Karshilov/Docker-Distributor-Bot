package dockerClient

import (
	"Docker-Distributor-Bot/utils/config"
	"context"
	"time"
)

func DeleteContainer(hostId uint, qqnum int, containerID string) error {
	cli, err := getDockerClient(hostId)
	hostInfo := config.GetConfig()[hostId]
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	return nil
}
