package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/lflxp/gin-template/middlewares"
	"github.com/lflxp/gin-template/model"
)

func RegisterAdmin(router *gin.Engine) {
	// auth CRUD
	apiGroup := router.Group("/api/v1/admin")
	apiGroup.Use(middlewares.JWT())
	{
		apiGroup.GET("/auth/all", GetAllAuth)
		apiGroup.POST("/auth/add", AddAuth)
		apiGroup.PUT("/auth/update/:id", PutAuth)
		apiGroup.DELETE("/auth/del/:id", DelAuth)
		apiGroup.GET("/auth/verify/:user/:pwd", GetAuthVerify)
	}
}

// @Summary  获取所有auth
// @Description 获取所有auth
// @Tags Admin
// @Param token query string false "token"
// @Param user path string true "用户名"
// @Param pwd path string true "密码"
// @Success 200 {string} string "success"
// @Router /api/v1/admin/auth/verify/{user}/{pwd} [get]
func GetAuthVerify(c *gin.Context) {
	user := c.Params.ByName("user")
	if user == "" {
		c.String(200, "user is none")
		return
	}

	pwd := c.Params.ByName("pwd")
	if pwd == "" {
		c.String(200, "pwd is none")
		return
	}

	isok, err := model.VerifyAuth(user, pwd)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	if isok {
		c.JSONP(200, gin.H{
			"status": true,
			"data":   "success",
		})
	} else {
		c.JSONP(200, gin.H{
			"status": false,
			"data":   "failed",
		})
	}

}

// @Summary  获取所有auth
// @Description 获取所有auth
// @Tags Admin
// @Param token query string false "token"
// @Success 200 {string} string "success"
// @Router /api/v1/admin/auth/all [get]
func GetAllAuth(c *gin.Context) {
	data, err := model.GetAuth()
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

// @Summary  新增用户
// @Description content 新增用户，不包括全权限
// @Tags Admin
// @Param token query string false "token"
// @Param data body model.Auth true "data"
// @Success 200 {object} model.Auth model.Auth{}
// @Router /api/v1/admin/auth/add [post]
func AddAuth(c *gin.Context) {
	var data model.Auth
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.AddAuth(&data)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status":  true,
		"message": fmt.Sprintf("success add %d", num),
	})
}

// @Summary  更新Auth
// @Description content 修改用户名或密码
// @Tags Admin
// @Param token query string false "token"
// @Param id path string true "更新的目标auth id"
// @Param data body model.Auth true "data"
// @Success 200 {object} model.Auth model.Auth{}
// @Router /api/v1/admin/auth/update/{id} [put]
func PutAuth(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}
	var data model.Auth
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.UpdateAuth(id, &data)
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

// @Summary  删除Auth
// @Description 删除Auth
// @Tags Admin
// @Param token query string false "token"
// @Param id path string true "更新的目标auth id"
// @Success 200 {string} string success!
// @Router /api/v1/admin/auth/del/{id} [delete]
func DelAuth(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}

	num, err := model.DeleteAuth(id)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSONP(200, gin.H{
		"status": true,
		"data":   fmt.Sprintf("success delete %d", num),
	})
}
