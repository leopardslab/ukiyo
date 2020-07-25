package schedulerapilayer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ukiyo/internal/containerscheduler"
	"ukiyo/internal/util"
	"ukiyo/pkg/jencoder"
)

func SaveRepositoryScheduledTime(r *gin.Engine) {
	r.POST("/save-repository-scheduled-time", func(c *gin.Context) {
		var podsDetails containerscheduler.PodsDetails
		var responseObj containerscheduler.ResponseObj
		c.ShouldBindJSON(&podsDetails)
		log.Println("Pods save details" + jencoder.PrintJson(podsDetails))

		if len(podsDetails.Name) > 0 && util.BindPortValidator(podsDetails.BindingPort) {
			res := containerscheduler.InsertPodsData(podsDetails)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		c.JSON(http.StatusOK, responseObj)
	})
}

func EditRepositoryScheduledTime(r *gin.Engine) {
	r.POST("/edit-repository-scheduled-time", func(c *gin.Context) {
		var podsDetails containerscheduler.PodsDetails
		var responseObj containerscheduler.ResponseObj
		c.ShouldBindJSON(&podsDetails)
		log.Println("Pods Update details" + jencoder.PrintJson(podsDetails))

		if len(podsDetails.Name) > 0 && util.BindPortValidator(podsDetails.BindingPort) {
			res := containerscheduler.UpdatePodsData(podsDetails)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		c.JSON(http.StatusOK, responseObj)
	})
}

func DeleteRepositoryScheduledTime(r *gin.Engine) {
	r.DELETE("/remove-repository-scheduled-time/:name", func(c *gin.Context) {
		var responseObj containerscheduler.ResponseObj
		name := c.Param("name")
		log.Println("Pod Delete details name=" + name)
		if len(name) > 0 {
			res := containerscheduler.DeleteDockerRegData(name)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		c.JSON(http.StatusOK, responseObj)
	})
}
