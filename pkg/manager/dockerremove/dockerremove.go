package dockerremove

import (
	docker "github.com/fsouza/go-dockerclient"
	"log"
	"strings"
	"ukiyo/pkg/jencoder"
)

type ResponseObj struct {
	ResponseCode int
	ResponseDesc string
}

func RemoveRunningContainer(name string) (ResponseObj, error) {
	var responseObj ResponseObj
	log.Println("Starting Remove Running pods: " + name)

	client, err := docker.NewClientFromEnv()
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Creating Docker Client Env"
		return responseObj, err
	}

	imgs, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error List Running Docker Services"
		return responseObj, err
	}

	if len(imgs) > 0 {
		log.Println("List Running container ...." + jencoder.PrintJson(imgs))
		for _, img := range imgs {
			if len(img.Names) > 0 {
				for _, cName := range img.Names {
					if strings.EqualFold(cName[1:len(cName)], name) {
						err = client.StopContainer(img.ID, 0)
						if err != nil {
							responseObj.ResponseCode = 0
							responseObj.ResponseDesc = "Error getting Remove Running Container"
							return responseObj, err
						} else {
							client.RemoveContainer(docker.RemoveContainerOptions{ID: img.ID})
							responseObj.ResponseCode = 0
							responseObj.ResponseDesc = "Successfully Remove Running Container"
							return responseObj, err
						}
					}
				}
			}
		}
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "No Running Container for Remove"
		return responseObj, err
	}
	return responseObj, err
}
