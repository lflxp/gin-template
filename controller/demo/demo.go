package demo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/lflxp/gin-template/middlewares"
	"github.com/lflxp/gin-template/model"
)

func RegisterDemo(router *gin.Engine) {
	demoGroup := router.Group("/api/v1/demo")
	demoGroup.Use(middlewares.JWT())
	{
		demoGroup.GET("/get", getDemo)
		demoGroup.POST("/add", addDemo)
		demoGroup.PUT("/put/:id", putDemo)
		demoGroup.DELETE("/del/:id", delDemo)
	}
}

// @Summary  查询指定key的值
// @Description 获取value，逐级数据查询
// @Tags Demo
// @Param token query string false "token"
// @Param key query string true "查询key"
// @Success 200 {object} model.Demo model.Demo{}
// @Router /api/v1/demo/get [get]
func getDemo(c *gin.Context) {
	key, isok := c.GetQuery("key")
	if !isok {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    "key is none",
		})
	}

	demo := new(model.Demo)
	t, err := demo.GetByString(key)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
			"source": t,
		})
	} else {
		c.JSONP(200, gin.H{
			"status": true,
			"data":   demo,
			"source": t,
		})
	}
}

// @Summary  新增Demo记录
// @Description 新增记录，只针对持久化数据
// @Tags Demo
// @Param token query string false "token"
// @Param data body model.Demo true "数据"
// @Success 200 {string} string "success"
// @Router /api/v1/demo/add [post]
func addDemo(c *gin.Context) {
	var data model.Demo
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.AddDemo(&data)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
	} else {
		c.JSONP(200, gin.H{
			"status":  true,
			"message": fmt.Sprintf("success add %d", num),
		})
	}
}

// @Summary  修改Demo记录
// @Description 修改指定id的内容
// @Tags Demo
// @Param token query string false "token"
// @Param id path string true "更新的id"
// @Param data body model.Demo true "数据"
// @Success 200 {string} string "success"
// @Router /api/v1/demo/put/{id} [put]
func putDemo(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}

	var data model.Demo
	info, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(info, &data)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	num, err := model.UpdateDemo(id, &data)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
	} else {
		c.JSONP(200, gin.H{
			"status":  true,
			"message": fmt.Sprintf("success put %d", num),
		})
	}
}

// @Summary  删除Demo记录
// @Description 删除记录
// @Tags Demo
// @Param token query string false "token"
// @Param id path string true "要删除的id"
// @Success 200 {string} string "success"
// @Router /api/v1/demo/del/{id} [delete]
func delDemo(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.String(200, "id is none")
		return
	}

	num, err := model.DelDemo(id)
	if err != nil {
		c.JSONP(200, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
	} else {
		c.JSONP(200, gin.H{
			"status":  true,
			"message": fmt.Sprintf("success del %d", num),
		})
	}
}
