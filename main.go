package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"time"
	"ukiyo/api/registryapilayer"
	"ukiyo/api/schedulerapilayer"
	"ukiyo/pkg/process"
	"ukiyo/pkg/scheduler/eventsheduler"
	"ukiyo/pkg/webhook-listener"
)

func main() {

	file, err := os.OpenFile("dbs/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	webhook_listener.HealthCheck(r)
	webhook_listener.HooksListener(r, c)
	registryapilayer.SaveContainerAccessKeys(r, c)
	registryapilayer.EditContainerAccessKeys(r, c)
	registryapilayer.DeleteContainerAccessKeys(r, c)
	schedulerapilayer.SaveRepositoryScheduledTime(r, c)
	schedulerapilayer.EditRepositoryScheduledTime(r, c)
	schedulerapilayer.DeleteRepositoryScheduledTime(r, c)
	eventsheduler.ScheduledRunner("run", "", 0, c)
	go process.Process(c)
	r.Run(":8080")
}
