package containerhistory

import (
	"log"
	"ukiyo/internal/containeraccess"
	"ukiyo/internal/containerscheduler"
	"ukiyo/internal/dbconfig"
)

type HistoryDetails struct {
	Name string `json:"name"`
	Data []Data `json:"data"`
}

type Data struct {
	Id          int         `json:"id"`
	EventType   string      `json:"eventType"`
	EventObject EventObject `json:"eventObject"`
	EventData   EventData   `json:"eventData"`
}

type EventObject struct {
	EventCode int    `json:"EventCode"`
	EventDesc string `json:"eventDesc"`
	EventAt   int    `json:"eventAt"`
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
	ResponseCode   int
	ResponseDesc   string
	PageNumbers    int
	HistoryDetails HistoryDetails
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

func UpdateContainerHistory(name string, data Data) {
	var history = QueryHistoryRecodeInDB(name)
	if len(history.Name) > 0 {
		data.Id = len(history.Data) + 1
		history.Data = append(history.Data, data)
		UpdateDb(history)
	} else {
		var historyDetails HistoryDetails
		historyDetails.Name = name
		data.Id = 1
		historyDetails.Data = append(historyDetails.Data, data)
		InsertDb(historyDetails)
	}
}

func QueryAllHistoryRecodeInDB(pageNo string) HistoryResponse {
	var historyRes HistoryResponse
	var historyDetails HistoryDetails
	err := dbconfig.DbConfig().Open(HistoryDetails{}).First().AsEntity(&historyDetails)
	if err != nil {
		log.Println(err)
		historyRes.ResponseCode = 1
		historyRes.ResponseDesc = "Failed"
		return historyRes
	} else {
		historyRes.ResponseCode = 0
		historyRes.ResponseDesc = "Success"
		historyRes.HistoryDetails = historyDetails
		historyRes.PageNumbers = 1
	}
	return historyRes
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

func InsertDb(historyDetails HistoryDetails) {
	err := dbconfig.DbConfig().Insert(historyDetails)
	if err != nil {
		log.Println("Fail History Recode Insert")
	} else {
		log.Println("Successfully Installed History Recode")
	}
}

func UpdateDb(historyDetails HistoryDetails) {
	err := dbconfig.DbConfig().Update(historyDetails)
	if err != nil {
		log.Println("Fail History Recode Update")
	} else {
		log.Println("Successfully Updated History Recode")
	}
}
