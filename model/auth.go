package model

import (
	"errors"
	"fmt"

	"github.com/lflxp/gin-template/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("初始化表 Login")
	utils.Engine.Sync2(new(Auth), new(Claims))
	err := checkAdminUser()
	if err != nil {
		log.Error(err)
	}
}

// 用户表
type Auth struct {
	Id           int64  `json:"id"`
	Username     string `form:"username" json:"username" binding:"required" xorm:"varchar(255) notnull index unique"`
	Password     string `form:"password" json:"password" binding:"required" xorm:"varchar(255) not null"`
	Name         string `json:"name" xorm:"varchar(255)"`
	Avatar       string `json:"avatar" xorm:"varchar(255)"`
	Status       bool   `json:"status" xorm:"bool"`
	Telephone    string `json:"telephone" xorm:"varchar(255)"`
	LastLoginIp  string `json:"lastLoginIp" xorm:"varchar(255)"`
	CreateTime   string `json:"createTime" xorm:"varchar(255)"`
	CreatorId    string `json:"creatorId" xorm:"varchar(255)"`
	MerchantCode string `json:"merchantCode" xorm:"varchar(255)"`
	Deleted      bool   `json:"deleted" xorm:"bool"`
	RoleId       string `json:"roleId" xorm:"varchar(255)"`
}

// 检查是否有admin用户
// 如果不存在则在初始化的时候生成一个
func checkAdminUser() error {
	data := new(Auth)
	has, err := utils.Engine.Where("username = ?", "admin").Get(data)
	if err != nil {
		return err
	}

	if !has {
		data.Username = "admin"
		data.Password = utils.Jiami("admin")
		data.RoleId = "admin"
		data.Name = "管理员"
		data.Status = true
		num, err := AddAuth(data)
		if err != nil {
			return err
		}
		log.Infof("admin用户新增完毕，添加记录条数: %d", num)
	}
	return nil
}

func GetAuthByUsername(user string) (*Auth, bool, error) {
	data := new(Auth)
	has, err := utils.Engine.Where("username = ?", user).Get(data)
	return data, has, err
}

// 修改用户
func UpdateAuth(id string, data *Auth) (int64, error) {
	// if data.Password != "" {
	// 	data.Password = utils.Jiami(data.Password)
	// }
	affected, err := utils.Engine.Table(new(Auth)).ID(id).Update(data)
	return affected, err
}

// 验证用户名和密码是否正确
func VerifyAuth(username, pwd string) (bool, error) {
	var auth Auth
	has, err := utils.Engine.Where("username = ?", username).Get(&auth)
	if err != nil {
		return false, err
	}
	if has {
		// 服务端md5处理
		if utils.Jiami(pwd) == auth.Password {
			// 前端md5处理
			// if pwd == auth.Password {
			return true, nil
		}
		return false, errors.New("user or pwd is error")
	}
	return false, errors.New("user or pwd is not right")
}

// 查询用户
func GetAuth() ([]Auth, error) {
	data := make([]Auth, 0)
	err := utils.Engine.Desc("id").Find(&data)
	return data, err
}

// 新增用户
func AddAuth(data *Auth) (int64, error) {
	if data.Username == "" || data.Password == "" {
		return -1, errors.New("user or pwd is none")
	}
	// data.Password = utils.Jiami(data.Password)
	affected, err := utils.Engine.Insert(data)
	return affected, err
}

// 删除用户
func DeleteAuth(id string) (int64, error) {
	auth := new(Auth)
	affected, err := utils.Engine.ID(id).Delete(auth)
	return affected, err
}

// 用户权限表
type Claims struct {
	Id    int64  `json:"id"`
	Auth  string `json:"auth" xorm:"varchar(255) unique(only)"`  // 对应Auth => Username  eg: admin
	Type  string `json:"type" xorm:"varchar(255) unique(only)"`  // 权限类型 eg: nav
	Value string `json:"value" xorm:"varchar(255) unique(only)"` // 权限指 eg: dashboard
}

// 查询用户权限
func GetClaimsByAuthAndType(auth, types string) ([]Claims, error) {
	data := make([]Claims, 0)
	err := utils.Engine.Where("auth = ? and type = ?", auth, types).Desc("id").Find(&data)
	return data, err
}

// 查询用户权限
func GetClaimsByAuth(auth string) ([]Claims, error) {
	data := make([]Claims, 0)
	err := utils.Engine.Where("auth = ?", auth).Desc("id").Find(&data)
	return data, err
}

// 查询所有权限
func GetClaims() ([]Claims, error) {
	data := make([]Claims, 0)
	err := utils.Engine.Desc("id").Find(&data)
	return data, err
}

// 新增用户权限
func AddClaims(data *Claims) (int64, error) {
	affected, err := utils.Engine.Insert(data)
	return affected, err
}

// 删除用户权限
func DeleteClaims(id string) (int64, error) {
	c := new(Claims)
	affected, err := utils.Engine.ID(id).Delete(c)
	return affected, err
}

// 修改用户权限
func UpdateClaims(id string, data map[string]interface{}) (int64, error) {
	affected, err := utils.Engine.Table(new(Claims)).ID(id).Update(data)
	return affected, err
}

// 用户表权限
type User struct {
	Auth     // 用户基本信息
	Username string
	Claims   []Claims `json:"claims"` // 接口权限
	Role     Role     `json:"role"`   // 界面权限
}

func GetUser(username string) (*User, error) {
	var auth Auth
	user := &User{
		Username: username,
	}
	has, err := utils.Engine.Where("username = ?", username).Get(&auth)
	if err != nil {
		return user, err
	}
	if has {
		user.Auth = auth
		user.Password = "******"
		claims, err := GetClaimsByAuth(auth.Username)
		if err != nil {
			return user, err
		}
		user.Claims = claims
		return user, nil
	}
	return user, errors.New(fmt.Sprintf("%s not such user", username))
}

func GetUserClaims(username string) ([]Claims, error) {
	var auth Auth
	has, err := utils.Engine.Where("username = ?", username).Get(&auth)
	if err != nil {
		return nil, err
	}
	if has {
		claims, err := GetClaimsByAuth(auth.Username)
		if err != nil {
			return claims, err
		}
		return claims, nil
	}
	return nil, errors.New(fmt.Sprintf("%s not such user", username))
}
