package test

import (
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
	"ukiyo/pkg/manager/dockercreater"
)

func TestContainerCreate(t *testing.T) {
	c := cache.New(1*time.Minute, 1*time.Minute)
	docker, _ := dockercreater.ContainerCreate("demo-nginx", "demo-nginx", c)
	if docker.ResponseCode != 0 {
		t.Errorf("failed expected %v but got value %v", 0, docker.ResponseCode)
	} else {
		t.Logf("passed expected  %v and got value %v", 0, docker.ResponseCode)
	}
}
