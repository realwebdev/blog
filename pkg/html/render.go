package html

import (
	"github.com/gin-gonic/gin"
	"github.com/realwebdev/blog/internal/providers/view"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)

	c.HTML(code, name, data)
}
