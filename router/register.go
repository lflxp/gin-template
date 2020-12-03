package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lflxp/gin-template/controller/auth"
	"github.com/lflxp/gin-template/controller/demo"
	"github.com/lflxp/gin-template/controller/menu"
	"github.com/lflxp/gin-template/controller/role"
	_ "github.com/lflxp/gin-template/docs"
	"github.com/lflxp/gin-template/middlewares"
	jwt "github.com/lflxp/gin-template/middlewares"
	"github.com/lflxp/gin-template/middlewares/permission"
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
	// admin.RegisterAdmin(r)
	// 注册demo接口
	demo.RegisterDemo(r)

	apiV1 := r.Group("/api/v1")

	apiV1.Use(jwt.JWT()) // token 验证
	apiV1.GET("/userInfo", auth.GetUserInfo)
	apiV1.Use(permission.CasbinMiddleware()) // 权限  验证

	{

		apiV1.GET("/menus", menu.GetMenus)
		apiV1.POST("/menus", menu.AddMenu)
		apiV1.PUT("/menus/:id", menu.EditMenu)
		apiV1.DELETE("/menus/:id", menu.DeleteMenu)

		apiV1.GET("/roles", role.GetRoles)
		apiV1.POST("/roles", role.AddRole)
		apiV1.PUT("/roles/:id", role.EditRole)
		apiV1.DELETE("/roles/:id", role.DeleteRole)

		apiV1.GET("/users", auth.GetUsers)
		apiV1.POST("/users", auth.AddUser)
		apiV1.PUT("/users/:id", auth.EditUser)
		apiV1.DELETE("/users/:id", auth.DeleteUser)
	}
}
