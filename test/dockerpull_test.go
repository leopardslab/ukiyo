package test

import (
	"testing"
	"ukiyo/pkg/manager/dockerpull"
	"ukiyo/pkg/webhook"
)

func TestPullToDocker(t *testing.T) {
	var webHook = webhook.Response{
		RepoName:  "demo-nginx",
		Namespace: "name",
		Name:      "demo-nginx",
		Tag:       "02",
		PushedAt:  0,
	}
	docker, _ := dockerpull.PullToDocker(webHook)
	if docker.ResponseCode != 0 {
		t.Errorf("failed expected %v but got value %v", 0, docker.ResponseCode)
	} else {
		t.Logf("passed expected  %v and got value %v", 0, docker.ResponseCode)
	}
}
