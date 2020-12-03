package permission

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lflxp/gin-template/middlewares"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwtGet "github.com/lflxp/gin-template/utils"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")
		t, _ := jwt.Parse(token[1], func(*jwt.Token) (interface{}, error) {
			return jwtGet.JwtSecret, nil
		})
		fmt.Println(jwtGet.GetIdFromClaims("username", t.Claims), c.Request.URL.Path, c.Request.Method)

		if b, err := middlewares.Obj.Enforcer.EnforceSafe(jwtGet.GetIdFromClaims("username", t.Claims), c.Request.URL.Path, c.Request.Method); err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusOK,
				"data": err,
				"msg":  "ok",
			})
			c.Abort()
			return
		} else if !b {

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusForbidden,
				"data": "登录用户 没有权限",
				"msg":  "ok",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
