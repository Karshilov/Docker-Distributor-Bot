package dockerClient

import (
	"Docker-Distributor-Bot/pkg/simpledb"
	"context"
	"time"

	"github.com/docker/docker/api/types"
)

func DeleteContainer(hostId uint, qqnum int64, containerID string) error {
	cli, err := getDockerClient(hostId)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	err = simpledb.CheckOwner(containerID, qqnum)
	if err != nil {
		return err
	}
	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}
	err = simpledb.DeleteContainer(containerID, qqnum)
	if err != nil {
		return err
	}
	return nil
}
