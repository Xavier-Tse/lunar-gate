package core

import (
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/Xavier-Tse/lunar-gate/flags"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func SetConfig(c *config.Config) {
	byteData, _ := yaml.Marshal(c)
	err := os.WriteFile(flags.FlagOptions.File, byteData, 0666)
	if err != nil {
		logrus.Warnf("%v", err)
	}
}
