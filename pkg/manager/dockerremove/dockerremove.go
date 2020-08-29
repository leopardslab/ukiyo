package dockerremove

import (
	"encoding/json"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
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

func RemoveRunningContainer(name string, cash *cache.Cache) (ResponseObj, error) {
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
