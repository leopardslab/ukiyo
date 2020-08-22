package util

import (
	"log"
	"ukiyo/internal/containerscheduler"
)

func BindPortValidator(bindingPort []containerscheduler.BindingPort, name string) bool {

	pods := containerscheduler.QueryListRecodeInDB()
	var portArr []string
	if len(pods) > 0 {
		for _, pod := range pods {
			if len(pod.BindingPort) > 0 && pod.Name != name {
				for _, bindingPort := range pod.BindingPort {
					portArr = append(portArr, bindingPort.ExportPort)
				}
			}
		}
	}
	log.Println(portArr)

	if len(bindingPort) > 0 {
		for _, ports := range bindingPort {
			if len(ports.InternalPort) == 0 && len(ports.ExportPort) == 0 {
				return false
			}
			if contains(portArr, ports.ExportPort) {
				return false
			}
		}
		return true
	}
	return false
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
