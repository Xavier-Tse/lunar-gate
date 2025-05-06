package global

import (
	"github.com/lunarise-dev/lunar-gate/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
