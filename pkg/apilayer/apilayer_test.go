package apilayer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddContainer(t *testing.T) {
	r := gin.Default()
	AddContainer(r)

	registry := `{
					"username":"agentukiyo",
					"repoName":"Ukiyo Docker registry",
					"accessToken":"f44e334e-1440-4166-a16f-d8fc9d0eb188",
					"email":"hansika.16@itfac.mrt.ac.lk",
					"serverAddress":"http://docker.io/v1"	
					}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add-container", strings.NewReader(registry))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	log.Println(response)
	value, exists := response["ResponseCode"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, 0, value)
}

func TestEditContainer(t *testing.T) {
	r := gin.Default()
	EditContainerToken(r)

	registry := `{
					"username":"agentukiyo",
					"accessToken":"f44e334e-1440-4166-a16f-d8fc9d0eb188",
					}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add-container", strings.NewReader(registry))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	log.Println(response)
	value, exists := response["ResponseCode"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, 0, value)
}