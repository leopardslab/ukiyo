package webhooks

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"ukiyo/pkg/pullmanager"
	//"ukiyo/pkg/pushmanager"
	"ukiyo/pkg/util"

	"github.com/gin-gonic/gin"
)

type PushData struct {
	Pushed_at int      `json:"pushed_at"`
	images    []string `json:"images"`
	Tag       string   `json:"tag"`
	Pusher    string   `json:"pusher"`
}

type Repository struct {
	Status          string `json:"status"`
	Description     string `json:"description"`
	IsTrusted       bool   `json:"is_trusted"`
	FullDescription string `json:"full_description"`
	RepoUrl         string `json:"repo_url"`
	Owner           string `json:"owner"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	StarCount       int    `json:"star_count"`
	CommentCount    int    `json:"comment_count"`
	DateCreated     int    `json:"date_created"`
	RepoName        string `json:"repo_name"`
}

type DockerWebHook struct {
	PushData    PushData   `json:"push_data"`
	CallbackUrl string     `json:"callback_url"`
	Repository  Repository `json:"repository"`
}

var responseCode int
var responseDesc string
var imageName string

func HooksListener(r *gin.Engine) {
	r.POST("/ukiyo-web-hook", func(c *gin.Context) {
		var dockerWebHook DockerWebHook
		c.ShouldBindJSON(&dockerWebHook)

		b, err := json.Marshal(dockerWebHook)
		if err != nil {
			log.Println(err)
			return
		}

		var pullObj util.PullObj
		pullObj.Namespace = dockerWebHook.Repository.Namespace
		pullObj.RepoName = dockerWebHook.Repository.RepoName
		pullObj.Tag = dockerWebHook.PushData.Tag
		pullObj.PushedDate = dockerWebHook.PushData.Pushed_at

		log.Println("web-hook trigger" + string(b))

		imageName, responseCode, responseDesc, err = pullmanager.PullToDocker(pullObj)
		log.Println("pull Manager responseCode :" + strconv.Itoa(responseCode) + " responseDesc : " + responseDesc)

		c.String(http.StatusOK, "OK")
	})
}

func HealthCheck(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
