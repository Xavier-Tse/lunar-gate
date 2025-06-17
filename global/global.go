package global

import (
	"github.com/Xavier-Tse/lunar-gate/config"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
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
