package old

import "github.com/gin-gonic/gin"

func FromContext(c *gin.Context) map[string][]string {
	c.Request.ParseForm()
	return c.Request.PostForm
}
