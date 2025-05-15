package core

import (
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func SetConfig(c *config.Config) {
	byteData, _ := yaml.Marshal(c)
	err := os.WriteFile("settings.yaml", byteData, 0666)
	if err != nil {
		logrus.Warnf("%v", err)
	}
}
