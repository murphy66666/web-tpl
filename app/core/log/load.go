package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"sync"
	"web-tpl/app"
	"web-tpl/app/core/config"
	"web-tpl/app/utils/env"
)

var log *logrus.Entry
var logTimeTpl = "2006-01-02 15:04:05.000Z07:00"
var logLocker sync.RWMutex //配置实例化

// Load todo  未完成加log_id 日志存储不同文件 日志文件按照日期分割
func Load(homeDir string, logConf config.Log, currentenv string) *logrus.Entry {
	logLocker.RLocker()
	if log != nil {
		logLocker.RUnlock()
		return log
	}
	logLocker.RUnlock()

	// a b c
	logLocker.Lock()
	defer logLocker.Unlock()

	//二次判断
	if log != nil {
		return log
	}

	logNew := logrus.New()

	//栈日志 输出错误的文件行号
	logNew.SetReportCaller(true)

	//设置日志level
	setLogLever(logNew, logConf.Level)

	//设置日志格式 json text
	setLogFormat(logNew, logConf.LogFormat, currentenv)

	//设置日志输出
	setLogOutput(logNew, logConf, homeDir)

	//基础字段预设 比如项目名 环境 env local_ip hostname idc
	l := presetFields(logNew, currentenv)
	log = l

	return l
}

func loadLogFile(conf config.Log, homeDir string) (io.Writer, error) {
	logPath := "logs/app.log"
	if conf.Name != "" {
		logPath = conf.Name
	}

	//判断是相对路径还是绝对路径
	if !filepath.IsAbs(logPath) {
		logPath = app.Config.HomeDir + "/" + logPath
	}

	//检测这个文件是否存在，如果不存在就新建文件
	//_, err := os.Stat(logPath)
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		return nil, err
	}

	return f, err
}

func setLogLever(logger *logrus.Logger, level string) {
	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
}

func setLogFormat(logger *logrus.Logger, format, currentenv string) {
	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: logTimeTpl,
		})
	} else {
		//如果非dev环境禁用color 节约性能
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: logTimeTpl,
			DisableColors:   currentenv != "dev",
		})
	}

}

func setLogOutput(logger *logrus.Logger, logConf config.Log, homeDir string) {
	if logConf.Output == "file" {
		f, e := loadLogFile(logConf, homeDir)
		if e != nil {
			panic(e)
		}
		logger.SetOutput(f)
	} else {
		logger.SetOutput(os.Stdout)
	}
}

func presetFields(logger *logrus.Logger, currentenv string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"env":      currentenv,
		"local_ip": env.LocalIP(),
		"hostname": env.Hostname(),
	})
}
