package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthMiddleware(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
