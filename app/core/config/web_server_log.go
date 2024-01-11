package config

type WebServerLog struct {
	Enable          bool     `json:"enable"`          //是否开启日志
	LogIDShowHeader bool     `yaml:"logIDShowHeader"` //日志头
	LogPath         string   `yaml:"logPath"`         //日志路径
	LogFormat       string   `yaml:"logFormat"`       //日志格式
	SkipPaths       []string `yaml:"skipPaths"`       //忽略路径
	Output          string   `yaml:"output"`          //输出格式
}
