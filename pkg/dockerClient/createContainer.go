package dockerClient

import (
	"Docker-Distributor-Bot/pkg/simpledb"
	"Docker-Distributor-Bot/utils/config"
	"Docker-Distributor-Bot/utils/random"
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func GetAvailablePort(host string) (int, error) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", host))
	if err != nil {
		return 0, err
	}
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func CreateContainer(hostId uint, pubKey string, qqnum int64) error {
	cli, err := getDockerClient(hostId)
	hostInfo := config.GetConfig()[hostId]
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	passwd := random.GetRandomPassword()
	port, err := GetAvailablePort(hostInfo.Host)
	if err != nil {
		return err
	}
	cfg := &container.Config{
		Env: []string{
			"ROOT_PASSWORD=" + passwd,
			"AUTHORIZED_KEYS=" + pubKey,
		},
	}
	curID, err := simpledb.GetLatestContainerId(hostId, qqnum)
	if err != nil {
		return err
	}
	containerName := fmt.Sprintf("%d-%d", qqnum, curID+1)
	hostCfg := &container.HostConfig{
		PortBindings: nat.PortMap{
			"2222/tcp": []nat.PortBinding{{
				HostIP:   hostInfo.Host,
				HostPort: strconv.Itoa(port),
			}},
		},
		Binds: []string{fmt.Sprintf("/var/%s/persist-data", containerName)},
	}
	res, err := cli.ContainerCreate(ctx, cfg, hostCfg, nil, nil, containerName)
	if err != nil {
		return err
	}
	err = simpledb.UpdateLatestContainer(simpledb.ContainerInfo{
		Id:            res.ID,
		Cid:           curID + 1,
		Name:          containerName,
		Port:          strconv.Itoa(port),
		HostID:        strconv.Itoa(int(hostId)),
		InitialPasswd: passwd,
		UserID:        qqnum,
	})
	if err != nil {
		return err
	}
	return nil
}
