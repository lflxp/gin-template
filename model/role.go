package model

import (
	"github.com/lflxp/gin-template/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("初始化PageRight")
	utils.Engine.Sync2(new(Nav))
}

// Action
type Action struct {
	Action       string `json:"action" xorm:"varchar(255)"`
	DefaultCheck bool   `json:"defaultCheck" xorm:"bool"`
	Describe     string `json:"describe" xorm:"varchar(255)"`
}

// 用户页面列表
type Page struct {
	RoleId          int64    `json:"roleId"`
	PermissionId    string   `json:"permissionId"`
	PermissionName  string   `json:"permissionName"`
	Actions         []Action `json:"actions"`
	ActionEntitySet []Action `json:"actionEntitySet"`
}

// 用户角色管理
type Role struct {
	Id          string `json:"id" xorm:"varchar(255) notnull index"`
	Name        string `json:"name" xorm:"varchar(255)"`
	Describe    string `json:"describe" xorm:"varchar(255)"`
	Status      bool   `json:"name" xorm:"bool"`
	CreateTime  string `json:"createTime" xorm:"varchar(255)"`
	CreatorId   string `json:"creatorId" xorm:"varchar(255)"`
	Deleted     bool   `json:"deleted" xorm:"bool"`
	Permissions []Page `json:"permissions"`
}

// 动态路由
type Nav struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" xorm:"varchar(255) unique"`
	ParentId  int64  `json:"parentId"`
	Component string `json:"component" xorm:"varchar(255)"`
	Redirect  string `json:"redirect" xorm:"varchar(255)"`
	Path      string `json:"path" xorm:"varchar(255)"`
	Icon      string `json:"icon" xorm:"varchar(255)"`
	Title     string `json:"title" xorm:"varchar(255)"`
	Show      bool   `json:"show" xorm:"bool"`
	// Target    string `json:"target" xorm:"varchar(255)"`
}

// TODO: 判断etcd注册中心是否包含nav标签 没有则过滤
// 查询动态路由
func GetNav() ([]Nav, error) {
	data := make([]Nav, 0)
	err := utils.Engine.Asc("id").Find(&data)
	return data, err
}

// 新增动态路由
func AddNav(data *Nav) (int64, error) {
	affected, err := utils.Engine.Insert(data)
	return affected, err
}

// 删除动态路由
func DeleteNav(id string) (int64, error) {
	c := new(Nav)
	affected, err := utils.Engine.ID(id).Delete(c)
	return affected, err
}

// 修改动态路由
func UpdateNav(id string, data Nav) (int64, error) {
	affected, err := utils.Engine.Table(new(Nav)).ID(id).Update(data)
	return affected, err
}

// 查询前端格式的Nav
func GetNavToFront(username string) ([]interface{}, error) {
	var info []interface{}
	data, err := GetNav()
	if err != nil {
		return info, err
	}

	navdata, err := GetClaimsByAuthAndType(username, "nav")
	if err != nil {
		return info, err
	}

	info = []interface{}{}
	for _, x := range data {
		for _, y := range navdata {
			if x.Name == y.Value {
				tmp := map[string]interface{}{
					"name":     x.Name,
					"parentId": x.ParentId,
					"id":       x.Id,
				}

				if x.Component != "" {
					tmp["component"] = x.Component
				}
				if x.Path != "" {
					tmp["path"] = x.Path
				}
				if x.Redirect != "" {
					tmp["redirect"] = x.Redirect
				}

				meta := map[string]interface{}{
					"title": x.Title,
				}

				if x.Icon != "" {
					meta["icon"] = x.Icon
				}
				if x.Show == false {
					meta["show"] = false
				} else {
					meta["show"] = true
				}
				// if x.Target != "" {
				// 	meta["target"] = x.Target
				// }
				tmp["meta"] = meta

				info = append(info, tmp)
			}
		}
	}
	return info, nil
}
