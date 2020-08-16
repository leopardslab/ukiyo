package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"ukiyo/api/registryapilayer"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSaveContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	registryapilayer.SaveContainerAccessKeys(r)
	pushJson := `{
				"username":"name",
				"desc":"Docker registry keys",
				"accessToken":"290e587a-1790-41d7-a053-61a2ef377875",
				"email":"username@gmail.com",
				"serverAddress":"http://docker.io/v1"
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/save-container-access-keys", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestEditContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	registryapilayer.EditContainerAccessKeys(r)
	pushJson := `{
				"username":"name",
				"desc":"Docker registry keys",
				"accessToken":"290e587a-1790-41d7-a053-61a2ef377875",
				"email":"username@gmail.com",
				"serverAddress":"http://docker.io/v1"
				}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/edit-container-access-keys", strings.NewReader(pushJson))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestDeleteContainerAccessKeys(t *testing.T) {
	r := gin.Default()
	registryapilayer.DeleteContainerAccessKeys(r)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/delete-container-access-keys/registryName", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
