package dockerrunner

import (
	"encoding/json"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
	"ukiyo/pkg/jencoder"
)

type ResponseObj struct {
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

func ContainerRunner(containerId string, c *cache.Cache) (ResponseObj, error) {
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
		log.Println("Container Run failure :" + containerId)
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Container Run failure"
	} else {
		log.Println("Successfully Run The Container :" + containerId)
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Run The Container"
	}

	SetHistory(responseObj, c)
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
