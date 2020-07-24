package pushmanager

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func ContainerCreate(imageName string) (*docker.Container, error) {

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	client.RemoveContainer(docker.RemoveContainerOptions{
		ID: "",
	})

	ports := []string{"80", "8080"}
	portBindings := make(map[docker.Port][]docker.PortBinding)
	exposedPorts := map[docker.Port]struct{}{}

	for _, p := range ports {
		dockerPort := docker.Port(p + "/tcp")
		portBindings[dockerPort] = []docker.PortBinding{{HostIP: "0.0.0.0", HostPort: p}}
		exposedPorts[dockerPort] = struct{}{}
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "demo-app-name",
		Config: &docker.Config{
			Image:        imageName,
			ExposedPorts: exposedPorts,
		},
	})

	err = client.StopContainer("8e8e3b6a5423", 0)
	if err != nil {
		log.Println(err)
	}

	client.StartContainer(container.ID, &docker.HostConfig{
		PortBindings:    portBindings,
		PublishAllPorts: true,
	})

	log.Println(imageName)
	return client.InspectContainer(container.ID)
}
