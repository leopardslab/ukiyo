package cacheupdate

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"strconv"
	"time"
	"ukiyo/pkg/jencoder"
)

const (
	_TimeZone = "Asia/Kolkata"
)

type CashObj struct {
	ImageName   string `json:"imageName"`
	ScheduledAt int64  `json:"scheduledAt"`
}

type History struct {
	Time    string
	Action  string
	Status  int
	Comment string
}

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

func TimeCalculator(name string, c *cache.Cache) (int, int, string) {
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
			return mnt, (val - mnt*60000) / 1000, imageName
		} else {
			return 0, 0, "-"
		}
	} else {
		return 0, 0, "-"
	}
}

func HistoryFetch(c *cache.Cache) [][]string {
	var historyArr []History
	var historyData [][]string
	if x, _, found := c.GetWithExpiration("history"); found {
		err := json.Unmarshal(jencoder.PassJson(x), &historyArr)
		if err != nil {
			log.Println(err)
		}
	}
	if len(historyArr) > 10 {
		for x := 0; x < (len(historyArr) - 10); x++ {
			historyArr = RemoveIndex(historyArr, 8)
		}
	}
	if len(historyArr) > 0 {
		for _, histor := range historyArr {
			historyData = append(historyData, []string{histor.Time, histor.Action, strconv.Itoa(histor.Status), histor.Comment})
		}
	}
	return historyData
}

func RemoveIndex(s []History, index int) []History {
	return append(s[:index], s[index+1:]...)
}
