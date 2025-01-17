package yredis

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/vincenty1ung/yeung-go-study/utils/ylog"
)

const (
	LockKeyPre = "redis_lock:"
	LockExpire = 2000 // 2000毫秒
	ErrIsExist = "get lock fail"

	LockInterval = 1 << iota // 毫秒
	LockRepeated             // 重复次数
	LockError    = iota      // 错误
)

// RedisLock 单点redis，多点注意
type RedisLock struct {
	lockKey   string
	lockValue string
	Key       string
	Field     string
	Expire    int64 // 毫秒
	ctx       context.Context

	rd Redis
}

// 保证原子性（redis是单线程），避免del删除了，其他client获得的lock
var delScript = redis.NewScript(script)
var script = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`
var ContextKey = "Redis_lock"

// NewRedisLock NewRedisLock
func NewRedisLock(ctx context.Context, rd Redis, key string, lockExpire ...int64) *RedisLock {
	lock := new(RedisLock)
	lock.lockKey = LockKeyPre + key
	lock.Key = key
	lock.ctx = ctx

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		ylog.Error("rand.Read fail:", zap.Error(err))
	}

	lock.lockValue = base64.StdEncoding.EncodeToString(b)

	lock.Expire = LockExpire
	if len(lockExpire) > 0 {
		if lockExpire[0] > 0 {
			lock.Expire = lockExpire[0]
		}
	}
	lock.rd = rd
	sync.OnceFunc(
		func() {
			_ = lock.rd.ScriptLoad(ctx, script)
		},
	)()

	return lock
}
func (l *RedisLock) SetContext(ctx context.Context) RedisLock {
	l.ctx = ctx
	return *l
}

// LockNoWait 加锁 获取锁失败马上返回
func (l *RedisLock) LockNoWait() error {
	lockReply, err := l.rd.SetNX(l.ctx, l.lockKey, l.lockValue, time.Duration(l.Expire)*time.Millisecond).Result()
	if err != nil {
		return errors.New("redis fail")
	}

	if !lockReply {
		return errors.New(ErrIsExist)
	}

	return nil
}

// Lock 等待锁 等待50ms
func (l *RedisLock) Lock() error {
	value := l.ctx.Value(ContextKey).(string)
	var b = false
	now := time.Now()
	for i := 0; i < LockRepeated; i++ {
		fmt.Println(value + "....尝试获取")
		err := l.LockNoWait()
		if err != nil {
			if err.Error() != ErrIsExist {
				return err
			}
			time.Sleep(LockInterval * time.Millisecond) // 重试间隔 7毫秒
		} else {
			b = true
			break
		}
	}

	if !b {
		fmt.Print(value + " 获取锁共计:")
		fmt.Println(time.Since(now))
		return errors.New(ErrIsExist)
	}
	return nil
}

// Unlock 解锁
func (l *RedisLock) Unlock() error {
	value := l.ctx.Value(ContextKey).(string)
	fmt.Println(value + "....尝试解锁")
	return delScript.Run(l.ctx, l.rd, []string{l.lockKey}, l.lockValue).Err()
}

// GetString GetString
func (l *RedisLock) GetString() ([]byte, error) {
	if err := l.Lock(); err != nil {
		return nil, err
	}
	result, err := l.rd.Get(l.ctx, l.Key).Bytes()
	if err != nil {
		if err == Nil {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// SetString SetString
func (l *RedisLock) SetString(value interface{}, ex ...int64) error {
	t := DefaultTime
	if len(ex) > 0 {
		t = time.Duration(ex[0]) * time.Second
	}
	err := l.rd.Set(l.ctx, l.Key, value, t).Err()
	if unErr := l.Unlock(); unErr == nil {
		ylog.Error("unlock fail:", zap.Error(unErr))
	}
	return err
}

// Deprecated: 维护自己一套redis后，将弃用
// /////////////set
func (l *RedisLock) GetHash() ([]byte, error) {
	if err := l.Lock(); err != nil {
		return nil, err
	}
	result, err := l.rd.HGet(l.ctx, l.Key, l.Field).Bytes()
	if err != nil {
		if err == Nil {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// Deprecated: 维护自己一套redis后，将弃用
func (l *RedisLock) SetHash(value interface{}, ex ...int64) error {
	err := l.rd.HSet(l.ctx, l.Key, l.Field, value).Err()
	if unErr := l.Unlock(); unErr == nil {
		ylog.Error("unlock fail:", zap.Error(unErr))
	}
	return err
}
