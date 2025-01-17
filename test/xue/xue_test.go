package xue

import (
	"context"
	"fmt"
	"testing"

	redis2 "github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/vincenty1ung/yeung-go-study/utils/ylog"
)

func TestName(t *testing.T) {
	t2 := string2T()
	fmt.Println(t2)
	ctx := context.Background()
	key := "TestName"
	for i := range t2 {
		if t2[i].Source.Friends != nil {
			for j := range t2[i].Source.Friends {
				_ = redis2.Z{Score: 1, Member: t2[i].Source.Friends[j]}
				redis.ZIncrBy(ctx, key, 1, t2[i].Source.Friends[j])
			}
		}
	}

	result, err := redis.ZRevRange(ctx, key, 0, 29).Result()
	if err != nil {
		ylog.Error("chuan", zap.Error(err))
	}
	fmt.Println(result)

	// redis.Del(ctx, key)
}
