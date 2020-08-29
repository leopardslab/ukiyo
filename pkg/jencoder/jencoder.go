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

func PassJson(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
