package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
	"ukiyo/api/historyapilayer"
	"ukiyo/api/registryapilayer"
	"ukiyo/api/schedulerapilayer"
	"ukiyo/pkg/scheduler/eventsheduler"
	"ukiyo/pkg/webhook-listener"
)

func main() {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	webhook_listener.HealthCheck(r)
	webhook_listener.HooksListener(r, c)
	registryapilayer.SaveContainerAccessKeys(r)
	registryapilayer.EditContainerAccessKeys(r)
	registryapilayer.DeleteContainerAccessKeys(r)
	schedulerapilayer.SaveRepositoryScheduledTime(r, c)
	schedulerapilayer.EditRepositoryScheduledTime(r, c)
	schedulerapilayer.DeleteRepositoryScheduledTime(r, c)
	historyapilayer.GetAllContainerHistory(r)
	historyapilayer.GetAllContainerHistoryByName(r)
	eventsheduler.ScheduledRunner("run", "", 0, c)
	r.Run(":8080")
}
