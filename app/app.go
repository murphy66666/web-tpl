package app

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"web-tpl/app/core/config"
	"web-tpl/app/core/db"
	"web-tpl/app/core/log"
	rds "web-tpl/app/core/redis"
)

// Config 所有配置挂载到app.Config
var Config config.Model

// Init 找到2个配置文件路径
// 根据配置文件路径加载内容
// 考虑重载机制
// 再考虑环境变量重载
func Init(prjHome string) error {
	//配置文件组件 viper了解
	return Config.LoadConfig(prjHome)

	//viper配置管理
}

// DBW 动态参数 设置默认值
func DBW(keys ...string) *gorm.DB {
	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}
	conf, ok := Config.DB[k]
	if !ok {
		panic(fmt.Sprintf("db config %s not found", k))
	}

	cacheKey := fmt.Sprintf("%s_write", k)

	return db.Load(conf.Write, conf.Log, cacheKey, Config.Env, Config.HomeDir)
}

func DBR(keys ...string) *gorm.DB {
	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}
	conf, ok := Config.DB[k]
	if !ok {
		panic(fmt.Sprintf("db config  %s not found", k))
	}

	cacheKey := fmt.Sprintf("%s_read", k)

	return db.Load(conf.Read, conf.Log, cacheKey, Config.Env, Config.HomeDir)
}

// Log logrus日志初始化
func Log() *logrus.Entry {

	return log.Load(Config.HomeDir, Config.Log, Config.Env)
}

// ReidsR redis初始化 懒加载
func ReidsR(key ...string) *redis.Client {
	return rds.Load("read", &Config, key...)
}
func ReidsW(key ...string) *redis.Client {
	return rds.Load("write", &Config, key...)
}
