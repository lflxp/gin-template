package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	mjwt "github.com/lflxp/gin-template/middlewares/jwt"
	"github.com/lflxp/gin-template/model"
)

func RegisterClaims(router *gin.Engine) {
	var claimsMiddleware = mjwt.NewGinJwtMiddlewares(mjwt.AllUserAuthorizator)
	// claims CRUD
	apiGroup := router.Group("/api/v1/admin")
	apiGroup.Use(claimsMiddleware.MiddlewareFunc())
	{
		apiGroup.GET("/claims/all", GetAllClaims)
		apiGroup.GET("/claims/get", GetAllClaimsByAuth)
		apiGroup.POST("/claims/add", AddClaims)
		apiGroup.PUT("/claims/update/:id", PutClaims)
		apiGroup.DELETE("/claims/del/:id", DelClaims)
	}
}

// @Summary  获取当前用户所有claims
// @Description 获取当前用户所有claims
// @Tags Claims
// @Success 200 {string} string "success"
// @Router /api/v1/admin/claims/all [get]
func GetAllClaims(c *gin.Context) {
	data, err := model.GetClaims()
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status": true,
		"data":   data,
	})
}

// @Summary  获取当前用户所有claims
// @Description 获取当前用户所有claims
// @Tags Claims
// @Param auth query string false "指定用户"
// @Success 200 {string} string "success"
// @Router /api/v1/admin/claims/get [get]
func GetAllClaimsByAuth(c *gin.Context) {
	var (
		data []model.Claims
		err  error
	)

	auth := c.DefaultQuery("auth", "")
	if auth == "" {
		claims := jwt.ExtractClaims(c)
		data, err = model.GetClaimsByAuth(claims["userName"].(string))
		if err != nil {
			c.JSONP(200, gin.H{
				"status": false,
				"msg":    err.Error(),
			})
			return
		}
	} else {
		data, err = model.GetClaimsByAuth(auth)
		if err != nil {
			c.JSONP(200, gin.H{
				"status": false,
				"msg":    err.Error(),
			})
			return
		}
	}

	c.JSONP(200, gin.H{
		"status": true,
		"data":   data,
	})
}

// @Summary  新增用户
// @Description content 新增用户，不包括全权限
// @Tags Claims
// @Param data body model.Claims true "data"
// @Success 200 {object} model.Claims model.Claims{}
// @Router /api/v1/admin/claims/add [post]
func AddClaims(c *gin.Context) {
	var data model.Claims
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.AddClaims(&data)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status": true,
		"msg":    fmt.Sprintf("success add %d", num),
	})
}

// @Summary  更新Claims
// @Description content 修改用户名或密码
// @Tags Claims
// @Param id path string true "更新的目标claims id"
// @Param data body model.Claims true "data"
// @Success 200 {object} model.Claims model.Claims{}
// @Router /api/v1/admin/claims/update/{id} [put]
func PutClaims(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}
	var data map[string]interface{}
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.UpdateClaims(id, data)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status": true,
		"msg":    fmt.Sprintf("success update %d", num),
	})
}

// @Summary  删除Claims
// @Description 删除Claims
// @Tags Claims
// @Param id path string true "更新的目标claims id"
// @Success 200 {string} string success!
// @Router /api/v1/admin/claims/del/{id} [delete]
func DelClaims(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}

	num, err := model.DeleteClaims(id)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status": true,
		"msg":    fmt.Sprintf("success delete %d", num),
	})
}
