package test

import (
	"encoding/json"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/vincenty1ung/yeung-go-study/utils/ylog"
	"github.com/vincenty1ung/yeung-go-study/utils/yredis"
)

/*func timeSleep() {
	time.Sleep(10 * time.Second)
}
*/

func found(err error) {
	if yredis.IsNotFound(err) {
		ylog.Error("IsNotFound")
	} else {
		ylog.Info("Found")
	}
}

func TestIsNotFound(t *testing.T) {
	found(nil)
	found(yredis.Nil)
}

func TestGetSet(t *testing.T) {
	type testStruc struct {
		A string
		B int
	}
	ctx1 := ctx
	data, err := json.Marshal(testStruc{"1", 1})
	yredis.AssertError(err)

	key := "test_struct"
	redisX.Set(ctx1, key, data, time.Second*120)

	result := redis.Get(ctx1, key)
	// 这里有坑，.String方法是调试用的，会把命令打印出来
	ylog.Info("info", zap.Any("Val", result.String()))

	// 推荐使用这种方式
	ylog.Info("info", zap.Any("Val", result.Val()))

	str, err := result.Result()
	yredis.AssertError(err)
	ylog.Info("info", zap.String("str", str))

	bytes, err := result.Bytes()
	yredis.AssertError(err)
	ylog.Info("info", zap.String("", string(bytes)))
}

// func TestNewCacheXAndCache(t *testing.T) {
// 	// reqLog := xfilter.PoolGetApiLog()
// 	ctx1 := context.Background()
// 	fmt.Println("redisX")
// 	// redisX
// 	{
// 		set := redisX.Set(ctx1, "test_1", "test_value", time.Second*2)
// 		 ylog.Info(set.Val())
// 		assert.Equal(t, set.Val(), "OK")
//
// 		key1 := "test_1"
// 		get := redisX.Get(ctx1, key1)
// 		 ylog.Info(get.Val())
// 		assert.Equal(t, get.Val(), "test_value")
// 		if xredis.IsNotFound(get.Err()) {
// 			 ylog.Info("not found key: " + key1)
// 		}
//
// 		hmSet := redisX.HMSet(ctx1, "test_2", "key1", "value1", "key3", "value3", "key2", "value2")
// 		 ylog.Info(hmSet.Val())
// 		assert.Equal(t, hmSet.Val(), true)
//
// 		hmGet := redisX.HMGet(ctx1, "test_2", "key1", "key2", "key3")
// 		values := hmGet.Val()
// 		 ylog.Info(values)
// 		assert.Equal(t, values[0].(string), "value1")
//
// 		all := redisX.HGetAll(ctx1, "test_2")
// 		allValues := all.Val()
// 		 ylog.Info(allValues)
// 		assert.Equal(t, allValues["key1"], "value1")
// 	}
//
// 	fmt.Println("redisN")
// 	// redisN
// 	{
// 		ctx2 := context.Background()
// 		result, err2 := redis.HGetAll(ctx2, "test").Result()
// 		xredis.AssertError(err2)
// 		if xredis.IsNotFound(err2) {
// 			 ylog.Error("err", zap.Error(err2))
// 		}
//
// 		 ylog.Info("result",zap.Any("result",result))
// 	}
//
// 	fmt.Println("redisService")
// 	{
// 		man := struct {
// 			Name string `json:"name"`
// 			Age  int    `json:"age"`
// 		}{}
//
// 		redisService := xredis.NewRedisService(redis, RedisKey)
// 		man.Name = "yangbo"
// 		man.Age = 19
// 		err := redisService.HSet(
// 			ctx1, EvaluationPlayerLike, "key1", "f1", man,
// 		)
// 		xredis.AssertError(err)
//
// 		man.Name = "yangbo1"
// 		err = redisService.HSet(
// 			ctx1, EvaluationPlayerLike, "key1", "f2", man,
// 		)
// 		xredis.AssertError(err)
//
// 		man.Name = "yangbo2"
// 		err = redisService.HSet(
// 			ctx1, EvaluationPlayerLike, "key1", "f3", man,
// 		)
// 		xredis.AssertError(err)
//
// 		hGetAll, err := redisService.HGetAll(ctx1, EvaluationPlayerLike, "key1")
// 		xredis.AssertError(err)
// 		assert.Equal(t, len(hGetAll), 3)
// 		 ylog.Info(hGetAll)
//
// 		err = redisService.SetCacheV2(ctx1, EvaluationPlayerMsgListTop200, "", man)
// 		xredis.AssertError(err)
//
// 		v2, err := redisService.GetCacheV2(ctx1, EvaluationPlayerMsgListTop200, "")
// 		 ylog.Info(string(v2), err)
//
// 		existsV2, err := redisService.IsExistsV2(ctx, EvaluationPlayerMsgListTop200, "")
// 		xredis.AssertError(err)
// 		 ylog.Info(existsV2)
//
// 		lock := redisService.NewCacheRedisLock("asd", ctx1)
// 		lock.Lock()
// 		lock.Unlock()
//
// 	}
// }
