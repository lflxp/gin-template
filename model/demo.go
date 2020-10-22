package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lflxp/gin-template/utils"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

func init() {
	log.Info("初始化表 Demo")
	utils.Engine.Sync2(new(Demo))
}

// 三级缓存查询
type Demo struct {
	Id         int64  `json:"id"`
	Uuid       string `json:"uuid" xorm:"varchar(255) notnull index unique"`
	Country    string `json:"country" xorm:"varchar(255) not null"`
	Zoom       string `json:"zoom" xorm:"varchar(255) not null"`
	Company    string `json:"company" xorm:"varchar(255) not null"`
	Items      string `json:"items" xorm:"varchar(255) not null"`
	Production string `json:"production" xorm:"varchar(255) not null"`
	Count      string `json:"count" xorm:"varchar(255) not null"`
	Serial     string `json:"serial" xorm:"varchar(255) not null"`
	Extend     string `json:"extend" xorm:"varchar(255) not null"`
}

func (d *Demo) GetByString(key string) (string, error) {
	var t string // 执行结束流程
	// 本地缓存获取
	t = "cache"
	if c1, ok := utils.NewCacheCli().Get(key); !ok {
		t = "redis"
		// redis缓存获取
		if c2, err := utils.NewRedisCli().Get(ctx, key).Result(); err != nil {
			t = "db"
			// sqlite3 获取
			data, has, err := getDemoByUUID(key)
			if err != nil {
				return t, err
			}

			// 不存在
			if !has {
				return t, errors.New(fmt.Sprintf("uuid %s not exists", key))
			}

			d = data

			// 放入缓存
			value, err := json.Marshal(d)
			if err != nil {
				return t, err
			}

			// local cache
			utils.NewCacheCli().Set(key, string(value))
			// redis cluster
			err = utils.NewRedisCli().Set(ctx, key, string(value), 0).Err()
			if err != nil {
				return t, err
			}
		} else {
			// 放入本地缓存
			utils.NewCacheCli().Set(key, c2)
			err = json.Unmarshal([]byte(c2), d)
			if err != nil {
				return t, err
			}
		}
	} else {
		err := json.Unmarshal([]byte(c1.(string)), d)
		if err != nil {
			return t, err
		}
	}
	return t, nil
}

func getDemoByUUID(uuid string) (*Demo, bool, error) {
	data := new(Demo)
	has, err := utils.Engine.Where("uuid = ?", uuid).Get(data)
	return data, has, err
}

func AddDemo(data *Demo) (int64, error) {
	affected, err := utils.Engine.Insert(data)
	return affected, err
}

func DelDemo(id string) (int64, error) {
	data := new(Demo)
	affected, err := utils.Engine.ID(id).Delete(data)
	return affected, err
}

func UpdateDemo(id string, data *Demo) (int64, error) {
	affected, err := utils.Engine.Table(new(Demo)).ID(id).Update(data)
	return affected, err
}
