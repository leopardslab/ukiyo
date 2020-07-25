package scheduler

import (
	"log"
	"time"
	"ukiyo/internal/containerscheduler"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/manager/dockercreater"
	"ukiyo/pkg/manager/dockerrunner"
)

func ScheduledRunner() {
	pods := containerscheduler.QueryListRecodeInDB()
	log.Println("ScheduledRunner trigger ...." + jencoder.PrintJson(pods))
	time.AfterFunc(1*time.Minute, func() {
		Foo()
	})
}

func WebHookScheduler(name string, imageName string) {
	pod := containerscheduler.QueryRecodeInDB(name)
	if len(pod.Name) > 0 && !pod.ScheduledDowntime {
		log.Println("WebHookScheduler trigger ...." + jencoder.PrintJson(pod))
		res, _ := dockercreater.ContainerCreate(name, imageName)
		log.Println("WebHookScheduler trigger ...." + jencoder.PrintJson(res))

		log.Println("Starting Container runner - WebHookScheduler")
		resObj, _ := dockerrunner.ContainerRunner(res.ContainerId)
		log.Println("Ending Container runner - WebHookScheduler ...." + jencoder.PrintJson(resObj))
	} else {
		log.Println("No saved pod details to schedule images")
	}
}

func Foo() {
	log.Println("Foo run for more than a minute.")
}
