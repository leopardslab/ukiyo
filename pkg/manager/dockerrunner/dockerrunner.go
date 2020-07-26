package dockerrunner

import (
	docker "github.com/fsouza/go-dockerclient"
	"log"
)

type ResponseObj struct {
	ResponseCode int
	ResponseDesc string
}

func ContainerRunner(containerId string) (ResponseObj, error) {
	var responseObj ResponseObj
	log.Println("Starting container Runner : containerId: " + containerId)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Running Docker Container Id: " + containerId
		return responseObj, err
	}

	err = client.StartContainer(containerId, &docker.HostConfig{
		PublishAllPorts: true,
	})
	if err != nil {
		log.Println(err)
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Container Run failure :" + containerId
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Run The Container :" + containerId
	}

	return responseObj, err
}
