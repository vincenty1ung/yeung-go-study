package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/vincenty1ung/yeung-go-study/utils/yredis"
)

func TestTest(t *testing.T) {
	key := "test"
	ctx := context.Background()
	defer func() {
		result, err := redisX.Del(context.Background(), key).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}()

	go hand(ctx, key, "test1")
	go hand(ctx, key, "test2")
	go hand(ctx, key, "test3")

	time.Sleep(time.Second * 10)

}

func hand(ctx context.Context, key string, s string) bool {
	fmt.Println(s)
	b, _ := yredis.SetNx(ctx, redisX, key, "1")
	if !b {
		fmt.Println(s + "：" + "得不到锁")
		return true
	}
	fmt.Println(s + "：" + "得到锁")
	defer func() {
		fmt.Println(s + "：" + "会执行吗")
		redisX.Del(ctx, key)
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(s + "：" + "做些事情")
	return false
}
