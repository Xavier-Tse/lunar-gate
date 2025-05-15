package core

import (
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/Xavier-Tse/lunar-gate/flags"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() *config.Config {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		logrus.Fatalf("配置文件读取失败 %v", err)
	}
	var c config.Config
	if err := yaml.Unmarshal(byteData, &c); err != nil {
		logrus.Fatalf("配置文件格式错误 %v", err)
	}
	logrus.Infof("配置文件读取成功 %s", flags.FlagOptions.File)
	return &c
}
