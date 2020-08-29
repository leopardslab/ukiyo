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
				"scheduledTime": "Aug 17 2020 00:19:50 AM",
				"scheduledDowntime": false
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/save-repository-scheduled-time", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Added pod details\"}", w.Body.String())
}

func TestEditRepositoryScheduledTime(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	schedulerapilayer.EditRepositoryScheduledTime(r, c)
	pushJson := `{
				"name": "demo-nginx",
				"bindingPort": [{
					"exportPort": "8180",
					"internalPort": "80"
					}],
				"scheduledTime": "Aug 17 2020 00:19:50 AM",
				"scheduledDowntime": false
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/edit-repository-scheduled-time", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Updated pod details\"}", w.Body.String())
}

func TestDeleteRepositoryScheduledTime(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	schedulerapilayer.DeleteRepositoryScheduledTime(r, c)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/remove-repository-scheduled-time/demo-nginx", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Deleted pod details\"}", w.Body.String())
}