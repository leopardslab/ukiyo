package pullmanager

import (
	"testing"
	"ukiyo/pkg/util"

	"github.com/stretchr/testify/assert"
)

var responseCode int
var responseDesc string

func TestAddContainer(t *testing.T) {

	var pullObj util.PullObj
	pullObj.Namespace = "agentukiyo"
	pullObj.RepoName = "agentukiyo/ukiyo"
	pullObj.Tag = "02"
	pullObj.PushedDate = 2020121234
	imageName, responseCode, responseDesc, err := PullToDocker(pullObj)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "", imageName)
	assert.Equal(t, 0, responseCode)
	assert.Equal(t, "Successfully pull the images", responseDesc)
}
