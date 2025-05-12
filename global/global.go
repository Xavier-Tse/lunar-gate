package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config     *config.Config
	DB         *gorm.DB
	Redis      *redis.Client
	Casbin     *casbin.CachedEnforcer
	RoutesInfo gin.RoutesInfo
)
