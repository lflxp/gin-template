package auth

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	user_service "github.com/lflxp/gin-template/controller/service"
	"github.com/lflxp/gin-template/utils"
	"github.com/lflxp/gin-template/utils/e"
)

type auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}

// @Summary   获取登录token 信息
// @Tags auth
// @Accept json
// @Produce  json
// @Param   body  body   models.AuthSwag   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /auth  [POST]
func Auth(c *gin.Context) {

	var reqInfo auth
	err := c.BindJSON(&reqInfo)
	//dataByte, _ := ioutil.ReadAll(c.Request.Body)
	//fsion := gofasion.NewFasion(string(dataByte))
	//fmt.Println(fsion.Get("username").ValueStr())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": nil,
		})
		return
	}

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.Username, 100, "username").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if valid.HasErrors() {
		utils.MarkErrors(valid.Errors)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_ADD_FAIL),
			"data": valid.Errors,
		})
		return
	}

	authService := user_service.User{Username: reqInfo.Username, Password: reqInfo.Password}
	isExist, err := authService.Check()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
			"data": nil,
		})
		return
	}

	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  e.GetMsg(e.ERROR_AUTH),
			"data": nil,
		})
		return
	}

	user, err := authService.Get()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_AUTH_TOKEN),
			"data": nil,
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, reqInfo.Username, reqInfo.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  e.GetMsg(e.ERROR_AUTH_TOKEN),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": map[string]string{
			"token": token,
		},
	})
}
