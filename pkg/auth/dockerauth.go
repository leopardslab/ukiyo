package auth

import (
	"log"
	"strings"
	"ukiyo/pkg/util"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/sonyarouje/simdb/db"
)

type DockerRegistry struct {
	Username      string `json:"userName"`
	RepoName      string `json:"repoName"`
	AccessToken   string `json:"accessToken"`
	Email         string `json:"email"`
	ServerAddress string `json:"serverAddress"`
}

type RegistryUpdate struct {
	Username    string `json:"userName"`
	AccessToken string `json:"accessToken"`
}

var val int

func DockerLogin(userName string) docker.AuthConfiguration {

	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}

	var registries DockerRegistry
	err = driver.Open(DockerRegistry{}).Where("userName", "=", userName).First().AsEntity(&registries)
	if err != nil {
		panic(err)
	}

	if len(registries.Username) == 0 {
		panic("User's Registry Credentials is not exist. : " + userName)
	}

	return docker.AuthConfiguration{
		Username:      registries.Username,
		Password:      registries.AccessToken,
		Email:         registries.Email,
		ServerAddress: registries.ServerAddress,
	}
}

func InsertDockerRegData(dockerRegistry DockerRegistry) util.ResponseObj {
	var responseObj util.ResponseObj

	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}

	var registries []DockerRegistry
	err = driver.Open(DockerRegistry{}).Where("userName", "=", dockerRegistry.Username).Get().AsEntity(&registries)
	if err != nil {
		panic(err)
	}

	for _, registry := range registries {
		if strings.EqualFold(registry.Username, dockerRegistry.Username) {
			val = val + 1
			responseObj.ResponseCode = 1
			responseObj.ResponseDesc = "Already Exist"
			log.Println("Already Exist The Registry : " + registry.Username)
		}
	}

	if val == 0 {
		err = driver.Insert(dockerRegistry)
		if err != nil {
			panic(err)
		}
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Added"
	}

	return responseObj
}

func UpdateDockerRegData(registryUpdate RegistryUpdate) util.ResponseObj {
	var responseObj util.ResponseObj

	driver, err := db.New("dbs")
	if err != nil {
		panic(err)
	}

	var registries DockerRegistry
	err = driver.Open(DockerRegistry{}).Where("userName", "=", registryUpdate.Username).First().AsEntity(&registries)
	if err != nil {
		panic(err)
	}

	registries.AccessToken = registryUpdate.AccessToken
	err = driver.Update(registries)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Update Failed"
		panic(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Updated"
	}

	return responseObj
}

func (c DockerRegistry) ID() (jsonField string, value interface{}) {
	value = c.Username
	jsonField = "userName"
	return
}
