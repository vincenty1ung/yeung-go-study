package yredis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/uncleyeung/yeung-go-study/utils/ylog"
)

// Nil reply returned by Redis when key does not exist.
const Nil = redis.Nil

// type RedisBase interface {
//	Context() context.Context
//	Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error
//
//	Conn(ctx context.Context) *redis.Conn
//	Process(ctx context.Context, cmd redis.Cmder) error
//	ProcessContext(ctx context.Context, cmd redis.Cmder) error
//
//	Options() *redis.Options
//	PoolStats() *redis.PoolStats
//	PubSub(ctx context.Context) *redis.PubSub
//	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
//	PSubscribe(ctx context.Context, channels ...string) *redis.PubSub
//	// Pipeline() redis.Pipeliner
//	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
//	TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
//	// TxPipeline(ctx context.Context) redis.Pipeliner
// }

var mu = sync.RWMutex{}

var redisMap = make(map[string]*redis.Client)

func GetRedis(key string, options ...redis.Options) *redis.Client {
	/*mu.RLock()
	cli, ok := redisMap[key]
	if ok {
		mu.RUnlock()
		return cli
	}

	mu.RUnlock()
	ulog2.Info("init redisMap key:" + key)
	mu.Lock()
	defer mu.Unlock()

	// 可能有多个被堵塞，排队进来，其实已经初始化好了，所以需要再判断一次s
	cli, ok = redisMap[key]
	if ok {
		ulog2.Warn("unlock, but it's existed key:" + key)
		return cli
	}*/

	/*configKey := ""
	if key == "" {
		configKey = "redis"
	} else {
		configKey = "redis_" + key
	}*/
	// tmpConf := config.Cfg.GetRedis(configKey)

	cli, err := NewClient("192.168.221.10", "", 2, options...)
	// cli, err := NewClient(tmpConf.Host, tmpConf.Password, tmpConf.Db, options...)
	if err != nil {
		panic(err)
	}
	ylog.Info("set redisMap key:" + key)
	redisMap[key] = cli

	return cli
}

func NewClient(host string, pwd string, db int, options ...redis.Options) (*redis.Client, error) {
	curRedisPoolSize := 100
	if len(options) > 0 && options[0].PoolSize > 0 {
		curRedisPoolSize = options[0].PoolSize
	}
	configs := strings.Split(host, ",")
	con := redis.NewClient(
		&redis.Options{
			PoolSize: curRedisPoolSize,
			Addr:     strings.TrimSpace(configs[0]),
			Password: pwd,
			DB:       db,
		},
	)

	// client := &ClientX{Client: con}
	// con.Ping(context.Background())
	if _, err := con.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	con.AddHook(_rH)
	return con, nil
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	} else {
		return errors.Is(err, redis.Nil)
	}
}

type redisHookTimeKey struct{}

var redisHookTimeK = redisHookTimeKey{}

type redisHook struct {
	redis.Hook
}

var _rH redis.Hook = redisHook{}

func (redisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	fmt.Println("=====redis====")
	withValue := context.WithValue(ctx, redisHookTimeK, time.Now())
	fmt.Println("你好世界,BeforeProcess")
	fmt.Println(cmd.String())
	return withValue, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	now := time.Now()
	defer AssertError(cmd.Err())
	times, ok := ctx.Value(redisHookTimeK).(time.Time)
	if !ok {
		times = time.Now()
	}
	// ulog2.LogRedis(ctx, times, cmd.Name(), cmd.Args())
	fmt.Println(now.Sub(times))
	// ylog.Info("处理时间", zap.Any("时长", now.Sub(times)))
	fmt.Println("你好世界,AfterProcess")
	fmt.Println("====redis=====")
	return nil
}

func AssertError(err error) {
	if err != nil {
		panic(err)
	}
}
