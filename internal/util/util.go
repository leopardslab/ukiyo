package util

import (
	"ukiyo/internal/containerscheduler"
)

func BindPortValidator(bindingPort []containerscheduler.BindingPort) bool {
	if len(bindingPort) > 0 {
		for _, ports := range bindingPort {
			if len(ports.InternalPort) == 0 && len(ports.ExportPort) == 0 {
				return false
			}
		}
		return true
	}
	return false
}
