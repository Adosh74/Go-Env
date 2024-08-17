package routes

import (
	"net/http"

	"github.com/Adosh74/Go-Env/pkg/html"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home page",
		})
	})
	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About page",
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"APP_NAME": viper.Get("App.Name"),
		})
	})
}
