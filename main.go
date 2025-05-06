package main

import (
	"github.com/lunarise-dev/lunar-gate/core"
	"github.com/lunarise-dev/lunar-gate/global"
)

func main() {
	core.InitLogger()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()
}
