package registryapilayer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
	"ukiyo/internal/containeraccess"
	"ukiyo/pkg/jencoder"
)

type History struct {
	Time    string
	Action  string
	Status  int
	Comment string
}

const (
	_TimeZone = "Asia/Kolkata"
)

func SaveContainerAccessKeys(r *gin.Engine, cache *cache.Cache) {
	r.POST("/save-container-access-keys", func(c *gin.Context) {
		var containerKey containeraccess.ContainerKeys
		var responseObj containeraccess.ResponseObj
		c.ShouldBindJSON(&containerKey)
		log.Println("save-container-access-keys | request : " + jencoder.PrintJson(containerKey))

		if len(containerKey.UserName) > 0 && len(containerKey.AccessToken) > 0 && len(containerKey.ServerAddress) > 0 {
			res := containeraccess.InsertDockerRegData(containerKey)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter : save-container key"
		}
		log.Println("save-container-access-keys | response : " + jencoder.PrintJson(responseObj))
		SetHistory(responseObj, cache)
		c.JSON(http.StatusOK, responseObj)
	})
}

func EditContainerAccessKeys(r *gin.Engine, cache *cache.Cache) {
	r.POST("/edit-container-access-keys", func(c *gin.Context) {
		var containerKey containeraccess.ContainerKeys
		var responseObj containeraccess.ResponseObj
		c.ShouldBindJSON(&containerKey)
		log.Println("edit-container-access-keys | request : " + jencoder.PrintJson(containerKey))

		if len(containerKey.UserName) > 0 && len(containerKey.AccessToken) > 0 {
			res := containeraccess.UpdateDockerRegData(containerKey)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter edit-container key"
		}
		log.Println("edit-container-access-keys | response :" + jencoder.PrintJson(responseObj))
		SetHistory(responseObj, cache)
		c.JSON(http.StatusOK, responseObj)
	})
}

func DeleteContainerAccessKeys(r *gin.Engine, cache *cache.Cache) {
	r.DELETE("/delete-container-access-keys/:userName", func(c *gin.Context) {
		var responseObj containeraccess.ResponseObj
		name := c.Param("userName")
		log.Println("delete-container-access-keys | request : userName=" + name)
		if len(name) > 0 {
			res := containeraccess.DeleteDockerRegData(name)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter delete-container key"
		}
		log.Println("delete-container-access-keys | response : " + jencoder.PrintJson(responseObj))
		SetHistory(responseObj, cache)
		c.JSON(http.StatusOK, responseObj)
	})
}

func SetHistory(obj containeraccess.ResponseObj, cash *cache.Cache) {
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
