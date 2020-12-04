package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lflxp/gin-template/controller"
	_ "github.com/lflxp/gin-template/docs"
	"github.com/lflxp/gin-template/middlewares"
	log "github.com/sirupsen/logrus"
)

// 注册插件和路由
func PreGinServe(r *gin.Engine) {
	log.Info("注册Gin路由")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	// swagger
	middlewares.RegisterSwaggerMiddleware(r)

	// 添加prometheus监控
	middlewares.RegisterPrometheusMiddleware(r, false)

	// 404
	r.NoRoute(middlewares.NoRouteHandler)

	// 健康检查
	r.GET("/health", middlewares.RegisterHealthMiddleware)

	// 注册admin接口
	controller.RegisterAdmin(r)
	// 注册demo接口
	controller.RegisterDemo(r)
	// 注册claims接口
	controller.RegisterClaims(r)
	// 注册auth接口
	controller.RegisterAuth(r)
	// 注册nav接口
	controller.RegisterPageRight(r)
}
