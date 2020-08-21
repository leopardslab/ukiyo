package cacheupdate

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"strconv"
	"time"
	"ukiyo/pkg/jencoder"
)

type CashObj struct {
	ImageName   string `json:"imageName"`
	ScheduledAt int64  `json:"scheduledAt"`
}

const (
	_TimeZone = "Asia/Kolkata"
)

var val int
var mnt int
var imageName string

func CacheUpdate(name string, imageName string, c *cache.Cache) {
	var cashObj CashObj
	if x, y, found := c.GetWithExpiration(name); found {
		err := json.Unmarshal(jencoder.PassJson(x), &cashObj)
		if err != nil {
			log.Println(err)
		}
		cashObj.ImageName = imageName
		c.Set(name, cashObj, time.Duration(time.Time(y).Sub(time.Now()).Minutes()+1)*time.Minute)
		log.Println("Successfully Cash update for the deployment")
	}
}

func TimeCalculator(name string, c *cache.Cache) (string, string, string) {
	var cashObj CashObj
	loc, _ := time.LoadLocation(_TimeZone)
	if x, _, found := c.GetWithExpiration(name); found {
		err := json.Unmarshal(jencoder.PassJson(x), &cashObj)
		if err != nil {
			log.Println(err)
		}
		val = int(cashObj.ScheduledAt-time.Now().In(loc).UnixNano()/int64(time.Millisecond)) + 60000
		if val > 0 {
			mnt = val / 60000
			if len(cashObj.ImageName) > 0 {
				imageName = cashObj.ImageName
			} else {
				imageName = "waiting"
			}
			return strconv.Itoa(mnt), strconv.Itoa((val - mnt*60000) / 1000), imageName
		} else {
			return "-", "-", "-"
		}
	} else {
		return "-", "-", "-"
	}
}
