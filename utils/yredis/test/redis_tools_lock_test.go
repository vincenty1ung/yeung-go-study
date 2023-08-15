package test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/uncleyeung/yeung-go-study/utils/ylog"
	"github.com/uncleyeung/yeung-go-study/utils/yredis"
	"go.uber.org/zap"
)

func TestLock(t *testing.T) {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	// uuid.Gen{}
	background := context.Background()

	go func() {
		v1, _ := uuid.NewV1()
		value1 := context.WithValue(background, yredis.ContextKey, "go1 id:"+v1.String())
		redisLock1 := yredis.NewRedisLock(value1, redisX, "test:lock")

		defer func() {
			err := redisLock1.Unlock()
			if err != nil {
				ylog.Error("", zap.Error(err))
			}
			wg.Done()
		}()

		err := redisLock1.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go1 dosomething")
		<-time.After(time.Millisecond * 2900)
		fmt.Println("go1 dosomething sussce")

	}()
	go func() {
		v2, _ := uuid.NewV1()
		value2 := context.WithValue(background, yredis.ContextKey, "go2 id:"+v2.String())
		redisLock2 := yredis.NewRedisLock(value2, redisX, "test:lock")

		defer wg.Done()
		err := redisLock2.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go2 dosomething")
		<-time.After(time.Millisecond * 100)
		fmt.Println("go2 dosomething sussce")
		err = redisLock2.Unlock()
		if err != nil {
			ylog.Error("", zap.Error(err))
		}
	}()

	go func() {
		v3, _ := uuid.NewV1()
		value3 := context.WithValue(background, yredis.ContextKey, "go3 id:"+v3.String())
		redisLock3 := yredis.NewRedisLock(value3, redisX, "test:lock")

		defer wg.Done()
		err := redisLock3.Lock()
		if err != nil {
			ylog.Error("", zap.Error(err))
			return
		}
		fmt.Println("go3 dosomething")
		<-time.After(time.Millisecond * 30)
		fmt.Println("go3 dosomething sussce")
		err = redisLock3.Unlock()
		if err != nil {
			ylog.Error("", zap.Error(err))
		}
	}()

	wg.Wait()
	fmt.Print("总计执行:")
	fmt.Println(time.Since(now))
}

func TestUUID(t *testing.T) {
	// u1 := uuid.Must(uuid.NewV4())

	// Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	// Parse a UUID from a string.
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u3, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	log.Printf("successfully parsed UUID %v", u3)
}
