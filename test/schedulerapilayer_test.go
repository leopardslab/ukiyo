package test

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"ukiyo/api/schedulerapilayer"
)

func TestSaveRepositoryScheduledTime(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	schedulerapilayer.SaveRepositoryScheduledTime(r, c)
	pushJson := `{
				"name": "demo-nginx",
				"bindingPort": [{
					"exportPort": "8180",
					"internalPort": "80"
					},
					{
					"exportPort": "443",
					"internalPort": "443"
					}],
				"scheduledAt": "1555438658",
				"scheduledDowntime": false
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/save-repository-scheduled-time", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestEditRepositoryScheduledTime(t *testing.T) {
	r := gin.Default()
	schedulerapilayer.EditRepositoryScheduledTime(r)
	pushJson := `{
				"name": "demo-nginx",
				"bindingPort": [{
					"exportPort": "8180",
					"internalPort": "80"
					}],
				"scheduledAt": "1555438658",
				"scheduledDowntime": false
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/edit-repository-scheduled-time", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestDeleteRepositoryScheduledTime(t *testing.T) {
	r := gin.Default()
	schedulerapilayer.DeleteRepositoryScheduledTime(r)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/remove-repository-scheduled-time/demo-nginx", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
