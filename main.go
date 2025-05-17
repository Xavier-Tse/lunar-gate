package main

import (
	"github.com/Xavier-Tse/lunar-gate/core"
	"github.com/Xavier-Tse/lunar-gate/flags"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/routers"
)

func main() {
	core.InitLogger()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	core.InitIPDB()
	global.Casbin = core.InitCasbin()
	global.Redis = core.InitRedis()

	// 绑定命令行参数
	flags.Run()

	// 启动web服务
	routers.Run()
}
