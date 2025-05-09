package redis_service

import (
	"context"
	"github.com/lunarise-dev/lunar-gate/global"
)

func TokenBlack(tokenString string) bool {
	_, err := global.Redis.Get(context.Background(), "black_"+tokenString).Result()
	if err != nil {
		return false
	}
	return true
}
