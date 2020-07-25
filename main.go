package main

import (
	"github.com/gin-gonic/gin"
	"ukiyo/api/historyapilayer"
	"ukiyo/api/registryapilayer"
	"ukiyo/api/schedulerapilayer"
	"ukiyo/pkg/scheduler"
	"ukiyo/pkg/webhook-listener"
)

func main() {
	r := gin.Default()
	webhook_listener.HealthCheck(r)
	webhook_listener.HooksListener(r)
	registryapilayer.SaveContainerAccessKeys(r)
	registryapilayer.EditContainerAccessKeys(r)
	registryapilayer.DeleteContainerAccessKeys(r)
	schedulerapilayer.SaveRepositoryScheduledTime(r)
	schedulerapilayer.EditRepositoryScheduledTime(r)
	schedulerapilayer.DeleteRepositoryScheduledTime(r)
	historyapilayer.GetAllContainerHistory(r)
	historyapilayer.GetAllContainerHistoryByName(r)
	scheduler.ScheduledRunner()
	r.Run(":8080")
}
