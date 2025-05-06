package core

import (
	"github.com/lunarise-dev/lunar-gate/config"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() *config.Config {
	byteData, err := os.ReadFile("settings.yaml")
	if err != nil {
		panic("配置文件读取失败 " + err.Error())
		return nil
	}
	var c config.Config
	if err := yaml.Unmarshal(byteData, &c); err != nil {
		panic("配置文件格式错误 " + err.Error())
		return nil
	}
	return &c
}
