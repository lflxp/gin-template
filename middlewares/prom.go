// prometheus中间件
package middlewares

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterPrometheusMiddleware(router *gin.Engine, isauth bool) {
	if isauth {
		group := router.Group("/metrics", gin.BasicAuth(gin.Accounts{
			"root": "system",
		}))
		{
			group.GET("/", ginprom.PromHandler(promhttp.Handler()))
		}
	} else {
		router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	}
}
