package yredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// fucnName (?<=\*redisN\)\s).[a-zA-z0-9_]*

type redisN struct {
	*redis.Client
}

func NewRedisN(conn *redis.Client) Redis {
	client := &redisN{Client: conn}
	return client
}

func GetRedisN(key string, options ...redis.Options) Redis {
	conn := GetRedis(key, options...)
	return NewRedisN(conn)
}

/*// context context
func (c *redisN) context() context.Context {
	return c.Client.Context()
}*/

// Ping Ping
func (c *redisN) Ping(ctx context.Context) *redis.StatusCmd {
	return c.Client.Ping(ctx)
}

// Pipelined Pipelined
func (c *redisN) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.Client.Pipelined(ctx, fn)
}

// TxPipelined TxPipelined
func (c *redisN) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.Client.TxPipelined(ctx, fn)
}

// Command Command
func (c *redisN) Command(ctx context.Context) *redis.CommandsInfoCmd {
	return c.Client.Command(ctx)
}

// ClientGetName ClientGetName
func (c *redisN) ClientGetName(ctx context.Context) *redis.StringCmd {
	return c.Client.ClientGetName(ctx)
}

func (c *redisN) _preHandler(ctx context.Context, startTime time.Time, cmd string, args ...interface{}) {
	// xlog.LogRedis(ctx, startTime, cmd, args)
}

// ReadOnly ReadOnly
func (c *redisN) ReadOnly(ctx context.Context) *redis.StatusCmd {
	return c.Client.ReadOnly(ctx)
}

// ReadWrite ReadWrite
func (c *redisN) ReadWrite(ctx context.Context) *redis.StatusCmd {
	return c.Client.ReadWrite(ctx)
}

// MemoryUsage MemoryUsage
func (c *redisN) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	return c.Client.MemoryUsage(ctx, key, samples...)
}

// Watch Watch
func (c *redisN) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	return c.Client.Watch(ctx, fn, keys...)
}

// Clone Clone
func (c *redisN) Clone(ctx context.Context) Redis {
	return &redisN{
		Client: c.Client.WithContext(ctx),
	}
}

// WithTimeout WithTimeout
func (c *redisN) WithTimeout(ctx context.Context, timeout time.Duration) Redis {
	if timeout <= 0 {
		timeout = 0
	}

	return &redisN{
		Client: c.Client.WithTimeout(timeout),
	}
}

// Context Context
func (c *redisN) Context() context.Context {
	return c.Client.Context()
}

// WithContext WithContext
func (c *redisN) WithContext(ctx context.Context) Redis {
	return &redisN{
		Client: c.Client.WithContext(ctx),
	}
}

func (c *redisN) Conn(ctx context.Context) *redis.Conn {
	return c.Client.Conn(ctx)
}

func (c *redisN) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	return c.Client.Do(ctx, args...)
}

// DoContext DoContext
func (c *redisN) DoContext(ctx context.Context, args ...interface{}) *redis.Cmd {
	return c.Client.Do(ctx, args...)
}

// Process Process
func (c *redisN) Process(ctx context.Context, cmd redis.Cmder) error {
	return c.Client.Process(ctx, cmd)
}

// ProcessContext ProcessContext
func (c *redisN) ProcessContext(ctx context.Context, cmd redis.Cmder) error {
	return c.Client.Process(ctx, cmd)
}

// Options Options
func (c *redisN) Options() *redis.Options {
	return c.Client.Options()
}

// PoolStats PoolStats
func (c *redisN) PoolStats() *redis.PoolStats {
	return c.Client.PoolStats()
}

// PubSub PubSub
func (c *redisN) PubSub(ctx context.Context) *redis.PubSub {
	return c.Client.Subscribe(ctx)
}

// Subscribe Subscribe
func (c *redisN) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return c.Client.Subscribe(ctx, channels...)
}

// PSubscribe PSubscribe
func (c *redisN) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return c.Client.PSubscribe(ctx, channels...)
}

