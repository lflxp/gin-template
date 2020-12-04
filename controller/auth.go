package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	newJwt "github.com/lflxp/gin-template/middlewares/jwt"
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func RegisterAuth(router *gin.Engine) {
	apiGroup := router.Group("/auth")
	// login
	apiGroup.POST("/login", Login)
	apiGroup.GET("/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": "success",
		})
	})

	var authMiddleware = newJwt.NewGinJwtMiddlewares(newJwt.AllUserAuthorizator)
	authGroup := router.Group("/auth")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		// Refresh time can be longer than token timeout
		authGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
		authGroup.GET("/hello", helloHandler)
	}
}

// @Summary  通用接口
// @Description 登陆、swagger、注销、404等
// @Tags Auth
// @Param token query string false "token"
// @Param data body model.Auth true "data"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var authMiddleware = newJwt.NewGinJwtMiddlewares(newJwt.AllUserAuthorizator)
	authMiddleware.LoginHandler(c)
}
