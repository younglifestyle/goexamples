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

	//  任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	master.InitLogMgr()

	select {}

ERR:
	fmt.Println(err)
}
