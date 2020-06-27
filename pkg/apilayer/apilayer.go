package apilayer

import (
	"encoding/json"
	"log"
	"net/http"
	"ukiyo/pkg/auth"
	"ukiyo/pkg/util"

	"github.com/gin-gonic/gin"
)

func AddContainer(r *gin.Engine) {
	r.POST("/add-container", func(c *gin.Context) {
		var dockerRegistry auth.DockerRegistry
		var responseObj util.ResponseObj
		c.ShouldBindJSON(&dockerRegistry)

		b, err := json.Marshal(dockerRegistry)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Requested container registry insert details" + string(b))

		if len(dockerRegistry.Username) > 0 && len(dockerRegistry.AccessToken) > 0 && len(dockerRegistry.ServerAddress) > 0 {
			res := auth.InsertDockerRegData(dockerRegistry)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		c.JSON(http.StatusOK, responseObj)
	})
}

func EditContainerToken(r *gin.Engine) {
	r.POST("/edit-container-token", func(c *gin.Context) {
		var registryUpdate auth.RegistryUpdate
		var responseObj util.ResponseObj
		c.ShouldBindJSON(&registryUpdate)

		b, err := json.Marshal(registryUpdate)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Requested container registry update details" + string(b))

		if len(registryUpdate.Username) > 0 && len(registryUpdate.AccessToken) > 0 {
			res := auth.UpdateDockerRegData(registryUpdate)
			responseObj = res
		} else {
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Invalid Parameter"
		}
		c.JSON(http.StatusOK, responseObj)
	})
}
