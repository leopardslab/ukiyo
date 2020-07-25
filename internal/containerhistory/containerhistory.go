package containerhistory

import (
	"log"
	"ukiyo/internal/containeraccess"
	"ukiyo/internal/containerscheduler"
	"ukiyo/internal/dbconfig"
)

type HistoryDetails struct {
	Name        string      `json:"name"`
	EventId     string      `json:"eventId"`
	EventObject EventObject `json:"eventObject"`
}

type EventObject struct {
	EventType string    `json:"eventType"`
	EventCode int       `json:"EventCode"`
	EventDesc string    `json:"eventDesc"`
	EventData EventData `json:"eventData"`
	EventAt   int       `json:"eventAt"`
}

type EventData struct {
	ContainerKeys containeraccess.ContainerKeys  `json:"containerKeys"`
	PodsDetails   containerscheduler.PodsDetails `json:"podsDetails"`
	ImagePuller   ImagePuller                    `json:"imagePuller"`
	DockerCreator DockerCreator                  `json:"dockerCreator"`
	DockerRunner  DockerRunner                   `json:"dockerRunner"`
}

type DockerCreator struct {
	Name        string `json:"name"`
	ContainerId string `json:"containerId"`
}

type ImagePuller struct {
	Name      string `json:"name"`
	ImageName string `json:"imageName"`
}

type DockerRunner struct {
	Name string `json:"name"`
}

type HistoryResponse struct {
	ResponseCode int
	ResponseDesc string
	PageNumbers  int
}

const (
	_podsPk                 = "name"
	_eventTypeContainerKeys = "Save Container Keys Event"
	_eventTypePodsDetails   = "Save Pod Details for Deployment Event"
	_eventTypeImagePuller   = "Pull Image Event"
	_eventTypeDockerCreator = "Create Docker Event"
	_eventTypeDockerRunner  = "Run Image Event"
)

func (c HistoryDetails) ID() (jsonField string, value interface{}) {
	value = c.Name
	jsonField = _podsPk
	return
}

func QueryAllHistoryRecodeInDB() []HistoryDetails {
	var history []HistoryDetails
	err := dbconfig.DbConfig().Open(HistoryDetails{}).Get().AsEntity(&history)
	if err != nil {
		log.Println(err)
		return history
	}
	return history
}

func QueryHistoryRecodeInDB(Name string) HistoryDetails {
	var history HistoryDetails
	err := dbconfig.DbConfig().Open(HistoryDetails{}).Where(_podsPk, "=", Name).First().AsEntity(&history)
	if err != nil {
		log.Println(err)
		return history
	}
	return history
}
