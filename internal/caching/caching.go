package caching

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
	"ukiyo/internal/containerscheduler"
	"ukiyo/pkg/jencoder"
)

var buildTimeMilSec int64
var buildTimeMnt int64

type CashObj struct {
	ImageName   string `json:"imageName"`
	ScheduledAt int64  `json:"scheduledAt"`
}

func ScheduledRunner() {
	pods := containerscheduler.QueryListRecodeInDB()
	log.Println("ScheduledRunner trigger ...." + jencoder.PrintJson(pods))

	if len(pods) > 0 {
		for _, pod := range pods {
			if len(pod.Name) > 0 && pod.ScheduledDowntime {
				loc, _ := time.LoadLocation("Asia/Kolkata")
				log.Print("current time")
				log.Print(time.Now().In(loc).UnixNano() / int64(time.Millisecond))
				if (pod.ScheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) >= 0 {
					log.Println("Activated scheduling :" + pod.Name)
					CacheRunner(pod.Name, "", pod.ScheduledAt)
				} else {
					log.Println("Time passed for the scheduling :" + pod.Name)
				}
			} else {
				log.Println("Not activate scheduling :" + pod.Name)
			}
		}
	}
}

func CacheRunner(name string, imageName string, scheduledAt int64) {
	var cashObj CashObj
	cashObj.ImageName = "FFF" //-->?
	cashObj.ScheduledAt = scheduledAt

	loc, _ := time.LoadLocation("Asia/Kolkata")
	buildTimeMilSec = scheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)
	buildTimeMnt = (buildTimeMilSec / 60000) + 1

	c := cache.New(1*time.Minute, 1*time.Minute)
	c.Set(name, cashObj, time.Duration(buildTimeMnt+2)*time.Minute) //--> 5

	log.Println(buildTimeMnt, buildTimeMnt+2)
	log.Println("Applied AfterFunc : " + name)
	log.Println(c.Items())

	time.AfterFunc(time.Duration(buildTimeMnt)*time.Minute, func() {
		ContainerRunFunc(name, scheduledAt, c)
	})

	for {
		log.Print("%%%")
		log.Print(c.Items())
		log.Print("$$$")
		time.Sleep(30 * time.Second)
	}

}

func ContainerRunFunc(name string, scheduledAt int64, c *cache.Cache) {

	log.Println(name + "$$$$$$$$$$")
	log.Println(scheduledAt)

	loc, _ := time.LoadLocation("Asia/Kolkata")
	if (scheduledAt-time.Now().In(loc).UnixNano()/int64(time.Millisecond)) > 60000 ||
		(scheduledAt-time.Now().In(loc).UnixNano()/int64(time.Millisecond)) < 60000 {

		log.Println("Trigger Start Container @ : " + name)

		var cashObj CashObj
		if x, found := c.Get(name); found {
			err := json.Unmarshal(jencoder.PassJson(x), &cashObj)
			if err != nil {
				log.Println(err)
			}
			log.Println(cashObj.ImageName)
			log.Println(cashObj.ScheduledAt)
		}

	} else {
		log.Println("Not Valid Time for Start Container")
	}

	for {
		log.Print("##")
		log.Print(c.Items())
		log.Print("**")
		time.Sleep(30 * time.Second)
	}

}
