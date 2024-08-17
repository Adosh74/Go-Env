package view

import "github.com/gin-gonic/gin"

func WithGlobalData(data gin.H) gin.H {
	data["APP_NAME"] = "Geek Blog"

	return data
}
