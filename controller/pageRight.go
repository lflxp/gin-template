package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	newJwt "github.com/lflxp/gin-template/middlewares/jwt"
	"github.com/lflxp/gin-template/model"
	log "github.com/sirupsen/logrus"
)

// 注册页面权限
func RegisterPageRight(router *gin.Engine) {
	var authMiddleware = newJwt.NewGinJwtMiddlewares(newJwt.AdminAuthorizator)
	// auth CRUD
	apiGroup := router.Group("/api/v1/nav")
	apiGroup.Use(authMiddleware.MiddlewareFunc())
	{
		apiGroup.GET("/get", GetNav)
		apiGroup.POST("/add", AddNav)
		apiGroup.PUT("/put/:id", NavPut)
		apiGroup.DELETE("/delete/:id", NavDelete)
	}
}

// @Summary Nav接口概要说明
// @Description GET获取所有模板
// @Tags Page
// @Success 200 {object} model.Nav model.Nav{} //成功返回的数据结构， 最后是示例
// @Security ApiKeyAuth
// @Router /api/v1/nav/get [get]
func GetNav(ctx *gin.Context) {
	data, err := model.GetNav()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}

// @Summary  新增接口
// @Description Nav POST ADD INTERFACE
// @Tags Page
// @Param data body model.Nav true "data"
// @Success 200 {object} model.Nav model.Nav{} //成功返回的数据结构， 最后是示例
// @Security ApiKeyAuth
// @Router /api/v1/nav/add [post]
func AddNav(c *gin.Context) {
	var temp model.Nav
	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Warning(string(data))
	err := json.Unmarshal(data, &temp)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		t, _ := json.Marshal(temp)
		log.Warning(string(t))
		n, err := model.AddNav(&temp)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, fmt.Sprintf("success add %d", n))
		}
	}
}

// @Summary  删除接口
// @Description DELETE INTERFACE
// @Tags Page
// @Param id path int true "ID"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /api/v1/nav/delete/{id} [delete]
func NavDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	n, err := model.DeleteNav(id)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, fmt.Sprintf("success delete %d", n))
	}
}

// @Summary  修改接口
// @Description UPDATE INTERFACE
// @Tags Page
// @Param id path int true "ID"
// @Param data body model.Nav true "data"
// @Success 200 {string} string "success"
// @Security ApiKeyAuth
// @Router /api/v1/nav/put/{id} [put]
func NavPut(c *gin.Context) {
	id := c.Params.ByName("id")
	var info model.Nav
	data, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(data, &info)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		n, err := model.UpdateNav(id, info)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, fmt.Sprintf("success update %d", n))
		}
	}
}
