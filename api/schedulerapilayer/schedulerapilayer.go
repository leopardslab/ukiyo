package schedulerapilayer

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
	"ukiyo/internal/containerscheduler"
	"ukiyo/internal/util"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/scheduler/eventsheduler"
)

const (
	_TimeZone = "Asia/Kolkata"
	_format   = "Jan _2 2006 3:04:05 PM"
)

type PodsDetailsObj struct {
	Name              string                           `json:"name"`
	BindingPort       []containerscheduler.BindingPort `json:"bindingPort"`
	ScheduledTime     string                           `json:"scheduledTime"`
	ScheduledDowntime bool                             `json:"scheduledDowntime"`
}

func SaveRepositoryScheduledTime(r *gin.Engine, cache *cache.Cache) {
	r.POST("/save-repository-scheduled-time", func(c *gin.Context) {
		var podsDetailsObj PodsDetailsObj
		var responseObj containerscheduler.ResponseObj
		c.ShouldBindJSON(&podsDetailsObj)
		log.Println("save-repository-scheduled-time | request : " + jencoder.PrintJson(podsDetailsObj))

		loc, _ := time.LoadLocation(_TimeZone)
		if len(podsDetailsObj.Name) > 0 && util.BindPortValidator(podsDetailsObj.BindingPort, "") {
			podsDetails, status := RequestDateConverter(podsDetailsObj)
			if podsDetails.ScheduledDowntime && status == 0 {
				if ((podsDetails.ScheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) / 60000) >= 5 {
					res := containerscheduler.InsertPodsData(podsDetails)
					if res.ResponseCode == 0 {
						eventsheduler.ScheduledRunner("save", podsDetails.Name, podsDetails.ScheduledAt, cache)
					}
					responseObj = res
				} else {
					responseObj.ResponseCode = 1
					responseObj.ResponseDesc = "Invalid Scheduled Time. Minimum 5Mnt Required"
				}
			} else if status == 0 {
				res := containerscheduler.InsertPodsData(podsDetails)
				responseObj = res
			} else {
				responseObj.ResponseCode = 1
				responseObj.ResponseDesc = "Invalid Parameter"
			}
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("save-repository-scheduled-time | response : " + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}

func EditRepositoryScheduledTime(r *gin.Engine, cache *cache.Cache) {
	r.POST("/edit-repository-scheduled-time", func(c *gin.Context) {
		var podsDetailsObj PodsDetailsObj
		var responseObj containerscheduler.ResponseObj
		c.ShouldBindJSON(&podsDetailsObj)
		log.Println("edit-repository-scheduled-time | request : " + jencoder.PrintJson(podsDetailsObj))

		if len(podsDetailsObj.Name) > 0 && util.BindPortValidator(podsDetailsObj.BindingPort, podsDetailsObj.Name) {
			podsDetails, status := RequestDateConverter(podsDetailsObj)
			loc, _ := time.LoadLocation(_TimeZone)
			if podsDetails.ScheduledDowntime && status == 0 {
				log.Println((podsDetails.ScheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) / 60000)
				if ((podsDetails.ScheduledAt - time.Now().In(loc).UnixNano()/int64(time.Millisecond)) / 60000) >= 5 {
					res := containerscheduler.UpdatePodsData(podsDetails)
					if res.ResponseCode == 0 {
						eventsheduler.ScheduledRunner("edit", podsDetails.Name, podsDetails.ScheduledAt, cache)
					}
					responseObj = res
				} else {
					responseObj.ResponseCode = 1
					responseObj.ResponseDesc = "Invalid Scheduled Time. Minimum 5Mnt Required"
				}
			} else if !podsDetails.ScheduledDowntime && status == 0 {
				res := containerscheduler.UpdatePodsData(podsDetails)
				if res.ResponseCode == 0 {
					eventsheduler.ScheduledRunner("delete", podsDetails.Name, 0, cache)
				}
				responseObj = res
			} else {
				responseObj.ResponseCode = 1
				responseObj.ResponseDesc = "Invalid Parameter."
			}
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("edit-repository-scheduled-time | response : " + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}

func DeleteRepositoryScheduledTime(r *gin.Engine, cache *cache.Cache) {
	r.DELETE("/remove-repository-scheduled-time/:name", func(c *gin.Context) {
		var responseObj containerscheduler.ResponseObj
		name := c.Param("name")
		log.Println("remove-repository-scheduled-time | request : name=" + name)
		if len(name) > 0 {
			res := containerscheduler.DeleteDockerRegData(name)
			if res.ResponseCode == 0 {
				eventsheduler.ScheduledRunner("delete", name, 0, cache)
			}
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("remove-repository-scheduled-time | response : " + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}

func RequestDateConverter(podsDetailsObj PodsDetailsObj) (containerscheduler.PodsDetails, int) {
	var podsDetails containerscheduler.PodsDetails
	podsDetails.Name = podsDetailsObj.Name
	podsDetails.BindingPort = podsDetailsObj.BindingPort
	podsDetails.ScheduledDowntime = podsDetailsObj.ScheduledDowntime
	loc, _ := time.LoadLocation(_TimeZone)
	time1, err := time.ParseInLocation(_format, podsDetailsObj.ScheduledTime, loc)
	if err != nil {
		log.Println("Error while parsing date :", err)
		return podsDetails, 1
	}
	podsDetails.ScheduledAt = time.Time(time1).In(loc).UnixNano() / int64(time.Millisecond)
	return podsDetails, 0
}
