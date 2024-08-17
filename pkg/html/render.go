package html

import (
	"github.com/Adosh74/Go-Env/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(data)
	c.HTML(code, name, data)
}
