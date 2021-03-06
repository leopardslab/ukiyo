package webhooks

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()
	HealthCheck(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestHookRoute(t *testing.T) {
	r := gin.Default()
	HooksListener(r)

	pushJson := `{
                 	"push_data": {
                 		"pushed_at": 1591175168,
                 		"images": [],
                 		"tag": "02",
                 		"pusher": "dabare"
                 	},
                 	"callback_url": "https://registry.hub.docker.com/u/dabare/testing/hook/22fbde3h00hi54a3jdf42j5cf44bebjf1/",
                 	"repository": {
                 		"status": "Active",
                 		"description": "",
                 		"is_trusted": false,
                 		"full_description": "",
                 		"repo_url": "https://hub.docker.com/r/dabare/testing",
                 		"owner": "dabare",
                 		"is_official": false,
                 		"is_private": false,
                 		"name": "testing",
                 		"namespace": "dabare",
                 		"star_count": 0,
                 		"comment_count": 0,
                 		"date_created": 1555438658,
                 		"repo_name": "dabare/testing"
                 	}
                 }`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ukiyo-web-hook", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
