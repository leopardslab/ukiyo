package containeraccess

import (
	"github.com/fsouza/go-dockerclient"
	"log"
	"strings"
	"ukiyo/internal/dbconfig"
)

type ContainerKeys struct {
	UserName      string `json:"userName"`
	Desc          string `json:"desc"`
	AccessToken   string `json:"accessToken"`
	Email         string `json:"email"`
	ServerAddress string `json:"serverAddress"`
}

type ResponseObj struct {
	ResponseCode int
	ResponseDesc string
}

const (
	_dockerRegPk = "userName"
)

func (c ContainerKeys) ID() (jsonField string, value interface{}) {
	value = c.UserName
	jsonField = _dockerRegPk
	return
}

var val int

func RequestLoginKeys(userName string) docker.AuthConfiguration {
	var registries = QueryRecodeInDB(userName)

	if len(registries.UserName) == 0 {
		panic("User's Registry Credentials is not exist. : " + userName)
	}

	return docker.AuthConfiguration{
		Username:      registries.UserName,
		Password:      registries.AccessToken,
		Email:         registries.Email,
		ServerAddress: registries.ServerAddress,
	}
}

func InsertDockerRegData(containerKeys ContainerKeys) ResponseObj {
	var responseObj ResponseObj
	var registries = QueryListRecodeInDB(containerKeys.UserName)
	if len(registries) > 0 {
		for _, registry := range registries {
			if strings.EqualFold(registry.UserName, containerKeys.UserName) {
				val++
				responseObj.ResponseCode = 1
				responseObj.ResponseDesc = "Already Exist Container key :" + containerKeys.UserName
				log.Println("Already Exist The Container key: " + registry.UserName)
			}
		}
	}
	if val == 0 {
		responseObj = InsertDb(containerKeys)
	}
	return responseObj
}

func UpdateDockerRegData(containerKeys ContainerKeys) ResponseObj {
	var responseObj ResponseObj

	var registries = QueryRecodeInDB(containerKeys.UserName)

	log.Println(registries)

	if len(registries.UserName) != 0 {
		registries.Desc = containerKeys.Desc
		registries.AccessToken = containerKeys.AccessToken
		registries.Email = containerKeys.Email
		registries.ServerAddress = containerKeys.ServerAddress
		responseObj = UpdateDb(registries)
	} else {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Update Failed Container Key :" + containerKeys.UserName
	}

	return responseObj
}

func DeleteDockerRegData(userName string) ResponseObj {
	var responseObj ResponseObj

	var registries = QueryRecodeInDB(userName)

	if len(registries.UserName) != 0 {
		responseObj = DeleteDb(registries)
	} else {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Delete Failed Container key :" + userName
	}
	return responseObj
}

func QueryListRecodeInDB(Username string) []ContainerKeys {
	var registries []ContainerKeys
	err := dbconfig.DbConfig().Open(ContainerKeys{}).Where(_dockerRegPk, "=", Username).Get().AsEntity(&registries)
	if err != nil {
		log.Println(err)
		return registries
	}
	return registries
}

func QueryRecodeInDB(Username string) ContainerKeys {
	var registries ContainerKeys
	err := dbconfig.DbConfig().Open(ContainerKeys{}).Where(_dockerRegPk, "=", Username).First().AsEntity(&registries)
	if err != nil {
		log.Println(err)
		return registries
	}
	return registries
}

func InsertDb(containerKeys ContainerKeys) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Insert(containerKeys)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Insert Failed Container key : " + containerKeys.UserName
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Added Container Key :" + containerKeys.UserName
	}
	return responseObj
}

func UpdateDb(containerKeys ContainerKeys) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Update(containerKeys)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Update Failed Container key :" + containerKeys.UserName
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Updated Container key :" + containerKeys.UserName
	}
	return responseObj
}

func DeleteDb(containerKeys ContainerKeys) ResponseObj {
	var responseObj ResponseObj
	err := dbconfig.DbConfig().Delete(containerKeys)
	if err != nil {
		responseObj.ResponseCode = 1
		responseObj.ResponseDesc = "Delete Failed Container Key :" + containerKeys.UserName
		log.Println(err)
	} else {
		responseObj.ResponseCode = 0
		responseObj.ResponseDesc = "Successfully Deleted Container Key :" + containerKeys.UserName
	}
	return responseObj
}
