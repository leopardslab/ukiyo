package test

import (
	"testing"
	"ukiyo/pkg/manager/dockercreater"
)

func TestContainerCreate(t *testing.T) {
	docker, _ := dockercreater.ContainerCreate("demo-nginx", "demo-nginx")
	if docker.ResponseCode != 0 {
		t.Errorf("failed expected %v but got value %v", 0, docker.ResponseCode)
	} else {
		t.Logf("passed expected  %v and got value %v", 0, docker.ResponseCode)
	}
}
