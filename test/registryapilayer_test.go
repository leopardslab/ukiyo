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
	"ukiyo/api/registryapilayer"
)

func TestSaveContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	registryapilayer.SaveContainerAccessKeys(r, c)
	pushJson := `{
				"username":"registry_name",
				"desc":"Docker registry keys",
				"accessToken":"290e587a-1790-41d7-a053-61a2ef377875",
				"email":"username@gmail.com",
				"serverAddress":"http://docker.io/v1"
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/save-container-access-keys", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Added Container Key :registry_name\"}", w.Body.String())
}

func TestEditContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	registryapilayer.EditContainerAccessKeys(r, c)
	pushJson := `{
				"username":"registry_name",
				"desc":"Docker registry keys",
				"accessToken":"290e587a-1790-41d7-a053-61a2ef377875",
				"email":"username@gmail.com",
				"serverAddress":"http://docker.io/v1"
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/edit-container-access-keys", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Updated Container key :registry_name\"}", w.Body.String())
}

func TestDeleteContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	c := cache.New(1*time.Minute, 1*time.Minute)
	registryapilayer.DeleteContainerAccessKeys(r, c)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/delete-container-access-keys/registry_name", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ResponseCode\":0,\"ResponseDesc\":\"Successfully Deleted Container Key :registry_name\"}", w.Body.String())
}