package containerscheduler

import (
	"log"
	"strings"
	"ukiyo/internal/dbconfig"
)

type PodsDetails struct {
	Name              string        `json:"name"`
	BindingPort       []BindingPort `json:"bindingPort"`
	ScheduledAt       int64         `json:"scheduledAt"`
	ScheduledDowntime bool          `json:"scheduledDowntime"`
}

type BindingPort struct {
	ExportPort   string `json:"exportPort"`
	InternalPort string `json:"internalPort"`
}

type ResponseObj struct {
	ResponseCode int
	ResponseDesc string
}

const (
	_podsPk = "name"
)

var val int
var buildTimeMilSec int64
var buildTimeMnt int64

type CashObj struct {
	ImageName   string `json:"imageName"`
	ScheduledAt int64  `json:"scheduledAt"`
}

func (c PodsDetails) ID() (jsonField string, value interface{}) {
	value = c.Name
	jsonField = _podsPk
	return
}

func GetBindingPortsForContainerCreate(name string) []string {
	var data []string
	pod := QueryRecodeInDB(name)
	if len(pod.BindingPort) > 0 {
		for _, port := range pod.BindingPort {
			data = append(data, port.ExportPort+":"+port.InternalPort)
		}
	}
	return data
}

func InsertPodsData(podsDetails PodsDetails) ResponseObj {
	var responseObj ResponseObj
	var pods = QueryListRecodeInDB()
	if len(pods) > 0 {
		for _, pod := range pods {
			if strings.EqualFold(pod.Name, podsDetails.Name) {
				val++
				responseObj.ResponseCode = 1
				responseObj.ResponseDesc = "Already Exist Pod details"
				log.Println("Already Exist The Pods : " + pod.Name)
			}
		}
	}
	if val == 0 {
		responseObj = InsertDb(podsDetails)
	}
	return responseObj
}

func UpdatePodsData(podsDetails PodsDetails) ResponseObj {
	var responseObj ResponseObj

	var pod = QueryRecodeInDB(podsDetails.Name)

	if len(pod.Name) != 0 {
		pod.BindingPort = podsDetails.BindingPort
		pod.ScheduledAt = podsDetails.ScheduledAt
		pod.ScheduledDowntime = podsDetails.ScheduledDowntime
		responseObj = UpdateDb(pod)
	} else {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Update Failed pod."
	}

	return responseObj
}

func DeleteDockerRegData(name string) ResponseObj {
	var responseObj ResponseObj

	var pod = QueryRecodeInDB(name)

	if len(pod.Name) != 0 {
		responseObj = DeleteDb(pod)
	} else {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Delete Failed pod."
	}
	return responseObj
}

func QueryListRecodeInDB() []PodsDetails {
	var pods []PodsDetails
	err := dbconfig.DbConfig().Open(PodsDetails{}).Get().AsEntity(&pods)
	if err != nil {
		log.Println(err)
		return pods
	}
	return pods
}

func QueryRecodeInDB(Name string) PodsDetails {
	var pod PodsDetails
	err := dbconfig.DbConfig().Open(PodsDetails{}).Where(_podsPk, "=", Name).First().AsEntity(&pod)
	if err != nil {
		log.Println(err)
		return pod
	}
	return pod
}

func InsertDb(pods PodsDetails) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Insert(pods)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Insert Failed pod details"
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Added pod details"
	}
	return responseObj
}

func UpdateDb(pod PodsDetails) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Update(pod)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Update Failed pod details"
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Updated pod details"
	}
	return responseObj
}

func DeleteDb(pod PodsDetails) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Delete(pod)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Delete Failed pod details"
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Deleted pod details"
	}
	return responseObj
}
