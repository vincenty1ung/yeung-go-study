package yredis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/stores/kv"
	"github.com/uncleyeung/yeung-go-study/utils/ylog"
	"go.uber.org/zap"
)

const (
	// RepeatedTimes RepeatedTimes
	RepeatedTimes = 5 // 次数
	// RepeatedInterval RepeatedInterval
	RepeatedInterval = 200 // 毫秒
	// DefaultTime DefaultTime
	DefaultTime = time.Duration(-1) * time.Second
)

// GetString GetString
func GetString(ctx context.Context, con Redis, key string) ([]byte, error) {
	count := 1
	for {
		result, err := con.Get(ctx, key).Bytes()
		if err == nil {
			return result, nil
		}

		if err == Nil {
			return nil, nil
		}

		count++
		if count > RepeatedTimes {
			ylog.Error("GetString失败 ", zap.Any("key", key), zap.Error(err))
			return nil, err
		} else {
			ylog.Error(
				fmt.Sprintf(
					"GetString请求key=%v数据失败，err = %s, "+
						"发起第%v请求", key, err, count,
				),
			)
			time.Sleep(RepeatedInterval * time.Millisecond) // 重试间隔
		}
	}
}

// SetString SetString
func SetString(ctx context.Context, con Redis, key string, value interface{}, ex ...int64) error {
	count := 1
	t := DefaultTime
	if len(ex) > 0 {
		t = time.Duration(ex[0]) * time.Second
	}

	for {
		err := con.Set(ctx, key, value, t).Err()
		if err == nil {
			return nil
		}
		count++
		if count > RepeatedTimes {
			ylog.Error("SetString失败", zap.Any("key", key), zap.Error(err))
			return err
		} else {
			ylog.Error(
				fmt.Sprintf(
					"SetString请求key=%v数据失败，err = %s, "+
						"发起第%v请求", key, err, count,
				),
			)
			time.Sleep(RepeatedInterval * time.Millisecond) // 重试间隔
		}
	}
}

// GetHash GetHash
func GetHash(ctx context.Context, con Redis, key string, field string) ([]byte, error) {
	count := 1
	for {
		result, err := con.HGet(ctx, key, field).Bytes()
		if err == nil {
			return result, nil
		}

		if err == Nil {
			return nil, nil
		}

		count++
		if count > RepeatedTimes {
			ylog.Error(fmt.Sprintf("GetHash请求key=%v数据失败 err=%s", key, err))
			return nil, err
		} else {
			ylog.Error(
				fmt.Sprintf(
					"GetHash请求key=%v数据失败，err = %s, "+
						"发起第%v请求", key, err, count,
				),
			)
			time.Sleep(RepeatedInterval * time.Millisecond) // 重试间隔
		}
	}
}

// Expire 设置过期时间
func Expire(ctx context.Context, con Redis, key string, t int64) error {
	err := con.Expire(ctx, key, time.Duration(t)*time.Second).Err()
	return err
}

// DelKey 删除key
func DelKey(ctx context.Context, con Redis, key ...string) error {
	err := con.Del(ctx, key...).Err()
	return err
}

// GetHashAll GetHashAll
func GetHashAll(ctx context.Context, con Redis, key string) (map[string]string, error) {
	count := 1
	for {
		result, err := con.HGetAll(ctx, key).Result()
		if err == nil {
			return result, nil
		}
		if err == Nil {
			return nil, nil
		}
		count++
		if count > RepeatedTimes {
			ylog.Error(fmt.Sprintf("GetHash请求key=%v数据失败 err=%s", key, err))
			return nil, err
		} else {
			ylog.Error(
				fmt.Sprintf(
					"GetHash请求key=%v数据失败，err = %s, "+
						"发起第%v请求", key, err, count,
				),
			)
			time.Sleep(RepeatedInterval * time.Millisecond) // 重试间隔
		}
	}
}

// GetHashLen 获取hash的长度
func GetHashLen(ctx context.Context, con Redis, key string) (int64, error) {
	result, err := con.HLen(ctx, key).Result()
	return result, err
}

