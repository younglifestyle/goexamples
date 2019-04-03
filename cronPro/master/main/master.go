package main

import (
	"flag"
	"fmt"
	"goexamples/cronPro/master"
	"runtime"
)

var (
	confFile string // 配置文件路径
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func main() {

	var (
		err error
	)

	initArgs()

	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 初始化服务发现模块
	if err = master.InitWorkerMgr(); err != nil {
		goto ERR
	}

	// 日志管理器
	master.InitLogMgr()

	//  任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// 启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	select {}

ERR:
	fmt.Println(err)
}
