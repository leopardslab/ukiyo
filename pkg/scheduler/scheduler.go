package scheduler

import (
	"github.com/patrickmn/go-cache"
	"log"
	"ukiyo/internal/caching/cacheupdate"
	"ukiyo/internal/containerscheduler"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/manager/dockercreater"
	"ukiyo/pkg/manager/dockerremove"
	"ukiyo/pkg/manager/dockerrunner"
)

func WebHookScheduler(name string, imageName string, c *cache.Cache) {
	pod := containerscheduler.QueryRecodeInDB(name)
	log.Println("WebHookScheduler " + jencoder.PrintJson(pod))
	if len(pod.Name) > 0 && !pod.ScheduledDowntime {
		log.Println("Starting Non scheduled deployment")
		DeploymentProcess(name, imageName)
	} else if len(pod.Name) > 0 && pod.ScheduledDowntime {
		log.Println("Starting scheduled deployment")
		cacheupdate.CacheUpdate(name, imageName, c)
	} else {
		log.Println("Error pod details to schedule images")
	}
}

func DeploymentProcess(name string, imageName string) {
	log.Println("Starting Container Remove - WebHookScheduler")
	removeObj, _ := dockerremove.RemoveRunningContainer(name)
	log.Println("Ending Container Remove - WebHookScheduler ...." + jencoder.PrintJson(removeObj))

	if removeObj.ResponseCode == 0 {

		log.Println("WebHookScheduler trigger ContainerCreate...." + jencoder.PrintJson(name))
		res, _ := dockercreater.ContainerCreate(name, imageName)
		log.Println("WebHookScheduler trigger ContainerCreate...." + jencoder.PrintJson(res))

		log.Println("Starting Container runner - WebHookScheduler")
		resObj, _ := dockerrunner.ContainerRunner(res.ContainerId)
		log.Println("Ending Container runner - WebHookScheduler ...." + jencoder.PrintJson(resObj))

	} else {
		log.Println("Stop Container runner - WebHookScheduler - Failed Running Container remove process")
	}
}
