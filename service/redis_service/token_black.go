package redis_service

import (
	"context"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/sirupsen/logrus"
	"time"
)

func TokenBlack(tokenString string) bool {
	_, err := global.Redis.Get(context.Background(), "black_"+tokenString).Result()
	if err != nil {
		return false
	}
	return true
}

func AddToken(token string, expiration time.Duration) {
	key := "black_" + token
	err := global.Redis.Set(context.Background(), key, "", expiration).Err()
	if err != nil {
		logrus.Errorf("token写入redis失败 %v", err)
	}
}
