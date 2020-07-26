package test

import (
	"testing"
	"ukiyo/pkg/manager/dockerrunner"
)

func TestContainerRunner(t *testing.T) {
	docker, _ := dockerrunner.ContainerRunner("6dc636d3c48dad91ee753440bd708888bf1b278040c6732e604b076ee44dd75f")
	if docker.ResponseCode != 0 {
		t.Errorf("failed expected %v but got value %v", 0, docker.ResponseCode)
	} else {
		t.Logf("passed expected  %v and got value %v", 0, docker.ResponseCode)
	}
}
