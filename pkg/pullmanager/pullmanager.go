package pullmanager

import (
	"encoding/json"
	"log"
	"ukiyo/pkg/util"
)

func PullToDocker(str util.PullObj) {
	b, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("pull images : " + string(b))
}
