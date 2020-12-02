package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lflxp/gin-template/controller/admin"
	"github.com/lflxp/gin-template/controller/auth"
	"github.com/lflxp/gin-template/controller/demo"
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

	// 获取登录token
	r.POST("/auth", auth.Auth)
	// 注册admin接口
	admin.RegisterAdmin(r)
	// 注册demo接口
	demo.RegisterDemo(r)
}
