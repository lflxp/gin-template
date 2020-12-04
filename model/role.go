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
