package role

import (
	"io/ioutil"
	"net/http"

	inject "github.com/lflxp/gin-template/middlewares"

	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/lflxp/gin-template/controller/service"
	"github.com/lflxp/gin-template/utils"
	"github.com/unknwon/com"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/lflxp/gin-template/model"

	"github.com/lflxp/gin-template/utils/e"
	"github.com/lflxp/gin-template/utils/setting"
)

// @Summary   获取所有角色
// @Tags role
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles  [GET]
func GetRoles(c *gin.Context) {
	RoleService := service.Role{
		ID:       com.StrTo(c.Query("id")).MustInt(),
		PageNum:  utils.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := RoleService.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_COUNT_FAIL),
			"data": nil,
		})
		return
	}

	articles, err := RoleService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_GET_S_FAIL),
			"data": nil,
		})
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  e.SUCCESS,
		"data": data,
	})
}

// @Summary   增加角色
// @Tags role
// @Accept json
// @Produce  json
// @Param   body  body   model.Role   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles  [POST]
func AddRole(c *gin.Context) {
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("username").ValueStr()
	menuId := com.StrTo(fsion.Get("menu_id").ValueInt()).MustInt()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "path").Message("名称最长为100字符")

	if valid.HasErrors() {
		utils.MarkErrors(valid.Errors)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_ADD_FAIL),
			"data": nil,
		})
		return
	}

	RoleService := service.Role{
		Name: name,
		Menu: menuId,
	}

	if id, err := RoleService.Add(); err != nil {

		err = inject.Obj.Common.RoleAPI.LoadPolicy(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  e.GetMsg(e.ERROR_EDIT_FAIL),
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  e.SUCCESS,
			"data": nil,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_ADD_FAIL),
			"data": nil,
		})
		return
	}

}

// @Summary   更新角色
// @Tags role
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Param   body  body   model.Role   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles/:id  [PUT]
func EditRole(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("username").ValueStr()
	menuId := com.StrTo(fsion.Get("menu_id").ValueInt()).MustInt()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "path").Message("名称最长为100字符")

	if valid.HasErrors() {
		utils.MarkErrors(valid.Errors)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_ADD_FAIL),
			"data": nil,
		})
		return
	}
	RoleService := service.Role{
		ID:   id,
		Name: name,
		Menu: menuId,
	}
	exists, err := RoleService.ExistByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_EXIST_FAIL),
			"data": nil,
		})
		return
	}
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  e.ERROR_EXIST_FAIL,
			"data": nil,
		})
		return
	}

	err = RoleService.Edit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_EDIT_FAIL),
			"data": nil,
		})
		return
	}

	err = inject.Obj.Common.RoleAPI.LoadPolicy(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_EDIT_FAIL),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  e.SUCCESS,
		"data": nil,
	})
}

// @Summary   删除角色
// @Tags role
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles/:id  [DELETE]
func DeleteRole(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		utils.MarkErrors(valid.Errors)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": nil,
		})
		return
	}

	RoleService := service.Role{ID: id}
	exists, err := RoleService.ExistByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_EXIST_FAIL),
			"data": nil,
		})
		return
	}
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  e.ERROR_EXIST_FAIL,
			"data": nil,
		})
		return
	}
	role, err := RoleService.Get()
	err = RoleService.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_DELETE_FAIL),
			"data": nil,
		})
		return
	}

	inject.Obj.Enforcer.DeleteUser(role.Name)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  e.SUCCESS,
		"data": nil,
	})
}
