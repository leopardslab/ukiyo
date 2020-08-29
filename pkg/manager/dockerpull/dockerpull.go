package dockerpull

import (
	"encoding/json"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
	"time"
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

type History struct {
	Time    string
	Action  string
	Status  int
	Comment string
}

const (
	_TimeZone = "Asia/Kolkata"
)

var val int
var ImageName string

func PullToDocker(str webhook.Response, cash *cache.Cache) (ResponseObj, error) {
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

	SetHistory(responseObj, cash)
	return responseObj, err
}

func SetHistory(obj ResponseObj, cash *cache.Cache) {
	var historyArray []History
	var history History
	loc, _ := time.LoadLocation(_TimeZone)
	history.Time = time.Now().In(loc).Format("2006-01-02 15:04:05")
	history.Status = obj.ResponseCode
	history.Action = obj.ResponseDesc
	history.Comment = obj.ResponseDesc

	if x, _, found := cash.GetWithExpiration("history"); found {
		err := json.Unmarshal(jencoder.PassJson(x), &historyArray)
		if err != nil {
			log.Println(err)
		}
	}

	historyArray = append(historyArray, history)
	cash.Set("history", historyArray, 5*time.Minute)
}
