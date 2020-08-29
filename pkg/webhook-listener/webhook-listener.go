package webhook_listener

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"ukiyo/pkg/jencoder"
	"ukiyo/pkg/manager/dockerpull"
	"ukiyo/pkg/scheduler"
	"ukiyo/pkg/webhook"
)

type PushData struct {
	PushedAt int      `json:"pushed_at"`
	Images   []string `json:"images"`
	Tag      string   `json:"tag"`
	Pusher   string   `json:"pusher"`
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

func HooksListener(r *gin.Engine, cash *cache.Cache) {
	r.POST("/ukiyo-web-hook", func(c *gin.Context) {
		var dockerWebHook DockerWebHook
		c.ShouldBindJSON(&dockerWebHook)
		log.Println("ukiyo web-hook trigger request" + jencoder.PrintJson(dockerWebHook))

		res := SetWebHookResponse(dockerWebHook)
		log.Println("ukiyo web-hook trigger response" + jencoder.PrintJson(res))

		resObj, _ := dockerpull.PullToDocker(res, cash)

		if len(resObj.Name) > 0 {
			scheduler.WebHookScheduler(resObj.Name, resObj.ImageName, cash)
		} else {
			log.Println("Error Creating Image .." + jencoder.PrintJson(resObj))
		}

		c.String(http.StatusOK, "OK")
	})
}

func HealthCheck(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func SetWebHookResponse(dockerWebHook DockerWebHook) webhook.Response {
	var response webhook.Response
	response.Namespace = dockerWebHook.Repository.Namespace
	response.RepoName = dockerWebHook.Repository.RepoName
	response.Name = dockerWebHook.Repository.Name
	response.Tag = dockerWebHook.PushData.Tag
	response.PushedAt = dockerWebHook.PushData.PushedAt
	return response
}
