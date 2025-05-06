package core

import (
	"context"
	"fmt"
	"github.com/lunarise-dev/lunar-gate/global"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	r := global.Config.Redis
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})

	_, err := rds.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("redis连接失败 %s", err)
	}
	logrus.Infof("redis连接成功")
	return rds
}
