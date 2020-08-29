package eventsheduler

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
	"ukiyo/internal/caching"
	"ukiyo/internal/containerscheduler"
	"ukiyo/pkg/jencoder"
)

const (
	_TimeZone = "Asia/Kolkata"
)

func ScheduledRunner(event string, name string, ScheduledAt int64, c *cache.Cache) {
	switch env := event; env {
	case "run":
		RunEvent(c)
	case "save":
		SaveEvent(name, ScheduledAt, c)
	case "edit":
		EditEvent(name, ScheduledAt, c)
	case "delete":
		DeleteEvent(name, c)
	default:
		log.Println("No Matching Event in ScheduledRunner")
	}

}

func RunEvent(c *cache.Cache) {
	log.Println("RunEvent")
	pods := containerscheduler.QueryListRecodeInDB()
	log.Println("ScheduledRunner trigger ...." + jencoder.PrintJson(pods))

	if len(pods) > 0 {
		for _, pod := range pods {
			if len(pod.Name) > 0 && pod.ScheduledDowntime {
				loc, _ := time.LoadLocation(_TimeZone)
				log.Print(time.Now().In(loc).UnixNano() / int64(time.Millisecond))
				if (pod.ScheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) >= 0 {
					log.Println("Activated scheduling :" + pod.Name)
					caching.CacheRunner(pod.Name, "", pod.ScheduledAt, c)
				} else {
					log.Println("Time passed for the scheduling :" + pod.Name)
				}
			} else {
				log.Println("Not activate scheduling :" + pod.Name)
			}
		}
	}
	log.Println(c.Items())
}

func SaveEvent(name string, scheduledAt int64, c *cache.Cache) {
	log.Println("SaveEvent")
	log.Println(c.Items())
	loc, _ := time.LoadLocation(_TimeZone)
	if (scheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) >= 0 {
		caching.CacheRunner(name, "", scheduledAt, c)
	} else {
		log.Println("Time passed for the scheduling :" + name)
	}
	log.Println(c.Items())
}

func EditEvent(name string, scheduledAt int64, c *cache.Cache) {
	log.Println("EditEvent")
	log.Println(c.Items())
	caching.CacheRunner(name, "", scheduledAt, c)
	log.Println(c.Items())
}

func DeleteEvent(name string, c *cache.Cache) {
	log.Println("DeleteEvent")
	log.Println(c.Items())
	c.Delete(name)
	log.Println(c.Items())
}
