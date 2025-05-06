package core

import (
	"github.com/lunarise-dev/lunar-gate/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() *config.Config {
	byteData, err := os.ReadFile("settings.yaml")
	if err != nil {
		logrus.Fatalf("配置文件读取失败 %v", err)
	}
	var c config.Config
	if err := yaml.Unmarshal(byteData, &c); err != nil {
		logrus.Fatalf("配置文件格式错误 %v", err)
	}
	logrus.Infof("配置文件读取成功")
	return &c
}