// SetHash SetHash
func SetHash(ctx context.Context, con Redis, key, field string, value []byte) error {
	count := 1
	// if len(ex) > 0 {
	//	t = time.Duration(ex[0]) * time.Second
	// }
	for {
		err := con.HSet(ctx, key, field, value).Err()
		if err == nil {
			return nil
		}
		count++
		if count > RepeatedTimes {
			ylog.Error(fmt.Sprintf("SetHash请求key=%v数据失败 err=%s", key, err))
			return err
		} else {
			ylog.Error(
				fmt.Sprintf(
					"SetHash请求key=%v数据失败，err = %s, "+
						"发起第%v请求", key, err, count,
				),
			)
			time.Sleep(RepeatedInterval * time.Millisecond) // 重试间隔
		}
	}
}

// DelHash 删除hash
func DelHash(ctx context.Context, con Redis, key string, field ...string) error {
	err := con.HDel(ctx, key, field...).Err()
	return err
}

// IncrByHash IncrByHash
func IncrByHash(ctx context.Context, con Redis, key string, field string, incr int64) error {
	err := con.HIncrBy(ctx, key, field, incr).Err()
	if err != nil {
		ylog.Error(fmt.Sprintf("IncrbyHash, key=%s, field=%s, incr=%v, err=%s", key, field, incr, err))
	}
	return err
}

// HMSetHash HMSetHash
func HMSetHash(ctx context.Context, con Redis, key string, fields map[string]interface{}) error {
	err := con.HMSet(ctx, key, fields).Err()
	if err != nil {
		ylog.Error(fmt.Sprintf("HMSetHash, key=%s, field=%v, err=%s", key, fields, err))
	}
	return err
}

// SetSortSet 有序集合
func SetSortSet(ctx context.Context, con Redis, key string, score int64, value interface{}) error {
	err := con.ZAdd(ctx, key, &redis.Z{Score: float64(score), Member: value}).Err()
	if err != nil {
		ylog.Error(
			fmt.Sprintf("设置有序集合数据失败, key=%s, score=%v, value=%v, err=%s", key, score, value, err),
		)
	}
	return err
}

// ZCardSortSet ZCardSortSet
func ZCardSortSet(ctx context.Context, con Redis, key string) int64 {
	return con.ZCard(ctx, key).Val()
}

// SetAllSortSet SetAllSortSet
func SetAllSortSet(ctx context.Context, con Redis, key string, value ...*redis.Z) error {
	err := con.ZAdd(ctx, key, value...).Err()
	if err != nil {
		ylog.Error(fmt.Sprintf("设置有序集合数据失败, key=%s, alue=%v, err=%s", key, value, err))
	}
	return err
}

