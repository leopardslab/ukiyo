package main

import (
	"ukiyo/pkg/webhooks"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	webhooks.HooksListener(r)
	webhooks.HealthCheck(r)
	r.Run(":8080")
}
