package caching

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/scheduler"
)

var buildTimeMilSec int64
var buildTimeMnt int64

const (
	_TimeZone = "Asia/Kolkata"
)

type CashObj struct {
	ImageName   string `json:"imageName"`
	ScheduledAt int64  `json:"scheduledAt"`
}

func CacheRunner(name string, imageName string, scheduledAt int64, c *cache.Cache) {
	var cashObj CashObj
	cashObj.ImageName = imageName
	cashObj.ScheduledAt = scheduledAt

	loc, _ := time.LoadLocation(_TimeZone)
	buildTimeMilSec = scheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)
	buildTimeMnt = (buildTimeMilSec / 60000) + 1

	c.Set(name, cashObj, time.Duration(buildTimeMnt+2)*time.Minute)

	log.Println(buildTimeMnt, buildTimeMnt+2)
	log.Println("Applied AfterFunc : " + name)
	log.Println(c.Items())

	time.AfterFunc(time.Duration(buildTimeMnt)*time.Minute, func() {
		ContainerRunFunc(name, scheduledAt, c)
	})
}

func ContainerRunFunc(name string, scheduledAt int64, c *cache.Cache) {
	loc, _ := time.LoadLocation(_TimeZone)
	if ((scheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) > 60000) ||
		((scheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) < 60000) {

		log.Println("Trigger Start Container @ : " + name)

		var cashObj CashObj
		if x, found := c.Get(name); found {
			err := json.Unmarshal(jencoder.PassJson(x), &cashObj)
			if err != nil {
				log.Println(err)
			}
			if len(cashObj.ImageName) > 0 {
				log.Println("Starting to create scheduled container : " + cashObj.ImageName)
				scheduler.DeploymentProcess(name, cashObj.ImageName)
			} else {
				log.Println("Failed to create scheduled container, No image in cache : null")
			}
		}

	} else {
		log.Println("Not Valid Time for Start Container")
	}
}
