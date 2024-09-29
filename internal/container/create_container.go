package container

import (
	"context"

	"github.com/docker/docker/api/types/container"
)

func (c *Container) CreateContainer(containerImage string, containerName string) error {
	_, err := c.Client.ContainerCreate(context.Background(), &container.Config{
		Image: containerImage,
	}, nil, nil, nil, containerName)
	if err != nil {
		return err
	}
	return nil
}
