package main

import (
	"flag"
	"log"
	"path/filepath"
	"runtime"
	"web-tpl/app"
	"web-tpl/app/http"
)

// 第一步加载配置
// 第二步启动http服务
func main() {
	// 解析项目根目录参数
	homePath := flag.String("prjHome", "", "项目的根目录路径")
	flag.Parse()
	if *homePath == "" {
		_, f, _, ok := runtime.Caller(0)
		if !ok {
			panic("尝试获取文件路径失败！")
		}
		*homePath = filepath.Dir(f) //解析runtime路径  e.g  E:\cz\GOPATH\web-tpl
	}

	//初始化项目
	err := app.Init(*homePath)
	if err != nil {
		panic(err)
	}

	//启动 server
	err = http.NewServer()
	if err != nil {
		log.Fatal(err)
	}
}
