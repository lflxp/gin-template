package jwt

import (
	"encoding/json"
	"time"

	"github.com/lflxp/gin-template/model"
	log "github.com/sirupsen/logrus"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "thisiskeyword"

type IdentityKey string

// 接口权限
type JwtAuthorizator func(data interface{}, c *gin.Context) bool

// 根据不同接口的权限规则生成不同权限的jwt中间件
func NewGinJwtMiddlewares(jwta JwtAuthorizator) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gateway zone",
		Key:         []byte("github.com/lflxp/gin-template"),
		Timeout:     7 * 24 * time.Hour,
		MaxRefresh:  7 * 24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				//get claims from username
				v.Claims, _ = model.GetUserClaims(v.Username)
				jsonClaim, _ := json.Marshal(v.Claims)
				// maps the claims in the JWT
				return jwt.MapClaims{
					"userName":   v.Username,
					"userClaims": string(jsonClaim),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			//extracts identity from claims
			jsonClaim := claims["userClaims"].(string)
			var userClaims []model.Claims
			json.Unmarshal([]byte(jsonClaim), &userClaims)
			//Set the identity
			return &model.User{
				Username: claims["userName"].(string),
				Claims:   userClaims,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals model.Auth
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if ok, err := model.VerifyAuth(userID, password); err == nil {
				if ok {
					return &model.User{
						Username: userID,
					}, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwta,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

//role is admin can access
func AdminAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*model.User); ok {
		// for _, itemClaim := range v.Claims {
		// 	if itemClaim.Type == "role" && itemClaim.Value == "admin" {
		// 		return true
		// 	}
		// }

		if v.Username == "admin" {
			return true
		}
	}

	return false
}

//username is test can access
func TestAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*model.User); ok && v.Username == "test" {
		return true
	}

	return false
}

// 不限制用户权限
func AllUserAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

//404 handler
func NoRouteHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": 404, "message": "Page Not Found"})
}
