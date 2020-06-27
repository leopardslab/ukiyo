package main

import (
	"ukiyo/pkg/apilayer"
	"ukiyo/pkg/webhooks"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	webhooks.HooksListener(r)
	webhooks.HealthCheck(r)
	apilayer.AddContainer(r)
	apilayer.EditContainerToken(r)
	r.Run(":8080")
}
