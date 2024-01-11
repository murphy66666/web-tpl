package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Model 解析yaml文件
type Model struct {
	Env          string            `yaml:"env" env:"GO_ENV"` //环境变量
	HomeDir      string            `yaml:"-"`
	Name         string            `yaml:"name"`
	HTTP         http              `yaml:"http"`  //业务模块
	Log          Log               `yaml:"log"`   //日志模块
	DB           map[string]dbItem `yaml:"db"`    //数据库模块
	WebServerLog WebServerLog      `yaml:"log"`   //自定义日志模块
	Redis        map[string]Redis  `yaml:"redis"` //redis
}

func (m *Model) LoadConfig(prjHome string) error {
	//找到两个配置文件的路径
	confDef := prjHome + "/config/config-default.yml" //默认配置
	confOverride := prjHome + "/config/config.yml"    //重载配置

	//根据配置文件路径加载内容
	err := m.parseConfig(confDef)
	if err != nil {
		return err
	}

	//考虑重载机制 重载覆盖之前的数据
	err = m.parseConfig(confOverride)
	if err != nil {
		return err
	}

	//考虑环境变量重载
	m.mergeEnv()

	m.HomeDir = strings.TrimRight(prjHome, "/")

	return nil
}

// 解析yaml格式的config文件
func (m *Model) parseConfig(conf string) error {
	data, err := os.ReadFile(conf)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

// 环境变量的重载
func (m *Model) mergeEnv() {
	assign(reflect.ValueOf(m))
}

// 变量反射之后的结构体的值
func assign(val reflect.Value) {
	v := reflect.Indirect(val) //判断是否是指针
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("env")
		processOne(v.Field(i), key)
	}
}

// 处理反射断言之后的值
func processOne(field reflect.Value, key string) {
	envVal, envOk := os.LookupEnv(key) //从环境变量获取值

	//获取反射之后的字段种类 todo 后面可以用泛型优化
	switch field.Type().Kind() {
	case reflect.String:
		if !envOk {
			return
		}
		field.SetString(envVal)

	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		if !envOk {
			return
		}
		val, e := strconv.ParseInt(envVal, 0, field.Type().Bits()) //转整型
		if e != nil {
			return
		}
		field.SetInt(val)
	case reflect.Float32, reflect.Float64:
		if !envOk {
			return
		}
		val, e := strconv.ParseFloat(envVal, field.Type().Bits()) //转类型
		if e != nil {
			return
		}
		field.SetFloat(val)
	case reflect.Bool:
		if !envOk {
			return
		}
		val, e := strconv.ParseBool(envVal) //转类型
		if e != nil {
			return
		}
		field.SetBool(val)
	case reflect.Struct: //结构体 递归处理
		assign(field)
	case reflect.Slice: //todo 逗号分隔处理
		//存到数组之中 再遍历
	}
}
