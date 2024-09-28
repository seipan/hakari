package container

import (
	"github.com/docker/docker/api/types/container"

	"context"
)

func (c *Container) GetContainer(target string) (bool, error) {
	if target == "" {
		return false, nil
	}

	containers, err := c.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return false, err
	}

	var targetContainerID string
	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+target {
				targetContainerID = container.ID
				break
			}
		}
	}

	if targetContainerID != "" {
		return true, nil
	}
	return false, nil
}
