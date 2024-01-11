package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
	"web-tpl/app/core/config"
)

const (
	dbTimeout      = 5000 //ms
	dbWriteTimeout = 5000
	dbReadTimeout  = 5000
)

// mysql持久连接
var dbInstance = make(map[string]*gorm.DB) //数据库实例 map在并发有问题
var dbLocker sync.RWMutex                  //读写锁

func Load(conf config.DBItemConf, confLog config.DBLog, key string, env string, homeDir string) *gorm.DB {
	dbLocker.RLock()
	db, ok := dbInstance[key]
	if ok {
		dbLocker.RUnlock()
		return db
	}
	dbLocker.RUnlock()

	dbLocker.Lock()
	defer dbLocker.Unlock()

	//todo 优化并发 原子操作 atomic包

	//double check
	if _, exist := dbInstance[key]; exist {
		return dbInstance[key]
	}

	dbInstance[key] = getDbInstance(conf, confLog, env, homeDir)

	return dbInstance[key]
}

// 获取数据库实例
func getDbInstance(conf config.DBItemConf, confLog config.DBLog, env string, homeDir string) *gorm.DB {
	timeout := dbTimeout
	if conf.TimeOut > 0 {
		timeout = conf.TimeOut
	}
	writeTimeout := dbWriteTimeout
	if conf.WriteTimeOut > 0 {
		writeTimeout = conf.WriteTimeOut
	}
	readTimeout := dbReadTimeout
	if conf.ReadTimeOut > 0 {
		readTimeout = conf.ReadTimeOut
	}

	//连接数据库
	dsnConf := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%dms&writeTimeout=%dms&readTimeout=%dms",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset, timeout, writeTimeout, readTimeout)

	//记录sql的log
	var l logger.LogLevel
	var dbLogger *log.Logger
	var gLogger logger.Interface

	if confLog.Enable {
		switch confLog.Level {
		case "silent":
			l = logger.Silent
		case "error":
			l = logger.Error
		case "info":
			l = logger.Info
		default:
			l = logger.Warn
		}

		if confLog.Type == "file" {
			logPath := confLog.Path
			if !filepath.IsAbs(confLog.Path) { //是否是相对路径
				logPath = homeDir + "/" + confLog.Path
			}

			f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0644)

			if err != nil {
				panic(err)
			}
			dbLogger = log.New(f, "\r\n", log.LstdFlags)

		} else if confLog.Type == "stdout" {
			dbLogger = log.New(os.Stdout, "\r\n", log.LstdFlags)
		}

		gLogger = New(dbLogger, logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  l,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}, env)
	}

	db, err := gorm.Open(mysql.Open(dsnConf), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),  //默认的sql日志
		Logger: gLogger,
	})
	if err != nil {
		panic(err)
	}

	//数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	return db
}
