package model

// import (
// 	"errors"

// 	"github.com/lflxp/gin-template/utils"
// 	log "github.com/sirupsen/logrus"
// )

// func init() {
// 	log.Info("初始化表 Login")
// 	utils.Engine.Sync2(new(Auth))
// 	err := checkAdminUser()
// 	if err != nil {
// 		log.Error(err)
// 	}
// }

// // 用户表
// type Auth struct {
// 	Id           int64  `json:"id"`
// 	Username     string `form:"username" json:"username" binding:"required" xorm:"varchar(255) notnull index unique"`
// 	Password     string `form:"password" json:"password" binding:"required" xorm:"varchar(255) not null"`
// 	Name         string `json:"name" xorm:"varchar(255)"`
// 	Avatar       string `json:"avatar" xorm:"varchar(255)"`
// 	Status       bool   `json:"status" xorm:"bool"`
// 	Telephone    string `json:"telephone" xorm:"varchar(255)"`
// 	LastLoginIp  string `json:"lastLoginIp" xorm:"varchar(255)"`
// 	CreateTime   string `json:"createTime" xorm:"varchar(255)"`
// 	CreatorId    string `json:"creatorId" xorm:"varchar(255)"`
// 	MerchantCode string `json:"merchantCode" xorm:"varchar(255)"`
// 	Deleted      bool   `json:"deleted" xorm:"bool"`
// 	RoleId       string `json:"roleId" xorm:"varchar(255)"`
// }

// // 检查是否有admin用户
// // 如果不存在则在初始化的时候生成一个
// func checkAdminUser() error {
// 	data := new(Auth)
// 	has, err := utils.Engine.Where("username = ?", "admin").Get(data)
// 	if err != nil {
// 		return err
// 	}

// 	if !has {
// 		tmp := utils.GetRandomSalt()
// 		data.Username = "admin"
// 		data.Password = utils.Jiami(tmp)
// 		data.RoleId = "admin"
// 		data.Name = "管理员"
// 		data.Status = true
// 		log.Infof("admin用户不存在，随机创建密码为： %s", tmp)
// 		num, err := AddAuth(data)
// 		if err != nil {
// 			return err
// 		}
// 		log.Infof("admin用户新增完毕，添加记录条数: %d", num)
// 	}
// 	return nil
// }

// func GetAuthByUsername(user string) (*Auth, bool, error) {
// 	data := new(Auth)
// 	has, err := utils.Engine.Where("username = ?", user).Get(data)
// 	return data, has, err
// }

// // 修改用户
// func UpdateAuth(id string, data *Auth) (int64, error) {
// 	// if data.Password != "" {
// 	// 	data.Password = utils.Jiami(data.Password)
// 	// }
// 	affected, err := utils.Engine.Table(new(Auth)).ID(id).Update(data)
// 	return affected, err
// }

// // 验证用户名和密码是否正确
// func VerifyAuth(username, pwd string) (bool, error) {
// 	var auth Auth
// 	has, err := utils.Engine.Where("username = ?", username).Get(&auth)
// 	if err != nil {
// 		return false, err
// 	}
// 	if has {
// 		// 服务端md5处理
// 		if utils.Jiami(pwd) == auth.Password {
// 			// 前端md5处理
// 			// if pwd == auth.Password {
// 			return true, nil
// 		}
// 		return false, errors.New("user or pwd is error")
// 	}
// 	return false, errors.New("user or pwd is not right")
// }

// // 查询用户
// func GetAuth() ([]Auth, error) {
// 	data := make([]Auth, 0)
// 	err := utils.Engine.Desc("id").Find(&data)
// 	return data, err
// }

// // 新增用户
// func AddAuth(data *Auth) (int64, error) {
// 	if data.Username == "" || data.Password == "" {
// 		return -1, errors.New("user or pwd is none")
// 	}
// 	data.Password = utils.Jiami(data.Password)
// 	affected, err := utils.Engine.Insert(data)
// 	return affected, err
// }

// // 删除用户
// func DeleteAuth(id string) (int64, error) {
// 	auth := new(Auth)
// 	affected, err := utils.Engine.ID(id).Delete(auth)
// 	return affected, err
// }
