package registryapilayer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ukiyo/internal/containeraccess"
	"ukiyo/pkg/jencoder"
)

func SaveContainerAccessKeys(r *gin.Engine) {
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
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("save-container-access-keys | response : " + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}

func EditContainerAccessKeys(r *gin.Engine) {
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
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("edit-container-access-keys | response :" + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}

func DeleteContainerAccessKeys(r *gin.Engine) {
	r.DELETE("/delete-container-access-keys/:userName", func(c *gin.Context) {
		var responseObj containeraccess.ResponseObj
		name := c.Param("userName")
		log.Println("delete-container-access-keys | request : userName=" + name)
		if len(name) > 0 {
			res := containeraccess.DeleteDockerRegData(name)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		log.Println("delete-container-access-keys | response : " + jencoder.PrintJson(responseObj))
		c.JSON(http.StatusOK, responseObj)
	})
}
