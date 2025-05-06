package main

import (
	"fmt"
	"github.com/lunarise-dev/lunar-gate/core"
	"github.com/lunarise-dev/lunar-gate/global"
)

func main() {
	global.Config = core.ReadConfig()
	fmt.Println(global.Config)
}
