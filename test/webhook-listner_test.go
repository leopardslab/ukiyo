package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	webHookListener "ukiyo/pkg/webhook-listener"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()
	webHookListener.HealthCheck(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestHookRoute(t *testing.T) {
	r := gin.Default()
	webHookListener.HooksListener(r)

	pushJson := `{
                 	"push_data": {
                 		"pushed_at": 1591175168,
                 		"images": [],
                 		"tag": "02",
                 		"pusher": "registryName"
                 	},
                 	"callback_url": "https://registry.hub.docker.com/u/registryName/demo-nginx/hook/22fbde3h00hi54a3jdf42j5cf44bebjf1/",
                 	"repository": {
                 		"status": "Active",
                 		"description": "",
                 		"is_trusted": false,
                 		"full_description": "",
                 		"repo_url": "https://hub.docker.com/r/registryName/demo-nginx",
                 		"owner": "registryName",
                 		"is_official": false,
                 		"is_private": false,
                 		"name": "demo-nginx",
                 		"namespace": "registryName",
                 		"star_count": 0,
                 		"comment_count": 0,
                 		"date_created": 1555438658,
                 		"repo_name": "registryName/demo-nginx"
                 	}
                 }`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ukiyo-web-hook", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
