package test

import (
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
	"ukiyo/pkg/manager/dockerpull"
	"ukiyo/pkg/webhook"
)

func TestPullToDocker(t *testing.T) {
	c := cache.New(1*time.Minute, 1*time.Minute)
	var webHook = webhook.Response{
		RepoName:  "demo-nginx",
		Namespace: "165090",
		Name:      "demo-nginx",
		Tag:       "02",
		PushedAt:  0,
	}
	docker, _ := dockerpull.PullToDocker(webHook, c)
	if docker.ResponseCode != 0 {
		t.Errorf("failed expected %v but got value %v", 0, docker.ResponseCode)
	} else {
		t.Logf("passed expected  %v and got value %v", 0, docker.ResponseCode)
	}
}
