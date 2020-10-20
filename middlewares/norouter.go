package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 404 handler
func NoRouteHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "%s", "Page Not Found")
}
