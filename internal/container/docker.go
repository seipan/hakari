package container

import (
	"github.com/docker/docker/client"
)

type Container struct {
	Client *client.Client
}

func NewContainer() (*Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Container{
		Client: cli,
	}, nil
}

func (c *Container) Close() error {
	return c.Client.Close()
}
