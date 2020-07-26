package dockercreater

import (
	docker "github.com/fsouza/go-dockerclient"
	"log"
	"strings"
	"ukiyo/internal/containerscheduler"
)

type ResponseObj struct {
	Name         string
	ContainerId  string
	ResponseCode int
	ResponseDesc string
}

func ContainerCreate(name string, imageName string) (ResponseObj, error) {
	var responseObj ResponseObj
	client, err := docker.NewClientFromEnv()
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Creating Docker Client Env for : " + name
		return responseObj, err
	}

	ports := containerscheduler.GetBindingPortsForContainerCreate(name)
	portBindings := make(map[docker.Port][]docker.PortBinding)
	exposedPorts := map[docker.Port]struct{}{}

	if len(ports) > 0 {
		for _, p := range ports {
			s := strings.Split(p, ":")
			extPort, intPort := s[0], s[1]
			log.Println("extPort : " + extPort + " intPort : " + intPort)
			dockerPort := docker.Port(intPort + "/tcp")
			portBindings[dockerPort] = []docker.PortBinding{{HostIP: "0.0.0.0", HostPort: extPort}}
			exposedPorts[dockerPort] = struct{}{}
		}
	} else {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "No valid Binding Ports for :" + name
		return responseObj, err
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image:        imageName,
			ExposedPorts: exposedPorts,
		},
		HostConfig: &docker.HostConfig{
			PortBindings: portBindings,
		},
	})
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Creating Container for :" + name
		return responseObj, err
	}
	log.Println(container)

	responseObj.ResponseCode = 0
	responseObj.ContainerId = container.ID
	responseObj.ResponseDesc = "Successfully Create The Container :" + name
	return responseObj, err
}
