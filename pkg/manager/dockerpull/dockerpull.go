package dockerpull

import (
	docker "github.com/fsouza/go-dockerclient"
	"log"
	"strings"
	"ukiyo/internal/containeraccess"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/webhook"
)

type ResponseObj struct {
	Name         string
	ImageName    string
	ResponseCode int
	ResponseDesc string
}

var val int
var ImageName string

func PullToDocker(str webhook.Response) (ResponseObj, error) {
	var responseObj ResponseObj
	log.Println("docker pull by repo name : " + jencoder.PrintJson(str))
	ImageName = str.RepoName + ":" + str.Tag

	client, err := docker.NewClientFromEnv()
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Creating Docker Client Env"
		return responseObj, err
	}

	err = client.PullImage(webhook.DockerPullImage(str), containeraccess.RequestLoginKeys(str.Namespace))
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error Pulling Images"
		return responseObj, err
	}

	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Error List images"
		return responseObj, err
	}

	for _, img := range imgs {
		for _, repo := range img.RepoTags {
			if strings.EqualFold(repo, ImageName) {
				val++
				log.Println("Successfully pull the images : " + repo)
			}
		}
	}

	if val > 0 {
		responseObj.ResponseCode = 0
		responseObj.Name = str.Name
		responseObj.ImageName = ImageName
		responseObj.ResponseDesc = "Successfully pull the images"
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Failed pull the images"
	}
	return responseObj, err
}