// Set Set
func (c *redisN) Set(
	ctx context.Context, key string, value interface{}, expiration time.Duration,
) *redis.StatusCmd {
	if expiration <= 0 {
		expiration = 0
	}
	defer c._preHandler(ctx, time.Now(), "Set", key, value, expiration)
	cmd := c.Client.Set(ctx, key, value, expiration)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Get Get
func (c *redisN) Get(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "Get", key)
	cmd := c.Client.Get(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Echo Echo
func (c *redisN) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "Echo", message)
	cmd := c.Client.Echo(ctx, message)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Quit Quit
func (c *redisN) Quit(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "Quit")
	cmd := c.Client.Quit(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Del Del
func (c *redisN) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Del", keys)
	cmd := c.Client.Del(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Unlink Unlink
func (c *redisN) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Unlink", keys)
	cmd := c.Client.Unlink(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Dump Dump
func (c *redisN) Dump(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "Dump", key)
	cmd := c.Client.Dump(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Exists Exists
func (c *redisN) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Exists", keys)
	cmd := c.Client.Exists(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Expire Expire
func (c *redisN) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if expiration <= 0 {
		expiration = 0
	}
	defer c._preHandler(ctx, time.Now(), "Expire", key, expiration)
	cmd := c.Client.Expire(ctx, key, expiration)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ExpireAt ExpireAt
func (c *redisN) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "ExpireAt", key, tm)
	cmd := c.Client.ExpireAt(ctx, key, tm)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Keys Keys
func (c *redisN) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "Keys", pattern)
	cmd := c.Client.Keys(ctx, pattern)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Migrate Migrate
func (c *redisN) Migrate(
	ctx context.Context, host, port, key string, db int, timeout time.Duration,
) *redis.StatusCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "Migrate", host, port, key, db, timeout)
	cmd := c.Client.Migrate(ctx, host, port, key, db, timeout)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Move Move
func (c *redisN) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "Move", key, db)
	cmd := c.Client.Move(ctx, key, db)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ObjectRefCount ObjectRefCount
func (c *redisN) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ObjectRefCount", key)
	cmd := c.Client.ObjectRefCount(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ObjectEncoding ObjectEncoding
func (c *redisN) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "ObjectEncoding", key)
	cmd := c.Client.ObjectEncoding(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ObjectIdleTime ObjectIdleTime
func (c *redisN) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	defer c._preHandler(ctx, time.Now(), "ObjectIdleTime", key)
	cmd := c.Client.ObjectIdleTime(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Persist Persist
func (c *redisN) Persist(ctx context.Context, key string) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "Persist", key)
	cmd := c.Client.Persist(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PExpire PExpire
func (c *redisN) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if expiration <= 0 {
		expiration = 0
	}

	defer c._preHandler(ctx, time.Now(), "PExpire", key, expiration)
	cmd := c.Client.PExpire(ctx, key, expiration)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PExpireAt PExpireAt
func (c *redisN) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "PExpireAt", key, tm)
	cmd := c.Client.PExpireAt(ctx, key, tm)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PTTL PTTL
func (c *redisN) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	defer c._preHandler(ctx, time.Now(), "PTTL", key)
	cmd := c.Client.PTTL(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RandomKey RandomKey
func (c *redisN) RandomKey(ctx context.Context) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "RandomKey")
	cmd := c.Client.RandomKey(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Rename Rename
func (c *redisN) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "Rename", key, newkey)
	cmd := c.Client.Rename(ctx, key, newkey)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RenameNX RenameNX
func (c *redisN) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "RenameNX", key, newkey)
	cmd := c.Client.RenameNX(ctx, key, newkey)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Restore Restore
func (c *redisN) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	if ttl <= 0 {
		ttl = 0
	}

	defer c._preHandler(ctx, time.Now(), "Restore", key, ttl, value)
	cmd := c.Client.Restore(ctx, key, ttl, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RestoreReplace RestoreReplace
func (c *redisN) RestoreReplace(
	ctx context.Context, key string, ttl time.Duration, value string,
) *redis.StatusCmd {
	if ttl <= 0 {
		ttl = 0
	}
	defer c._preHandler(ctx, time.Now(), "RestoreReplace", key, ttl, value)
	cmd := c.Client.RestoreReplace(ctx, key, ttl, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Sort Sort
func (c *redisN) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "Sort", key, sort)
	cmd := c.Client.Sort(ctx, key, sort)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SortStore SortStore
func (c *redisN) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SortStore", key, store, sort)
	cmd := c.Client.SortStore(ctx, key, store, sort)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SortInterfaces SortInterfaces
func (c *redisN) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	defer c._preHandler(ctx, time.Now(), "SortInterfaces", key, sort)
	cmd := c.Client.SortInterfaces(ctx, key, sort)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Touch Touch
func (c *redisN) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Touch", keys)
	cmd := c.Client.Touch(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// TTL TTL
func (c *redisN) TTL(ctx context.Context, key string) *redis.DurationCmd {
	defer c._preHandler(ctx, time.Now(), "TTL", key)
	cmd := c.Client.TTL(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Type Type
func (c *redisN) Type(ctx context.Context, key string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "Type", key)
	cmd := c.Client.Type(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Scan Scan
func (c *redisN) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	defer c._preHandler(ctx, time.Now(), "Scan", cursor, match, count)
	cmd := c.Client.Scan(ctx, cursor, match, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SScan SScan
func (c *redisN) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	defer c._preHandler(ctx, time.Now(), "SScan", key, cursor, match, count)
	cmd := c.Client.SScan(ctx, key, cursor, match, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HScan HScan
func (c *redisN) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	defer c._preHandler(ctx, time.Now(), "HScan", key, cursor, match, count)
	cmd := c.Client.HScan(ctx, key, cursor, match, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZScan ZScan
func (c *redisN) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	defer c._preHandler(ctx, time.Now(), "ZScan", key, cursor, match, count)
	cmd := c.Client.ZScan(ctx, key, cursor, match, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Append Append
func (c *redisN) Append(ctx context.Context, key, value string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Append", key, value)
	cmd := c.Client.Append(ctx, key, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitCount BitCount
func (c *redisN) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitCount", key, bitCount)
	cmd := c.Client.BitCount(ctx, key, bitCount)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitOpAnd BitOpAnd
func (c *redisN) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitOpAnd", destKey, keys)
	cmd := c.Client.BitOpAnd(ctx, destKey, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitOpOr BitOpOr
func (c *redisN) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitOpOr", destKey, keys)
	cmd := c.Client.BitOpOr(ctx, destKey, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitOpXor BitOpXor
func (c *redisN) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitOpXor", destKey, keys)
	cmd := c.Client.BitOpXor(ctx, destKey, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitOpNot BitOpNot
func (c *redisN) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitOpNot", destKey, key)
	cmd := c.Client.BitOpXor(ctx, destKey, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitPos BitPos
func (c *redisN) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "BitPos", key, bit, pos)
	cmd := c.Client.BitPos(ctx, key, bit, pos...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BitField BitField
func (c *redisN) BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd {
	defer c._preHandler(ctx, time.Now(), "BitField", key, args)
	cmd := c.Client.BitField(ctx, key, args...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Decr Decr
func (c *redisN) Decr(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Decr", key)
	cmd := c.Client.Decr(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// DecrBy DecrBy
func (c *redisN) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "DecrBy", key, decrement)
	cmd := c.Client.DecrBy(ctx, key, decrement)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GetBit GetBit
func (c *redisN) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "GetBit", key, offset)
	cmd := c.Client.GetBit(ctx, key, offset)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GetRange GetRange
func (c *redisN) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "GetRange", key, start, end)
	cmd := c.Client.GetRange(ctx, key, start, end)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GetSet GetSet
func (c *redisN) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "GetSet", key, value)
	cmd := c.Client.GetSet(ctx, key, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Incr Incr
func (c *redisN) Incr(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Incr", key)
	cmd := c.Client.Incr(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// IncrBy IncrBy
func (c *redisN) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "IncrBy", key, value)
	cmd := c.Client.IncrBy(ctx, key, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// IncrByFloat IncrByFloat
func (c *redisN) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "IncrByFloat", key, value)
	cmd := c.Client.IncrByFloat(ctx, key, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// MGet MGet
func (c *redisN) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	defer c._preHandler(ctx, time.Now(), "MGet", keys)
	cmd := c.Client.MGet(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// MSet MSet
func (c *redisN) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "MSet", values)
	cmd := c.Client.MSet(ctx, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// MSetNX MSetNX
func (c *redisN) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "MSetNX", values)
	cmd := c.Client.MSetNX(ctx, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SetBit SetBit
func (c *redisN) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SetBit", key, offset, value)
	cmd := c.Client.SetBit(ctx, key, offset, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SetNX SetNX
func (c *redisN) SetNX(
	ctx context.Context, key string, value interface{}, expiration time.Duration,
) *redis.BoolCmd {
	if expiration <= 0 {
		expiration = 0
	}
	defer c._preHandler(ctx, time.Now(), "SetNX", key, value, expiration)
	cmd := c.Client.SetNX(ctx, key, value, expiration)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SetXX SetXX
func (c *redisN) SetXX(
	ctx context.Context, key string, value interface{}, expiration time.Duration,
) *redis.BoolCmd {
	if expiration <= 0 {
		expiration = 0
	}
	defer c._preHandler(ctx, time.Now(), "SetXX", key, value, expiration)
	cmd := c.Client.SetXX(ctx, key, value, expiration)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SetRange SetRange
func (c *redisN) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SetRange", key, offset, value)
	cmd := c.Client.SetRange(ctx, key, offset, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// StrLen StrLen
func (c *redisN) StrLen(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "StrLen", key)
	cmd := c.Client.StrLen(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HDel HDel
func (c *redisN) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "HDel", key, fields)
	cmd := c.Client.HDel(ctx, key, fields...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HExists HExists
func (c *redisN) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "HExists", key, field)
	cmd := c.Client.HExists(ctx, key, field)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HGet HGet
func (c *redisN) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "HGet", key, field)
	cmd := c.Client.HGet(ctx, key, field)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HGetAll HGetAll
func (c *redisN) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	defer c._preHandler(ctx, time.Now(), "HGetAll", key)
	cmd := c.Client.HGetAll(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HIncrBy HIncrBy
func (c *redisN) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "HIncrBy", key, field, incr)
	cmd := c.Client.HIncrBy(ctx, key, field, incr)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HIncrByFloat HIncrByFloat
func (c *redisN) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "HIncrByFloat", key, field, incr)
	cmd := c.Client.HIncrByFloat(ctx, key, field, incr)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HKeys HKeys
func (c *redisN) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "HKeys", key)
	cmd := c.Client.HKeys(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HLen HLen
func (c *redisN) HLen(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "HLen", key)
	cmd := c.Client.HLen(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HMGet HMGet
func (c *redisN) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	defer c._preHandler(ctx, time.Now(), "HMGet", key, fields)
	cmd := c.Client.HMGet(ctx, key, fields...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HSet HSet
func (c *redisN) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "HSet", key, values)
	cmd := c.Client.HSet(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HMSet HMSet
func (c *redisN) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "HMSet", key, values)
	cmd := c.Client.HMSet(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HSetNX HSetNX
func (c *redisN) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "HSetNX", key, field, value)
	cmd := c.Client.HSetNX(ctx, key, field, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// HVals HVals
func (c *redisN) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "HVals", key)
	cmd := c.Client.HVals(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BLPop BLPop
func (c *redisN) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "BLPop", timeout, keys)
	cmd := c.Client.BLPop(ctx, timeout, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BRPop BRPop
func (c *redisN) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "BRPop", timeout, keys)
	cmd := c.Client.BRPop(ctx, timeout, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BRPopLPush BRPopLPush
func (c *redisN) BRPopLPush(
	ctx context.Context, source, destination string, timeout time.Duration,
) *redis.StringCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "BRPopLPush", source, destination, timeout)
	cmd := c.Client.BRPopLPush(ctx, source, destination, timeout)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LIndex LIndex
func (c *redisN) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "LIndex", key, index)
	cmd := c.Client.LIndex(ctx, key, index)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LInsert LInsert
func (c *redisN) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LInsert", key, op, pivot, value)
	cmd := c.Client.LInsert(ctx, key, op, pivot, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LInsertBefore LInsertBefore
func (c *redisN) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LInsertBefore", key, pivot, value)
	cmd := c.Client.LInsertBefore(ctx, key, pivot, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LInsertAfter LInsertAfter
func (c *redisN) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LInsertAfter", key, pivot, value)
	cmd := c.Client.LInsertAfter(ctx, key, pivot, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LLen LLen
func (c *redisN) LLen(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LLen", key)
	cmd := c.Client.LLen(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LPop LPop
func (c *redisN) LPop(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "LPop", key)
	cmd := c.Client.LPop(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LPush LPush
func (c *redisN) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LPush", key, values)
	cmd := c.Client.LPush(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LPushX LPushX
func (c *redisN) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LPushX", key, values)
	cmd := c.Client.LPushX(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LRange LRange
func (c *redisN) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "LRange", key, start, stop)
	cmd := c.Client.LRange(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LRem LRem
func (c *redisN) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LRem", key, count, value)
	cmd := c.Client.LRem(ctx, key, count, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LSet LSet
func (c *redisN) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "LSet", key, index, value)
	cmd := c.Client.LSet(ctx, key, index, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LTrim LTrim
func (c *redisN) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "LTrim", key, start, stop)
	cmd := c.Client.LTrim(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RPop RPop
func (c *redisN) RPop(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "RPop", key)
	cmd := c.Client.RPop(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RPopLPush RPopLPush
func (c *redisN) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "RPopLPush", source, destination)
	cmd := c.Client.RPopLPush(ctx, source, destination)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RPush RPush
func (c *redisN) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "RPush", key, values)
	cmd := c.Client.RPush(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// RPushX RPushX
func (c *redisN) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "RPushX", key, values)
	cmd := c.Client.RPushX(ctx, key, values...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SAdd SAdd
func (c *redisN) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SAdd", members...)
	cmd := c.Client.SAdd(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SCard SCard
func (c *redisN) SCard(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SCard", key)
	cmd := c.Client.SCard(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SDiff SDiff
func (c *redisN) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SDiff", keys)
	cmd := c.Client.SDiff(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SDiffStore SDiffStore
func (c *redisN) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SDiffStore", destination, keys)
	cmd := c.Client.SDiffStore(ctx, destination, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SInter SInter
func (c *redisN) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SInter", keys)
	cmd := c.Client.SInter(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SInterStore SInterStore
func (c *redisN) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SInterStore", destination, keys)
	cmd := c.Client.SInterStore(ctx, destination, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SIsMember SIsMember
func (c *redisN) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), "SIsMember", key, member)
	cmd := c.Client.SIsMember(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SMembers SMembers
func (c *redisN) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SMembers", key)
	cmd := c.Client.SMembers(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SMembersMap SMembersMap
func (c *redisN) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	defer c._preHandler(ctx, time.Now(), "SMembersMap", key)
	cmd := c.Client.SMembersMap(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SMove SMove
func (c *redisN) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	defer c._preHandler(ctx, time.Now(), source, destination, member)
	cmd := c.Client.SMove(ctx, source, destination, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SPop SPop
func (c *redisN) SPop(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "SPop", key)
	cmd := c.Client.SPop(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SPopN SPopN
func (c *redisN) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SPopN", key, count)
	cmd := c.Client.SPopN(ctx, key, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SRandMember SRandMember
func (c *redisN) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "SRandMember", key)
	cmd := c.Client.SRandMember(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SRandMemberN SRandMemberN
func (c *redisN) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SRandMemberN", key, count)
	cmd := c.Client.SRandMemberN(ctx, key, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SRem SRem
func (c *redisN) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SRem", key, members)
	cmd := c.Client.SRem(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SUnion SUnion
func (c *redisN) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "SUnion", keys)
	cmd := c.Client.SUnion(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SUnionStore SUnionStore
func (c *redisN) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "SUnionStore", destination, keys)
	cmd := c.Client.SUnionStore(ctx, destination, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XAdd XAdd
func (c *redisN) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "XAdd", a)
	cmd := c.Client.XAdd(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XDel XDel
func (c *redisN) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XDel", stream, ids)
	cmd := c.Client.XDel(ctx, stream, ids...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XLen XLen
func (c *redisN) XLen(ctx context.Context, stream string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XLen", stream)
	cmd := c.Client.XLen(ctx, stream)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XRange XRange
func (c *redisN) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XRange", stream, start, stop)
	cmd := c.Client.XRange(ctx, stream, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XRangeN XRangeN
func (c *redisN) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XRangeN", stream, start, stop, count)
	cmd := c.Client.XRangeN(ctx, stream, start, stop, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XRevRange XRevRange
func (c *redisN) XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XRevRange", stream, start, stop)
	cmd := c.Client.XRevRange(ctx, stream, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XRevRangeN XRevRangeN
func (c *redisN) XRevRangeN(
	ctx context.Context, stream string, start, stop string, count int64,
) *redis.XMessageSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XRevRangeN", stream, start, stop, count)
	cmd := c.Client.XRevRangeN(ctx, stream, start, stop, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XRead XRead
func (c *redisN) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XRead", a)
	cmd := c.Client.XRead(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XReadStreams XReadStreams
func (c *redisN) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XReadStreams", streams)
	cmd := c.Client.XReadStreams(ctx, streams...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XGroupCreate XGroupCreate
func (c *redisN) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "XGroupCreate", stream, group, start)
	cmd := c.Client.XGroupCreate(ctx, stream, group, start)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XGroupCreateMkStream XGroupCreateMkStream
func (c *redisN) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "XGroupCreateMkStream", stream, group, start)
	cmd := c.Client.XGroupCreate(ctx, stream, group, start)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XGroupSetID XGroupSetID
func (c *redisN) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "XGroupSetID", stream, group, start)
	cmd := c.Client.XGroupSetID(ctx, stream, group, start)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XGroupDestroy XGroupDestroy
func (c *redisN) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XGroupDestroy", stream, group)
	cmd := c.Client.XGroupDestroy(ctx, stream, group)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XGroupDelConsumer XGroupDelConsumer
func (c *redisN) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XGroupDelConsumer", stream, group, consumer)
	cmd := c.Client.XGroupDelConsumer(ctx, stream, group, consumer)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XReadGroup XReadGroup
func (c *redisN) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XReadGroup", a)
	cmd := c.Client.XReadGroup(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XAck XAck
func (c *redisN) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XAck", stream, group, ids)
	cmd := c.Client.XAck(ctx, stream, group, ids...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XPending XPending
func (c *redisN) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	defer c._preHandler(ctx, time.Now(), "XPending", stream, group)
	cmd := c.Client.XPending(ctx, stream, group)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XPendingExt XPendingExt
func (c *redisN) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	defer c._preHandler(ctx, time.Now(), "XPendingExt", a)
	cmd := c.Client.XPendingExt(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XClaim XClaim
func (c *redisN) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XClaim", a)
	cmd := c.Client.XClaim(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XClaimJustID XClaimJustID
func (c *redisN) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "XClaimJustID", a)
	cmd := c.Client.XClaimJustID(ctx, a)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XTrim XTrim
func (c *redisN) XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XTrim", key, maxLen)
	cmd := c.Client.XTrimMaxLen(ctx, key, maxLen)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XTrimApprox XTrimApprox
func (c *redisN) XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "XTrimApprox", key, maxLen, 0)
	cmd := c.Client.XTrimMaxLenApprox(ctx, key, maxLen, 0)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// XInfoGroups XInfoGroups
func (c *redisN) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	defer c._preHandler(ctx, time.Now(), "XInfoGroups", key)
	cmd := c.Client.XInfoGroups(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BZPopMax BZPopMax
func (c *redisN) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "BZPopMax", timeout, keys)
	cmd := c.Client.BZPopMax(ctx, timeout, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BZPopMin BZPopMin
func (c *redisN) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if timeout <= 0 {
		timeout = 0
	}
	defer c._preHandler(ctx, time.Now(), "BZPopMin", timeout, keys)
	cmd := c.Client.BZPopMin(ctx, timeout, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZAdd ZAdd
func (c *redisN) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAdd", key, members)
	cmd := c.Client.ZAdd(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZAddNX ZAddNX
func (c *redisN) ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAddNX", key, key, members)
	cmd := c.Client.ZAddNX(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZAddXX ZAddXX
func (c *redisN) ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAddXX", key, members)
	cmd := c.Client.ZAddXX(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

//
// func (c *redisN) _afterHandlerCheckPanic(err error) {
//	xglobal.AssertError(err)
// }

// ZAddCh ZAddCh
func (c *redisN) ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAddCh", members)
	cmd := c.Client.ZAddCh(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZAddNXCh ZAddNXCh
func (c *redisN) ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAddNXCh", key, members)
	cmd := c.Client.ZAddNXCh(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZAddXXCh ZAddXXCh
func (c *redisN) ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZAddXXCh", key, members)
	cmd := c.Client.ZAddXXCh(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZIncr ZIncr
func (c *redisN) ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "ZIncr", key, member)
	cmd := c.Client.ZIncr(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZIncrNX ZIncrNX
func (c *redisN) ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "ZIncrNX", key, member)
	cmd := c.Client.ZIncrNX(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZIncrXX ZIncrXX
func (c *redisN) ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "ZIncrXX", key)
	cmd := c.Client.ZIncrXX(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZCard ZCard
func (c *redisN) ZCard(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZCard", key)
	cmd := c.Client.ZCard(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZCount ZCount
func (c *redisN) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZCount", key, min, max)
	cmd := c.Client.ZCount(ctx, key, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZLexCount ZLexCount
func (c *redisN) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZLexCount", key, min, max)
	cmd := c.Client.ZLexCount(ctx, key, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZIncrBy ZIncrBy
func (c *redisN) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "ZIncrBy", key, increment, member)
	cmd := c.Client.ZIncrBy(ctx, key, increment, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZInterStore ZInterStore
func (c *redisN) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZInterStore", destination, store)
	cmd := c.Client.ZInterStore(ctx, destination, store)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZPopMax ZPopMax
func (c *redisN) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZPopMax", key, count)
	cmd := c.Client.ZPopMax(ctx, key, count...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZPopMin ZPopMin
func (c *redisN) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZPopMin", key, count)
	cmd := c.Client.ZPopMin(ctx, key, count...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRange ZRange
func (c *redisN) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRange", key, start, stop)
	cmd := c.Client.ZRange(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRangeWithScores ZRangeWithScores
func (c *redisN) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRangeWithScores", key, start, stop)
	cmd := c.Client.ZRangeWithScores(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRangeByScore ZRangeByScore
func (c *redisN) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRangeByScore", key, opt)
	cmd := c.Client.ZRangeByScore(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRangeByLex ZRangeByLex
func (c *redisN) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRangeByLex", key, opt)
	cmd := c.Client.ZRangeByLex(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRangeByScoreWithScores ZRangeByScoreWithScores
func (c *redisN) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRangeByScoreWithScores", key, opt)
	cmd := c.Client.ZRangeByScoreWithScores(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRank ZRank
func (c *redisN) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRank", key, member)
	cmd := c.Client.ZRank(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRem ZRem
func (c *redisN) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRem", key, members)
	cmd := c.Client.ZRem(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRemRangeByRank ZRemRangeByRank
func (c *redisN) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRemRangeByRank", key, start, stop)
	cmd := c.Client.ZRemRangeByRank(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRemRangeByScore ZRemRangeByScore
func (c *redisN) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRemRangeByScore", key, min, max)
	cmd := c.Client.ZRemRangeByScore(ctx, key, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRemRangeByLex ZRemRangeByLex
func (c *redisN) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRemRangeByLex", key, min, max)
	cmd := c.Client.ZRemRangeByLex(ctx, key, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRange ZRevRange
func (c *redisN) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRange", key, start, stop)
	cmd := c.Client.ZRevRange(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRangeWithScores ZRevRangeWithScores
func (c *redisN) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRangeWithScores", key, start, stop)
	cmd := c.Client.ZRevRangeWithScores(ctx, key, start, stop)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRangeByScore ZRevRangeByScore
func (c *redisN) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRangeByScore", key, opt)
	cmd := c.Client.ZRevRangeByScore(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRangeByLex ZRevRangeByLex
func (c *redisN) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRangeByLex", key, opt)
	cmd := c.Client.ZRevRangeByLex(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRangeByScoreWithScores ZRevRangeByScoreWithScores
func (c *redisN) ZRevRangeByScoreWithScores(
	ctx context.Context, key string, opt *redis.ZRangeBy,
) *redis.ZSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRangeByScoreWithScores", key, opt)
	cmd := c.Client.ZRevRangeByScoreWithScores(ctx, key, opt)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZRevRank ZRevRank
func (c *redisN) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZRevRank", key, member)
	cmd := c.Client.ZRevRank(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZScore ZScore
func (c *redisN) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "ZScore", key, member)
	cmd := c.Client.ZScore(ctx, key, member)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ZUnionStore ZUnionStore
func (c *redisN) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ZUnionStore", dest, store)
	cmd := c.Client.ZUnionStore(ctx, dest, store)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PFAdd PFAdd
func (c *redisN) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "PFAdd", key, els)
	cmd := c.Client.PFAdd(ctx, key, els...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PFCount PFCount
func (c *redisN) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "PFCount")
	cmd := c.Client.PFCount(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PFMerge PFMerge
func (c *redisN) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "PFMerge")
	cmd := c.Client.PFMerge(ctx, dest, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BgRewriteAOF BgRewriteAOF
func (c *redisN) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "BgRewriteAOF")
	cmd := c.Client.BgRewriteAOF(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// BgSave BgSave
func (c *redisN) BgSave(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "BgSave")
	cmd := c.Client.BgSave(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClientKill ClientKill
func (c *redisN) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClientKill")
	cmd := c.Client.ClientKill(ctx, ipPort)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClientKillByFilter ClientKillByFilter
func (c *redisN) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ClientKillByFilter")
	cmd := c.Client.ClientKillByFilter(ctx, keys...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClientList ClientList
func (c *redisN) ClientList(ctx context.Context) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "ClientList")
	cmd := c.Client.ClientList(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClientPause ClientPause
func (c *redisN) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	if dur <= 0 {
		dur = 0
	}
	defer c._preHandler(ctx, time.Now(), "ClientPause")
	cmd := c.Client.ClientPause(ctx, dur)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClientID ClientID
func (c *redisN) ClientID(ctx context.Context) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ClientID")
	cmd := c.Client.ClientID(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ConfigGet ConfigGet
func (c *redisN) ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd {
	defer c._preHandler(ctx, time.Now(), "ConfigGet")
	cmd := c.Client.ConfigGet(ctx, parameter)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ConfigResetStat ConfigResetStat
func (c *redisN) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ConfigResetStat")
	cmd := c.Client.ConfigResetStat(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ConfigSet ConfigSet
func (c *redisN) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ConfigSet")
	cmd := c.Client.ConfigSet(ctx, parameter, value)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ConfigRewrite ConfigRewrite
func (c *redisN) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ConfigRewrite")
	cmd := c.Client.ConfigRewrite(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// DBSize DBSize
func (c *redisN) DBSize(ctx context.Context) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "DBSize")
	cmd := c.Client.DBSize(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// FlushAll FlushAll
func (c *redisN) FlushAll(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "FlushAll")
	cmd := c.Client.FlushAll(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// FlushAllAsync FlushAllAsync
func (c *redisN) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "FlushAllAsync")
	cmd := c.Client.FlushAllAsync(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// FlushDB FlushDB
func (c *redisN) FlushDB(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "FlushDB")
	cmd := c.Client.FlushDB(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// FlushDBAsync FlushDBAsync
func (c *redisN) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "FlushDBAsync")
	cmd := c.Client.FlushDBAsync(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Info Info
func (c *redisN) Info(ctx context.Context, section ...string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "Info")
	cmd := c.Client.Info(ctx, section...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// LastSave LastSave
func (c *redisN) LastSave(ctx context.Context) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "LastSave")
	cmd := c.Client.LastSave(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Save Save
func (c *redisN) Save(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "Save")
	cmd := c.Client.Save(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Shutdown Shutdown
func (c *redisN) Shutdown(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "Shutdown")
	cmd := c.Client.Shutdown(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ShutdownSave ShutdownSave
func (c *redisN) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ShutdownSave")
	cmd := c.Client.ShutdownSave(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ShutdownNoSave ShutdownNoSave
func (c *redisN) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ShutdownNoSave")
	cmd := c.Client.ShutdownNoSave(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// SlaveOf SlaveOf
func (c *redisN) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "SlaveOf")
	cmd := c.Client.SlaveOf(ctx, host, port)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Time Time
func (c *redisN) Time(ctx context.Context) *redis.TimeCmd {
	defer c._preHandler(ctx, time.Now(), "Time")
	cmd := c.Client.Time(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Eval Eval
func (c *redisN) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	defer c._preHandler(ctx, time.Now(), "Eval", script, keys, args)
	cmd := c.Client.Eval(ctx, script, keys, args...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// EvalSha EvalSha
func (c *redisN) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	defer c._preHandler(ctx, time.Now(), "EvalSha", sha1, keys, args)
	cmd := c.Client.EvalSha(ctx, sha1, keys, args...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ScriptExists ScriptExists
func (c *redisN) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ScriptExists", hashes)
	cmd := c.Client.ScriptExists(ctx, hashes...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ScriptFlush ScriptFlush
func (c *redisN) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ScriptFlush")
	cmd := c.Client.ScriptFlush(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ScriptKill ScriptKill
func (c *redisN) ScriptKill(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ScriptKill")
	cmd := c.Client.ScriptKill(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ScriptLoad ScriptLoad
func (c *redisN) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "ScriptLoad")
	cmd := c.Client.ScriptLoad(ctx, script)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// DebugObject DebugObject
func (c *redisN) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "DebugObject", key)
	cmd := c.Client.DebugObject(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// Publish Publish
func (c *redisN) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "Publish")
	cmd := c.Client.Publish(ctx, channel, message)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PubSubChannels PubSubChannels
func (c *redisN) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "PubSubChannels")
	cmd := c.Client.PubSubChannels(ctx, pattern)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PubSubNumSub PubSubNumSub
func (c *redisN) PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd {
	defer c._preHandler(ctx, time.Now(), "PubSubNumSub")
	cmd := c.Client.PubSubNumSub(ctx, channels...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// PubSubNumPat PubSubNumPat
func (c *redisN) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "PubSubNumPat")
	cmd := c.Client.PubSubNumPat(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterSlots ClusterSlots
func (c *redisN) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterSlots")
	cmd := c.Client.ClusterSlots(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterNodes ClusterNodes
func (c *redisN) ClusterNodes(ctx context.Context) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterNodes")
	cmd := c.Client.ClusterNodes(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterMeet ClusterMeet
func (c *redisN) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterMeet")
	cmd := c.Client.ClusterMeet(ctx, host, port)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterForget ClusterForget
func (c *redisN) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterForget")
	cmd := c.Client.ClusterForget(ctx, nodeID)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterReplicate ClusterReplicate
func (c *redisN) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterReplicate")
	cmd := c.Client.ClusterReplicate(ctx, nodeID)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterResetSoft ClusterResetSoft
func (c *redisN) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterResetSoft")
	cmd := c.Client.ClusterResetSoft(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterResetHard ClusterResetHard
func (c *redisN) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterResetHard")
	cmd := c.Client.ClusterResetHard(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterInfo ClusterInfo
func (c *redisN) ClusterInfo(ctx context.Context) *redis.StringCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterInfo")
	cmd := c.Client.ClusterInfo(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterKeySlot ClusterKeySlot
func (c *redisN) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterKeySlot", key)
	cmd := c.Client.ClusterKeySlot(ctx, key)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterGetKeysInSlot ClusterGetKeysInSlot
func (c *redisN) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterGetKeysInSlot")
	cmd := c.Client.ClusterGetKeysInSlot(ctx, slot, count)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterCountFailureReports ClusterCountFailureReports
func (c *redisN) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterCountFailureReports")
	cmd := c.Client.ClusterCountFailureReports(ctx, nodeID)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterCountKeysInSlot ClusterCountKeysInSlot
func (c *redisN) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterCountKeysInSlot")
	cmd := c.Client.ClusterCountKeysInSlot(ctx, slot)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterDelSlots ClusterDelSlots
func (c *redisN) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterDelSlots")
	cmd := c.Client.ClusterDelSlots(ctx, slots...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterDelSlotsRange ClusterDelSlotsRange
func (c *redisN) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterDelSlotsRange")
	cmd := c.Client.ClusterDelSlotsRange(ctx, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterSaveConfig ClusterSaveConfig
func (c *redisN) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterSaveConfig")
	cmd := c.Client.ClusterSaveConfig(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterSlaves ClusterSlaves
func (c *redisN) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterSlaves", nodeID)
	cmd := c.Client.ClusterSlaves(ctx, nodeID)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterFailover ClusterFailover
func (c *redisN) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterFailover")
	cmd := c.Client.ClusterFailover(ctx)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterAddSlots ClusterAddSlots
func (c *redisN) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterAddSlots", slots)
	cmd := c.Client.ClusterAddSlots(ctx, slots...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// ClusterAddSlotsRange ClusterAddSlotsRange
func (c *redisN) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	defer c._preHandler(ctx, time.Now(), "ClusterAddSlotsRange", min, max)
	cmd := c.Client.ClusterAddSlotsRange(ctx, min, max)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoAdd GeoAdd
func (c *redisN) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "GeoAdd", key, geoLocation)
	cmd := c.Client.GeoAdd(ctx, key, geoLocation...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoPos GeoPos
func (c *redisN) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	defer c._preHandler(ctx, time.Now(), "GeoPos", key, members)
	cmd := c.Client.GeoPos(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoRadius GeoRadius
func (c *redisN) GeoRadius(
	ctx context.Context,
	key string, longitude, latitude float64, query *redis.GeoRadiusQuery,
) *redis.GeoLocationCmd {
	defer c._preHandler(ctx, time.Now(), "GeoRadius", key, longitude, latitude, query)
	cmd := c.Client.GeoRadius(ctx, key, longitude, latitude, query)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoRadiusStore GeoRadiusStore
func (c *redisN) GeoRadiusStore(
	ctx context.Context,
	key string, longitude, latitude float64, query *redis.GeoRadiusQuery,
) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "GeoRadiusStore", key, longitude, latitude, query)
	cmd := c.Client.GeoRadiusStore(ctx, key, longitude, latitude, query)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoRadiusByMember GeoRadiusByMember
func (c *redisN) GeoRadiusByMember(
	ctx context.Context, key, member string, query *redis.GeoRadiusQuery,
) *redis.GeoLocationCmd {
	defer c._preHandler(ctx, time.Now(), "GeoRadiusByMember", key, member, query)
	cmd := c.Client.GeoRadiusByMember(ctx, key, member, query)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoRadiusByMemberStore GeoRadiusByMemberStore
func (c *redisN) GeoRadiusByMemberStore(
	ctx context.Context, key, member string, query *redis.GeoRadiusQuery,
) *redis.IntCmd {
	defer c._preHandler(ctx, time.Now(), "GeoRadiusByMemberStore", key, member, query)
	cmd := c.Client.GeoRadiusByMemberStore(ctx, key, member, query)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoDist GeoDist
func (c *redisN) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	defer c._preHandler(ctx, time.Now(), "GeoDist", key, member1, member2, unit)
	cmd := c.Client.GeoDist(ctx, key, member1, member2, unit)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}

// GeoHash GeoHash
func (c *redisN) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	defer c._preHandler(ctx, time.Now(), "GeoHash", key, members)
	cmd := c.Client.GeoHash(ctx, key, members...)
	// c._afterHandlerCheckPanic(cmd.Err())
	return cmd
}
