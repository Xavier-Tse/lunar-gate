package global

import (
	"github.com/lunarise-dev/lunar-gate/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
