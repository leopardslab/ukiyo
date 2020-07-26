package historyapilayer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ukiyo/internal/containerhistory"
)

func GetAllContainerHistory(r *gin.Engine) {
	r.GET("/get-container-history/:pageNo", func(c *gin.Context) {
		pageNo := c.Param("pageNo")
		response := containerhistory.QueryAllHistoryRecodeInDB(pageNo)
		log.Println("Container History details by pageNo=" + pageNo)
		c.JSON(http.StatusOK, response)
	})
}

func GetAllContainerHistoryByName(r *gin.Engine) {
	r.GET("/get-container-history-by-name/:name", func(c *gin.Context) {
		var response containerhistory.HistoryResponse
		name := c.Param("name")
		log.Println("Container History details by name=" + name)
		c.JSON(http.StatusOK, response)
	})
}
