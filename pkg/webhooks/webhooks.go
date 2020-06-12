package webhooks

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PushData struct {
	Pushed_at int      `json:"pushed_at"`
	images    []string `json:"images"`
	Tag       string   `json:"tag"`
	Pusher    string   `json:"pusher"`
}

type Repository struct {
	Status           string `json:"status"`
	Description      string `json:"description"`
	Is_trusted       bool   `json:"is_trusted"`
	Full_description string `json:"full_description"`
	Repo_url         string `json:"repo_url"`
	Owner            string `json:"owner"`
	Is_official      bool   `json:"is_official"`
	Is_private       bool   `json:"is_private"`
	Name             string `json:"name"`
	Namespace        string `json:"namespace"`
	Star_count       int    `json:"star_count"`
	Comment_count    int    `json:"comment_count"`
	Date_created     int    `json:"date_created"`
	Repo_name        string `json:"repo_name"`
}

type DockerWebHook struct {
	PushData     PushData   `json:"push_data"`
	Callback_url string     `json:"callback_url"`
	Repository   Repository `json:"repository"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/ukiyo-web-hook", func(c *gin.Context) {
		var dockerWebHook DockerWebHook
		c.ShouldBindJSON(&dockerWebHook)
		b, err := json.Marshal(dockerWebHook)

		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(b))
		c.JSON(http.StatusOK, "foo")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