// ListSortSet ListSortSet
func ListSortSet(ctx context.Context, con Redis, key string, num ...int64) ([]string, error) {
	var n1, n2 int64 = 0, -1
	if len(num) > 0 {
		n1 = num[0]
	}
	if len(num) > 1 {
		n2 = num[1]
	}
	arr, err := con.ZRange(ctx, key, n1, n2).Result()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// ListSortSetRev 获取从大到小的有序集合
func ListSortSetRev(ctx context.Context, con Redis, key string, num ...int64) ([]string, error) {
	var n1, n2 int64 = 0, -1
	if len(num) > 0 {
		n1 = num[0]
	}
	if len(num) > 1 {
		n2 = num[1]
	}
	arr, err := con.ZRevRange(ctx, key, n1, n2).Result()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// DelSortSet DelSortSet
func DelSortSet(ctx context.Context, con Redis, key string, start, end int64) error {
	s := strconv.FormatInt(start, 10)
	e := strconv.FormatInt(end, 10)
	err := con.ZRemRangeByLex(ctx, key, s, e).Err()
	return err
}

// LRange LRange
func LRange(ctx context.Context, con Redis, key string, start int64, stop int64) ([]string, error) {
	arr, err := con.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// Ltrim Ltrim
func Ltrim(ctx context.Context, con Redis, key string, start int64, stop int64) error {
	err := con.LTrim(ctx, key, start, stop).Err()
	return err
}

// RPush RPush
func RPush(ctx context.Context, con Redis, key string, values ...interface{}) error {
	err := con.RPush(ctx, key, values...).Err()
	return err
}

// LPop LPop
func LPop(ctx context.Context, con Redis, key string) (string, error) {
	data, err := con.LPop(ctx, key).Result()
	return data, err
}

// SAdd SAdd
func SAdd(ctx context.Context, con Redis, key string, value interface{}, ex ...int64) error {
	err := con.SAdd(ctx, key, value).Err()
	if err == nil && len(ex) > 0 {
		t := time.Duration(ex[0]) * time.Second
		err = con.Expire(ctx, key, t).Err()
		return err
	}
	return err
}

// SMembers SMembers
func SMembers(ctx context.Context, con Redis, key string) ([]string, error) {
	arr, err := con.SMembers(ctx, key).Result()
	return arr, err
}

// SetNx SetNx
func SetNx(ctx context.Context, con Redis, key string, value interface{}, ex ...int64) (bool, error) {
	if len(ex) > 0 {
		t := time.Duration(ex[0]) * time.Second
		return con.SetNX(ctx, key, value, t).Result()
	} else {
		return con.SetNX(ctx, key, value, 0).Result()
	}
}

// SetNxWait SetNxWait
func SetNxWait(ctx context.Context, con Redis, key string, value interface{}, ex ...int64) (bool, error) {
	inval := 10
	flag := false
	err := kv.ErrNoRedisNode
	for i := 0; i <= 5; i++ {
		flag, err = SetNx(ctx, con, key, value, ex...)
		if err != nil || flag {
			fmt.Println(err)
			return flag, err
		}
		fmt.Println(i)
		time.Sleep(time.Duration(inval) * time.Millisecond)
	}
	if !flag && err == nil {
		return flag, nil
	}
	return flag, err
}

// Keys Keys
func Keys(ctx context.Context, con Redis, key string) ([]string, error) {
	return con.Keys(ctx, key).Result()
}

// Scan Scan
func Scan(ctx context.Context, con Redis, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return con.Scan(ctx, cursor, match, count).Result()
}

// ZRange ZRange
func ZRange(ctx context.Context, con Redis, key string, start int64, end int64) ([]string, error) {
	return con.ZRange(ctx, key, start, end).Result()
}

// ZRem ZRem
func ZRem(ctx context.Context, con Redis, key string, members ...interface{}) error {
	return con.ZRem(ctx, key, members...).Err()
}

// ZCard ZCard
func ZCard(ctx context.Context, con Redis, key string) (int64, error) {
	return con.ZCard(ctx, key).Result()
}

// ZRemRangeByRank ZRemRangeByRank
func ZRemRangeByRank(ctx context.Context, con Redis, key string, start int64, end int64) error {
	return con.ZRemRangeByRank(ctx, key, start, end).Err()
}

// ZScore ZScore
func ZScore(ctx context.Context, con Redis, key string, member string) (float64, error) {
	return con.ZScore(ctx, key, member).Result()
}

// LPush LPush
func LPush(ctx context.Context, con Redis, key string, value interface{}) error {
	return con.LPush(ctx, key, value).Err()
}

// IsExists IsExists
func IsExists(ctx context.Context, con Redis, name string, key string) (bool, error) {
	k := name + key
	ENum, err := con.Exists(ctx, k).Result()
	if err == nil {
		if ENum > 0 {
			return true, err
		}
	}
	return false, err
}

// IncrKey IncrKey
func IncrKey(ctx context.Context, con Redis, name string) (int64, error) {
	k := name
	i, err := con.Incr(ctx, k).Result()
	if err != nil {
		return 0, err
	}
	return i, err
}

// DecrKey DecrKey
func DecrKey(ctx context.Context, con Redis, name string) (int64, error) {
	k := name
	i, err := con.Decr(ctx, k).Result()
	if err != nil {
		return -1, err
	}
	return i, err
}

// Incr Incr
func Incr(ctx context.Context, con Redis, name string, key string) error {
	k := name + key
	err := con.Incr(ctx, k).Err()
	return err
}

// IncrBy IncrBy
func IncrBy(ctx context.Context, con Redis, name string, key string, val int64) (int64, error) {
	k := name + key
	val, err := con.IncrBy(ctx, k, val).Result()
	return val, err
}

// Decr Decr
func Decr(ctx context.Context, con Redis, name string, key string) error {
	k := name + key
	err := con.Decr(ctx, k).Err()
	return err
}
