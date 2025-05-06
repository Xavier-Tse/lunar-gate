package core

import (
	"fmt"
	"github.com/lunarise-dev/lunar-gate/config"
	"gopkg.in/yaml.v3"
	"os"
)

func SetConfig(c *config.Config) {
	byteData, _ := yaml.Marshal(c)
	err := os.WriteFile("settings.yaml", byteData, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
}
