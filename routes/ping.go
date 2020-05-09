package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Writer.WriteString("Ping")
}
