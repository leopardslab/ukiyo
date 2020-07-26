package jencoder

import (
	"encoding/json"
	"log"
)

func PrintJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return "json marshal error"
	}
	return string(b)
}
