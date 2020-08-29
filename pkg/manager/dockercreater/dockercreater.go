package dockercreater

import (
	"encoding/json"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
	"time"
	"ukiyo/internal/containerscheduler"
	"ukiyo/pkg/jencoder"
)

type ResponseObj struct {
	Name         string
	ContainerId  string
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

func ContainerCreate(name string, imageName string, c *cache.Cache) (ResponseObj, error) {
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
