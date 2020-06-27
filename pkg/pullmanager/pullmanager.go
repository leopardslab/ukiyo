package pullmanager

import (
	"encoding/json"
	"log"
	"strings"
	"ukiyo/pkg/auth"
	"ukiyo/pkg/util"

	docker "github.com/fsouza/go-dockerclient"
)

var val int
var imageName string

func PullToDocker(str util.PullObj) (string, int, string, error) {
	b, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
		return "", 1, "Invalid pull image Obj", err
	}
	imageName = str.RepoName + ":" + str.Tag
	log.Println("pull images : " + string(b))

	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	err = client.PullImage(DockerPullImage(str), auth.DockerLogin(str.Namespace))

	if err != nil {
		panic(err)
	}

	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})

	if err != nil {
		panic(err)
	}

	for _, img := range imgs {
		for _, repo := range img.RepoTags {
			if strings.EqualFold(repo, imageName) {
				val = val + 1
				log.Println("Successfully pull the images : " + repo)
			}
		}
	}

	if val > 0 {
		return imageName, 0, "Successfully pull the images", err
	} else {
		return "", 1, "Failed pull the images", err
	}
}

func DockerPullImage(str util.PullObj) docker.PullImageOptions {
	return docker.PullImageOptions{
		Repository:        str.RepoName,
		Tag:               str.Tag,
		Platform:          "",
		Registry:          "",
		OutputStream:      nil,
		RawJSONStream:     false,
		InactivityTimeout: 0,
		Context:           nil,
	}
}
