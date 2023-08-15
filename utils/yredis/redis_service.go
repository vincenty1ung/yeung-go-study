package yredis

//
// import (
// 	"context"
// 	"errors"
// 	"math/rand"
// 	"time"
// 	"unsafe"
//
// 	"github.com/go-redis/redis/v8"
// 	jsoniter "github.com/json-iterator/go"
// 	"github.com/uncleyeung/yeung-go-study/utils/encoding/ujson"
// 	"github.com/uncleyeung/yeung-go-study/utils/ulog"
// 	"go.uber.org/zap"
// )
//
// var (
// 	// 例子
// 	randExampleMin            = -3
// 	randExampleMax            = 3
// 	randFloatExampleMin       = 0.05
// 	randFloatExampleMax       = -0.05
// 	defaultScanMax            = 100
// 	defaultScanCount    int64 = 500
//
// 	IJson = jsoniter.ConfigCompatibleWithStandardLibrary
// )
//
// type (
// 	RedisService struct {
// 		redisClient Redis
// 		redisKeyMap map[string]int64
// 	}
//
// 	DelQueryObj struct {
// 		Name string
// 		Key  string
// 	}
//
// 	CacheRedisLock struct {
// 		redisLock   *RedisLock
// 		redisClient Redis
// 	}
//
// 	SortSetZ struct {
// 		Score  float64
// 		Member interface{}
// 	}
// )
//
// func NewRedisService(redisClient Redis, redisKeyMap map[string]int64) *RedisService {
// 	return &RedisService{
// 		redisClient: redisClient,
// 		redisKeyMap: redisKeyMap,
// 	}
// }
//
// func (s *RedisService) NewCacheRedisLock(key string, ctx context.Context) *CacheRedisLock {
// 	return &CacheRedisLock{
// 		redisLock:   NewRedisLock(ctx, key),
// 		redisClient: s.redisClient,
// 	}
// }
//
// func (receiver *CacheRedisLock) Lock() bool {
// 	return receiver.redisLock.Lock(receiver.redisClient) == nil
// }
//
// func (receiver *CacheRedisLock) Unlock() bool {
// 	return receiver.redisLock.Unlock(receiver.redisClient) == nil
// }
//
// /*string*/
// func (s *RedisService) GetCache(
// 	ctx context.Context, name string, key string,
// ) ([]byte, error) {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return nil, errors.New("name not exist")
// 	}
// 	k := name + key
// 	b, err := GetString(ctx, s.redisClient, k)
// 	return b, err
// }
// func (s *RedisService) GetCacheV2(
// 	ctx context.Context, name string, key string,
// ) ([]byte, error) {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return nil, errors.New("name not exist")
// 	}
// 	k := name + key
// 	get := s.redisClient.Get(ctx, k)
// 	return get.Bytes()
// }
//
// func (s *RedisService) SetCache(ctx context.Context, name string, key string, value interface{}) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	Expire := s.GeyExTimeInt64ByRedisKeyMapForRandomNum(name)
//
// 	k := name + key
// 	v, err := IJson.Marshal(value)
// 	if err != nil {
// 		xlog.Error2(ctx, "Marshal fail", err)
// 	}
// 	xlog.Info2(ctx, "SetCache", zap.Any("k", k))
// 	if Expire > 0 {
// 		return SetString(ctx, s.redisClient, k, s.BytesToString(v), Expire)
// 	}
// 	return SetString(ctx, s.redisClient, k, s.BytesToString(v))
// }
// func (s *RedisService) SetCacheV3(ctx context.Context, name string, key string, value interface{}) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	Expire := s.GeyExTimeInt64ByRedisKeyMapForRandomNum(name)
// 	/*if !ok {
// 		return errors.New("name not exist")
// 	}*/
// 	k := name + key
// 	v, err := IJson.Marshal(value)
// 	if err != nil {
// 		xlog.Error2(ctx, "Marshal fail", err)
// 	}
// 	if Expire > 0 {
// 		err := SetString(ctx, s.redisClient, k, s.BytesToString(v), Expire)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return SetString(ctx, s.redisClient, k, s.BytesToString(v))
// }
//
// // 支持不过期redisKey
// func (s *RedisService) SetCacheV2(ctx context.Context, name string, key string, value interface{}) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	Expire := s.GeyExTimeByRedisKeyMap(name)
// 	k := name + key
// 	v, err := IJson.Marshal(value)
// 	if err != nil {
// 		xlog.Error2(ctx, "Marshal fail", err)
// 	}
// 	xlog.Info2(ctx, "SetCacheV2 Success", zap.Any("Expire", Expire))
// 	if Expire > 0 {
// 		return SetString(ctx, s.redisClient, k, s.BytesToString(v), int64(Expire))
// 	}
// 	return SetString(ctx, s.redisClient, k, s.BytesToString(v))
// }
//
// func (s *RedisService) DelCache(ctx context.Context, name string, key string) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	k := name + key
// 	xlog.Info2(ctx, "DelCache", zap.Any("k", k))
// 	return DelKey(ctx, s.redisClient, k)
// }
//
// func (s *RedisService) DelCacheV2(ctx context.Context, name string, key string) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	k := name + key
// 	s.redisClient.Del(ctx, k)
// 	return nil
// }
//
// func (s *RedisService) DelByQueryObj(ctx context.Context, delQueryObj []DelQueryObj) error {
// 	keys := make([]string, 0, len(delQueryObj))
// 	for _, obj := range delQueryObj {
// 		_, ok := s.redisKeyMap[obj.Name]
// 		if !ok {
// 			continue
// 		}
// 		keys = append(keys, obj.Name+obj.Key)
// 	}
// 	// if len(keys) < 0 {
// 	//	return errors.New("name not exist")
// 	// }
// 	// xlog.Debug2(ctx, "DelByQueryObj", zap.Any("keys", keys))
// 	return s.redisClient.Del(ctx, keys...).Err()
// }
//
// // Scan 展示不限制扫描次数
// func (s *RedisService) Scan(ctx context.Context, match string) ([]string, error) {
// 	// cursor uint64, match string, count int64
// 	cursor := uint64(0)
// 	strings := make([]string, 0)
// 	for {
// 		result, cursorTmp, err := s.redisClient.Scan(ctx, cursor, match, defaultScanCount).Result()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Scan fail", err)
// 			return nil, err
// 		}
// 		strings = append(strings, result...)
// 		cursor = cursorTmp
// 		if cursorTmp == 0 {
// 			break
// 		}
// 		time.Sleep(time.Millisecond * 30)
// 	}
// 	return strings, nil
// }
//
// func (s *RedisService) Expire(ctx context.Context, preKey, keyStr string, setExpire ...int64) error {
// 	_, ok := s.redisKeyMap[preKey]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	var ExpireTime int64
// 	if len(setExpire) == 0 {
// 		tmpExpire := s.GeyExTimeInt64ByRedisKeyMapForRandomNum(preKey)
// 		/*tmpExpire, ok := RedisKey[preKey]
// 		if !ok {
// 			return errors.New("name not exist")
// 		}*/
// 		ExpireTime = tmpExpire
// 	} else {
// 		ExpireTime = setExpire[0]
// 	}
//
// 	if ExpireTime > 0 {
// 		err := Expire(ctx, s.redisClient, preKey+keyStr, ExpireTime)
// 		if err != nil {
// 			xlog.Error2(ctx, "redis,超时 出错", err)
// 		}
// 		return err
// 	}
// 	return nil
// }
//
// // LimitFrequency 缓存限频
// func (s *RedisService) LimitFrequency(ctx context.Context, key string, duration time.Duration, limit int64) bool {
// 	result, err := s.redisClient.SetNX(ctx, key, 1, duration).Result()
// 	if err != nil {
// 		return false
// 	}
// 	if !result {
// 		if incr, err := s.redisClient.Incr(ctx, key).Result(); err != nil || incr > limit {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// // LimitFrequency 缓存限频
// func (s *RedisService) LimitFrequencyV2(
// 	ctx context.Context, key string, duration time.Duration, limit int64,
// ) bool {
// 	if !s.redisClient.SetNX(ctx, key, 1, duration).Val() {
// 		if incr := s.redisClient.Incr(ctx, key).Val(); incr > limit {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func (s *RedisService) IsExists(ctx context.Context, name string, key string) (bool, error) {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	k := name + key
// 	ENum, err := s.redisClient.Exists(ctx, k).Result()
// 	if err == nil {
// 		if ENum > 0 {
// 			return true, err
// 		}
// 	}
// 	return false, err
// }
// func (s *RedisService) IsExistsV2(ctx context.Context, name string, key string) (bool, error) {
// 	return s.redisClient.Exists(ctx, name+key).Val() > 0, nil
// }
//
// func (s *RedisService) Set(ctx context.Context, key, keyname string, value interface{}) error {
// 	k, exp, bin, err := s.getCacheInfo(key, keyname, value)
// 	if err != nil {
// 		return err
// 	}
// 	return s.redisClient.Set(ctx, k, s.BytesToString(bin), time.Second*time.Duration(exp)).Err()
// }
// func (s *RedisService) SetV2(ctx context.Context, key, keyname string, value interface{}) error {
// 	k, exp, bin, err := s.getCacheInfo(key, keyname, value)
// 	if err != nil {
// 		return err
// 	}
// 	s.redisClient.Set(ctx, k, s.BytesToString(bin), time.Second*time.Duration(exp))
// 	return nil
// }
//
// func (s *RedisService) Get(ctx context.Context, key, keyname string, value interface{}) error {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		return err
// 	}
// 	bin, err := s.redisClient.Get(ctx, k).Bytes()
// 	if err != nil && err != Nil {
// 		return err
// 	}
// 	if len(bin) <= 0 {
// 		return nil
// 	}
// 	return IJson.Unmarshal(bin, value)
// }
//
// func (s *RedisService) GetV2(ctx context.Context, key, keyname string, value interface{}) error {
// 	k := key + keyname
// 	bin := s.redisClient.Get(ctx, k).Val()
// 	if len(bin) <= 0 {
// 		return nil
// 	}
// 	return IJson.Unmarshal(s.StringToBytes(bin), value)
// }
//
// func (s *RedisService) SetInt(ctx context.Context, name string, key string, val int) error {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	expire := s.GeyExTimeInt64ByRedisKeyMapForRandomNum(name)
// 	/*expire, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}*/
// 	k := name + key
// 	return s.redisClient.Set(ctx, k, val, time.Duration(expire)*time.Second).Err()
// }
//
// func (s *RedisService) GetInt(ctx context.Context, name string, key string) (int, error) {
// 	_, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return 0, errors.New("name not exist")
// 	}
// 	k := name + key
// 	num, err := s.redisClient.Get(ctx, k).Int()
// 	if err == Nil {
// 		return 0, nil
// 	}
// 	return num, err
// }
//
// /*list*/
// func (s *RedisService) LRange(ctx context.Context, keyname, key string, pageNum, pageSize int32) ([]string, error) {
// 	keyResult, start, end, err1, done := s.buildRangeValues(ctx, keyname, key, pageNum, pageSize)
// 	if done {
// 		return nil, err1
// 	}
// 	stringSliceCmd := s.redisClient.LRange(ctx, keyResult, int64(start), int64(end))
// 	if err := stringSliceCmd.Err(); err != nil && err != Nil {
// 		xlog.Warn2(ctx, "LRange", err)
// 		return nil, err
// 	}
// 	values := stringSliceCmd.Val()
// 	xlog.Debug2(ctx, "LRange redis", zap.Any("values", values))
// 	return values, nil
// }
// func (s *RedisService) LRangeV2(
// 	ctx context.Context, keyname, key string, pageNum, pageSize int32,
// ) (
// 	[]string, error,
// ) {
// 	keyResult, start, end, err1, done := s.buildRangeValuesV2(ctx, keyname, key, pageNum, pageSize)
// 	if done {
// 		return nil, err1
// 	}
// 	values := s.redisClient.LRange(ctx, keyResult, int64(start), int64(end)).Val()
// 	xlog.Debug2(ctx, "LRange redis", zap.Any("values", values))
// 	xlog.Debug2(ctx, "LRange redis", values)
// 	return values, nil
// }
//
// func (s *RedisService) LGetALL(ctx context.Context, keyname, key string) ([]string, error) {
// 	exists, err := s.IsExists(ctx, keyname, key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if exists {
// 		return nil, nil
// 	}
// 	keyResult := keyname + key
// 	stringSliceCmd := s.redisClient.LRange(ctx, keyResult, int64(0), int64(-1))
// 	if err := stringSliceCmd.Err(); err != nil && err != Nil {
// 		xlog.Warn2(ctx, "MultimediaMusicWorksTagDao", err)
// 		return nil, err
// 	}
// 	values := stringSliceCmd.Val()
// 	xlog.Debug2(ctx, "LRange  redis", zap.Any("values", values))
// 	return values, nil
// }
//
// func (s *RedisService) RPushAll(ctx context.Context, keyname, key string, values []interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "RPushAll redis", zap.Any("key", key), zap.Any("val", values))
// 	exists, err := s.IsExists(ctx, keyname, key)
// 	if err != nil {
// 		xlog.Warn2(ctx, "RPushAll", err)
// 		return false, err
// 	}
// 	if exists {
// 		return false, nil
// 	}
// 	keyResult := keyname + key
// 	err = s.redisClient.RPush(ctx, keyResult, values...).Err()
// 	if err != nil {
// 		xlog.Warn2(ctx, "RPushAll", err)
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) RPush(ctx context.Context, keyname, key string, value interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "RPush redis", zap.Any("key", key), zap.Any("val", value))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.RPush(ctx, keyResult, value).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) LPush(ctx context.Context, keyname, key string, value interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "LPush redis", zap.Any("key", key), zap.Any("val", value))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.LPush(ctx, keyResult, value).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) LPushContainsExpired(
// 	ctx context.Context, keyname, key string, value interface{},
// ) (
// 	bool, error,
// ) {
// 	xlog.Debug2(ctx, "LPushContainsExpired redis", zap.Any("key", key), zap.Any("val", value))
// 	ex, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.LPush(ctx, keyResult, value).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	s.redisClient.Expire(ctx, keyResult, time.Duration(ex)*time.Second)
// 	return true, nil
// }
//
// func (s *RedisService) LPop(ctx context.Context, keyname, key string) (string, error) {
// 	xlog.Debug2(ctx, "LPop redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return "", errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.LPop(ctx, keyResult).Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }
//
// func (s *RedisService) BLPop(ctx context.Context, keyname, key string, timeout time.Duration) ([]string, error) {
// 	xlog.Debug2(ctx, "LPop redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return []string{}, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.BLPop(ctx, timeout, keyResult).Result()
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	return result, nil
// }
//
// func (s *RedisService) LTrim(ctx context.Context, keyname, key string, start, end int64) (string, error) {
// 	xlog.Debug2(ctx, "LTrim redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return "", errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.LTrim(ctx, keyResult, start, end).Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }
//
// func (s *RedisService) LZore(ctx context.Context, keyname, key string) error {
// 	xlog.Debug2(ctx, "LTrim redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	_, err := s.redisClient.LTrim(ctx, keyResult, 1, 0).Result()
// 	if err != nil {
// 		return err
// 	}
// 	marshal, err := ujson.Marshal(
// 		struct {
// 		}{},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	s.redisClient.LPush(ctx, keyResult, marshal)
// 	return nil
// }
// func (s *RedisService) LZoreV2(ctx context.Context, keyname, key string) error {
// 	xlog.Debug2(ctx, "LTrim redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	_ = s.redisClient.LTrim(ctx, keyResult, 1, 0)
// 	marshal, err := ujson.Marshal(
// 		struct {
// 		}{},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	s.redisClient.LPush(ctx, keyResult, marshal)
// 	return nil
// }
//
// // 字段
// func (s *RedisService) LRem(ctx context.Context, keyname, key string, field interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "LRem redis", zap.Any("key", key), zap.Any("val", field))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.LRem(ctx, keyResult, 0, field).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) LLen(ctx context.Context, keyname, key string) (int64, error) {
// 	xlog.Debug2(ctx, "LLen redis", zap.Any("key", key))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return 0, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.LLen(ctx, keyResult).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result, nil
// }
//
// func (s *RedisService) buildRangeValues(
// 	ctx context.Context, keyname string, key string, pageNum int32, pageSize int32,
// ) (
// 	string, int32, int32, error, bool,
// ) {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return "", 0, 0, errors.New("name not exist"), true
// 	}
// 	exists, err := s.IsExists(ctx, keyname, key)
// 	if err != nil {
// 		return "", 0, 0, err, true
// 	}
// 	if !exists {
// 		return "", 0, 0, nil, true
// 	}
// 	keyResult := keyname + key
// 	start := (pageNum - 1) * pageSize
// 	end := start + pageSize - 1
// 	return keyResult, start, end, nil, false
// }
// func (s *RedisService) buildRangeValuesV2(
// 	ctx context.Context, keyname string, key string, pageNum int32,
// 	pageSize int32,
// ) (
// 	string, int32, int32, error, bool,
// ) {
// 	exists, err := s.IsExistsV2(ctx, keyname, key)
// 	if err != nil {
// 		return "", 0, 0, err, true
// 	}
// 	if !exists {
// 		return "", 0, 0, nil, true
// 	}
// 	keyResult := keyname + key
// 	start := (pageNum - 1) * pageSize
// 	end := start + pageSize - 1
// 	return keyResult, start, end, nil, false
// }
//
// /*zset*/
//
// // ZAdd key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
// func (s *RedisService) ZAdd(ctx context.Context, keyname, key string, field ...SortSetZ) (bool, error) {
// 	xlog.Debug2(ctx, "ZAdd redis", zap.Any("key", key), zap.Any("val", field))
// 	exp, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	zSet := make([]*redis.Z, 0, len(field))
// 	for _, z := range field {
// 		zset := new(redis.Z)
// 		zset.Score = z.Score
// 		zset.Member = z.Member
// 		zSet = append(zSet, zset)
// 	}
// 	keyResult := keyname + key
// 	if len(zSet) > 0 {
// 		err := s.redisClient.ZAdd(ctx, keyResult, zSet...).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Redis ZAdd failed", err)
// 			return false, err
// 		}
// 		return true, nil
// 	}
// 	if exp > 0 {
// 		err := s.redisClient.Expire(ctx, keyResult, time.Second*time.Duration(exp)).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Redis ZINCRBY Expire failed", err)
// 		}
// 	}
// 	return false, nil
// }
//
// // ZAddV2 key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
// func (s *RedisService) ZAddXV2(ctx context.Context, keyname, key string, field ...SortSetZ) (
// 	bool, error,
// ) {
// 	xlog.Debug2(ctx, "ZAdd redis", zap.Any("key", key), zap.Any("val", field))
// 	exp, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	zSet := make([]*redis.Z, 0, len(field))
// 	for _, z := range field {
// 		zset := new(redis.Z)
// 		zset.Score = z.Score
// 		zset.Member = z.Member
// 		zSet = append(zSet, zset)
// 	}
// 	keyResult := keyname + key
// 	if len(zSet) > 0 {
// 		s.redisClient.ZAdd(ctx, keyResult, zSet...)
// 		return true, nil
// 	}
// 	if exp > 0 {
// 		s.redisClient.Expire(ctx, keyResult, time.Second*time.Duration(exp))
//
// 	}
// 	return false, nil
// }
//
// // ZRem key member [member ...] 移除有序集合中的一个或多个成员
// func (s *RedisService) ZRem(ctx context.Context, keyname, key string, field ...interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "ZRem redis", zap.Any("key", key), zap.Any("member", field))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.ZRem(ctx, keyResult, field...).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
// func (s *RedisService) ZRemV2(
// 	ctx context.Context, keyname, key string, field ...interface{},
// ) (bool, error) {
// 	xlog.Debug2(ctx, "ZRem redis", zap.Any("key", key), zap.Any("member", field))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	s.redisClient.ZRem(ctx, keyResult, field...)
// 	return true, nil
// }
//
// // ZRem key member [member ...] 移除有序集合中的一个或多个成员
// func (s *RedisService) ZPopMax(ctx context.Context, keyname, key string, count int64) ([]SortSetZ, error) {
// 	xlog.Debug2(ctx, "ZPopMax redis", zap.Any("key", key), zap.Any("count", count))
// 	_, ok := s.redisKeyMap[keyname]
// 	setZS := make([]SortSetZ, 0)
// 	if !ok {
// 		return setZS, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.ZPopMax(ctx, keyResult, count).Result()
// 	if err != nil {
// 		return setZS, err
// 	}
// 	for _, z := range result {
// 		setZS = append(
// 			setZS, SortSetZ{
// 				Score:  z.Score,
// 				Member: z.Member,
// 			},
// 		)
// 	}
// 	return setZS, nil
// }
//
// // ZRem key member [member ...] 移除有序集合中的一个或多个成员
// func (s *RedisService) ZPopMin(ctx context.Context, keyname, key string, count int64) ([]SortSetZ, error) {
// 	xlog.Debug2(ctx, "ZPopMax redis", zap.Any("key", key), zap.Any("count", count))
// 	_, ok := s.redisKeyMap[keyname]
// 	setZS := make([]SortSetZ, 0)
// 	if !ok {
// 		return setZS, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	result, err := s.redisClient.ZPopMin(ctx, keyResult, count).Result()
// 	if err != nil {
// 		return setZS, err
// 	}
// 	for _, z := range result {
// 		setZS = append(
// 			setZS, SortSetZ{
// 				Score:  z.Score,
// 				Member: z.Member,
// 			},
// 		)
// 	}
// 	return setZS, nil
// }
//
// // ZRevRange key start stop [WITHSCORES] 返回有序集中指定区间内的成员，通过索引，分数从高到低
// func (s *RedisService) ZRevRange(ctx context.Context, keyname, key string, pageNum, pageSize int32) ([]string, error) {
// 	keyResult, start, end, err2, done := s.buildRangeValues(ctx, keyname, key, pageNum, pageSize)
// 	if done {
// 		return nil, err2
// 	}
// 	stringSliceCmd := s.redisClient.ZRevRange(ctx, keyResult, int64(start), int64(end))
// 	if err := stringSliceCmd.Err(); err != nil && err != Nil {
// 		xlog.Warn2(ctx, "ZRevRange", err)
// 		return nil, err
// 	}
// 	values := stringSliceCmd.Val()
// 	xlog.Debug2(ctx, "ZRevRange redis", zap.Any("values", values))
// 	return values, nil
// }
//
// // ZRange key start stop [WITHSCORES]  通过索引区间返回有序集合指定区间内的成员
// func (s *RedisService) ZRange(ctx context.Context, keyname, key string, pageNum, pageSize int32) ([]string, error) {
// 	keyResult, start, end, err2, done := s.buildRangeValues(ctx, keyname, key, pageNum, pageSize)
// 	if done {
// 		return nil, err2
// 	}
// 	stringSliceCmd := s.redisClient.ZRange(ctx, keyResult, int64(start), int64(end))
// 	if err := stringSliceCmd.Err(); err != nil && err != Nil {
// 		xlog.Warn2(ctx, "ZRange", err)
// 		return nil, err
// 	}
// 	values := stringSliceCmd.Val()
// 	xlog.Debug2(ctx, "ZRange redis", zap.Any("values", values))
// 	return values, nil
// }
//
// // ZCard key 获取有序集合的成员数
// func (s *RedisService) ZCard(ctx context.Context, keyname, key string) (int64, error) {
// 	xlog.Debug2(ctx, "ZCard redis", zap.Any("key", key), zap.Any("keyname", keyname))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return 0, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	intCmd := s.redisClient.ZCard(ctx, keyResult)
// 	if err := intCmd.Err(); err != nil && err != Nil {
// 		xlog.Warn2(ctx, "ZRange", err)
// 		return 0, err
// 	}
// 	values := intCmd.Val()
// 	xlog.Debug2(ctx, "ZCard redis", zap.Any("values", values))
// 	return values, nil
// }
//
// // 返回有序集中指定区间内的成员，通过索引，分数从高到低
// func (s *RedisService) ZRevRangeWithScores(ctx context.Context, keyname string, start, stop int64) []redis.Z {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return make([]redis.Z, 0)
// 	}
// 	list, err := s.redisClient.ZRevRangeWithScores(ctx, keyname, start, stop).Result()
// 	if err != nil && err != Nil {
// 		xlog.Error2(ctx, "ZRevRangeWithScores  redis错误", err)
// 		return make([]redis.Z, 0)
// 	}
// 	return list
// }
//
// // 获取返回有序集中，成员的分数值
// func (s *RedisService) ZScore(ctx context.Context, keyname string, id string) float64 {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return 0
// 	}
// 	value, err := s.redisClient.ZScore(ctx, keyname, id).Result()
// 	if err == Nil {
// 		return 0
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return 0
// 	}
// 	return value
// }
//
// // 获取返回有序集中，成员的分数值
// func (s *RedisService) ZScoreV3(ctx context.Context, keyname, key string, id string) float64 {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return 0
// 	}
// 	keyname = keyname + key
// 	value, err := s.redisClient.ZScore(ctx, keyname, id).Result()
// 	if err == Nil {
// 		return 0
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return 0
// 	}
// 	return value
// }
//
// func (s *RedisService) ZIsExist(ctx context.Context, keyname, key string, id string) (bool, error) {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, nil
// 	}
// 	keyname = keyname + key
// 	_, err := s.redisClient.ZScore(ctx, keyname, id).Result()
// 	if err == Nil {
// 		return false, nil
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return false, err
// 	}
// 	return true, nil
// }
//
// // 获取返回有序集中，成员的分数值--不判断keyname是否存在
// func (s *RedisService) ZScoreV2(ctx context.Context, keyname string, id string) float64 {
// 	value, err := s.redisClient.ZScore(ctx, keyname, id).Result()
// 	if err == Nil {
// 		return 0
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return 0
// 	}
// 	return value
// }
//
// // ZAdd key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
// func (s *RedisService) ZAddV2(ctx context.Context, keyname, key string, members []*SortSetZ) (bool, error) {
// 	xlog.Debug2(ctx, "ZAdd redis", zap.Any("key", key), zap.Any("val", members))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	zSet := make([]*redis.Z, 0, len(members))
// 	for _, z := range members {
// 		zset := new(redis.Z)
// 		zset.Score = z.Score
// 		zset.Member = z.Member
// 		zSet = append(zSet, zset)
// 	}
// 	if len(zSet) > 0 {
// 		keyResult := keyname + key
// 		err := s.redisClient.ZAdd(ctx, keyResult, zSet...).Err()
// 		if err != nil {
// 			return false, err
// 		}
// 		return true, nil
// 	}
// 	return false, nil
// }
//
// // 有序集合中对指定成员的分数加上增量 increment
// func (s *RedisService) ZINCRBY(ctx context.Context, name, key, sKey string, score float64) error {
// 	exp, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return nil
// 	}
// 	k := name + key
// 	err := s.redisClient.ZIncrBy(ctx, k, score, sKey).Err()
// 	xlog.Info2(
// 		ctx,
// 		"ZINCRBY", zap.Any("name", name), zap.Any("sKey", sKey), zap.Any("score", score), zap.Any("exp", exp),
// 		zap.Any("err", err),
// 	)
// 	if err == Nil {
// 		return nil
// 	}
//
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", k), err)
// 		return err
// 	}
// 	if exp > 0 {
// 		err = s.redisClient.Expire(ctx, k, time.Second*time.Duration(exp)).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Redis ZINCRBY Expire failed", err)
// 		}
// 	}
// 	return nil
// }
//
// // 有序集合中对指定成员的分数加上增量 increment，不检测keyname
// func (s *RedisService) ZINCRBYV2(ctx context.Context, keyname, sKey string, score int64) error {
// 	err := s.redisClient.ZIncrBy(ctx, keyname, float64(score), sKey).Err()
// 	if err == Nil {
// 		return nil
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx, "获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return err
// 	}
// 	return nil
// }
//
// // 有序集合中对指定成员的分数加上增量 increment
// func (s *RedisService) ZINCRBYV3(
// 	ctx context.Context, name, key, sKey string, score float64,
// ) error {
// 	exp, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return nil
// 	}
// 	k := name + key
// 	res := s.redisClient.ZIncrBy(ctx, k, score, sKey)
// 	xlog.Info2(
// 		ctx,
// 		"ZINCRBY", zap.Any("name", name), zap.Any("sKey", sKey), zap.Any("score", score), zap.Any("exp", exp),
// 		zap.Any("res", res),
// 	)
//
// 	if exp > 0 {
// 		s.redisClient.Expire(ctx, k, time.Second*time.Duration(exp))
// 	}
// 	return nil
// }
//
// // Scan 展示不限制扫描次数
// func (s *RedisService) ZScan(ctx context.Context, key, match string) ([]string, error) {
// 	// cursor uint64, match string, count int64
// 	cursor := uint64(0)
// 	strings := make([]string, 0)
// 	for {
// 		result, cursorTmp, err := s.redisClient.ZScan(ctx, key, cursor, match, defaultScanCount).Result()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Scan fail", err)
// 			return nil, err
// 		}
// 		strings = append(strings, result...)
// 		cursor = cursorTmp
// 		if cursorTmp == 0 {
// 			break
// 		}
// 		time.Sleep(time.Millisecond * 30)
// 	}
// 	return strings, nil
// }
//
// /*
// func (s *RedisService) ZIsExistByPipeliner(pipeliner redis.Pipeliner, keyname string, id string) (bool, error) {
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, nil
// 	}
// 	_, err := pipeliner.ZScore(context.Background(), keyname, id).Result()
// 	if err == Nil {
// 		return false, nil
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx,"获取有序集合key出错", zap.Any("keyname", keyname), err)
// 		return false, err
// 	}
// 	return true, nil
// }
//
// // 有序集合中对指定成员的分数加上增量 increment
// func (s *RedisService) ZINCRBYByPipeliner(pipeliner redis.Pipeliner, name, key, sKey string, score float64) error {
// 	ctx := context.Background()
// 	exp, ok := s.redisKeyMap[name]
// 	if !ok {
// 		return nil
// 	}
// 	k := name + key
// 	err := pipeliner.ZIncrBy(ctx,k, score, sKey).Err()
// 	if err == Nil {
// 		return nil
// 	}
// 	if err != nil {
// 		xlog.Error2(ctx,"获取有序集合key出错", zap.Any("keyname", k), err)
// 		return err
// 	}
// 	if exp > 0 {
// 		err = pipeliner.Expire(ctx,k, time.Second*time.Duration(exp)).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx,"Redis ZINCRBY Expire failed", err)
// 		}
// 	}
// 	return nil
// }*/
//
// /*set*/
//
// func (s *RedisService) SAdd(ctx context.Context, keyname, key string, field ...interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "SAdd redis", zap.Any("key", key), zap.Any("val", field))
// 	exp, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.SAdd(ctx, keyResult, field...).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	if exp > 0 {
// 		err = s.redisClient.Expire(ctx, keyResult, time.Second*time.Duration(exp)).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Redis SAdd Expire failed", err)
// 		}
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) SRem(ctx context.Context, keyname, key string, field ...interface{}) (bool, error) {
// 	xlog.Debug2(ctx, "ZRem redis", zap.Any("key", key), zap.Any("member", field))
// 	_, ok := s.redisKeyMap[keyname]
// 	if !ok {
// 		return false, errors.New("name not exist")
// 	}
// 	keyResult := keyname + key
// 	err := s.redisClient.SRem(ctx, keyResult, field...).Err()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
//
// func (s *RedisService) SInter(ctx context.Context, keyname1, key1, keyname2, key2 string) ([]string, error) {
// 	_, ok := s.redisKeyMap[keyname1]
// 	if !ok {
// 		return nil, errors.New("name not exist")
// 	}
// 	keyResult1 := keyname1 + key1
// 	_, ok1 := s.redisKeyMap[keyname2]
// 	if !ok1 {
// 		return nil, errors.New("name not exist")
// 	}
// 	keyResult2 := keyname2 + key2
// 	xlog.Debug2(ctx, "SInter redis", zap.Any("keyResult1", keyResult1), zap.Any("keyResult2", keyResult2))
// 	result, err := s.redisClient.SInter(ctx, keyResult1, keyResult2).Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
//
// /*hash*/
// func (s *RedisService) HDel(ctx context.Context, key string, field ...string) error {
// 	return s.redisClient.HDel(ctx, key, field...).Err()
// }
//
// /*func (s *RedisService) Dump(ctx context.Context,  key string) ([]byte, error) {
// 	res, err := s.redisClient.Dump(key).Bytes()
// 	return res, err
// }
//
// func (s *RedisService) HScan(ctx context.Context,  key string, cursor uint64, count int64, match string) ([]string, uint64, error) {
// 	return s.redisClient.HScan(key, cursor, match, count).Result()
// }*/
//
// /*func (s *RedisService) HLen(ctx context.Context,  key string) (int64, error) {
// 	return client.HLen(key).Result()
// }*/
//
// func (s *RedisService) HSet(ctx context.Context, key, keyname string, field string, value interface{}) error {
// 	k, exp, bin, err := s.getCacheInfo(key, keyname, value)
// 	if err != nil {
//
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return err
// 	}
// 	err = s.redisClient.HSet(ctx, k, field, s.BytesToString(bin)).Err()
// 	if err != nil {
// 		xlog.Warn2(ctx, "Redis HSet failed", err)
// 		return err
// 	}
// 	if exp > 0 {
// 		err = s.redisClient.Expire(ctx, k, time.Second*time.Duration(exp)).Err()
// 		if err != nil {
// 			xlog.Warn2(ctx, "Redis Hash Expire failed", err)
// 		}
// 	}
// 	return nil
// }
// func (s *RedisService) HSetV2(
// 	ctx context.Context, key, keyname string, field string, value interface{},
// ) error {
// 	k, exp, bin, err := s.getCacheInfo(key, keyname, value)
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return err
// 	}
// 	s.redisClient.HSet(ctx, k, field, s.BytesToString(bin))
// 	if exp > 0 {
// 		s.redisClient.Expire(ctx, k, time.Second*time.Duration(exp))
// 	}
// 	return nil
// }
// func (s *RedisService) HIncrBy(ctx context.Context, key, keyname string, field string, incr int64) (int64, error) {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return 0, err
// 	}
// 	redisIncr, err := s.redisClient.HIncrBy(ctx, k, field, incr).Result()
// 	if err != nil {
// 		xlog.Warn2(ctx, "Redis HIncrBy failed", err)
// 		return 0, err
// 	}
// 	return redisIncr, nil
// }
//
// func (s *RedisService) HGet(
// 	ctx context.Context, key, keyname string, field string, value interface{},
// ) error {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return err
// 	}
// 	bin, err := s.redisClient.HGet(ctx, k, field).Result()
// 	if err != nil && err != Nil {
// 		xlog.Warn2(ctx, "redis HGet failed", err)
// 		return err
// 	}
// 	if len(bin) == 0 {
// 		return nil
// 	}
// 	err = IJson.Unmarshal([]byte(bin), value)
// 	if err != nil {
// 		xlog.Warn2(ctx, "bytes IJson.Unmarshal failed", err)
// 	}
// 	return err
// }
//
// func (s *RedisService) HGetV2(
// 	ctx context.Context, key, keyname string, field string, value interface{},
// ) error {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return err
// 	}
// 	bin := s.redisClient.HGet(ctx, k, field).Val()
// 	if len(bin) == 0 {
// 		return nil
// 	}
// 	err = IJson.Unmarshal([]byte(bin), value)
// 	if err != nil {
// 		xlog.Warn2(ctx, "bytes IJson.Unmarshal failed", err)
// 	}
// 	return err
// }
// func (s *RedisService) HMGet(ctx context.Context, key, keyname string, field []string) ([]interface{}, error) {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis key不存在", err)
// 		return nil, err
// 	}
// 	result, err := s.redisClient.HMGet(ctx, k, field...).Result()
// 	if err != nil && err != Nil {
// 		xlog.Warn2(ctx, "redis HMGet failed", err)
// 		return nil, err
// 	}
// 	if len(result) == 0 {
// 		return nil, nil
// 	}
// 	return result, err
// }
//
// func (s *RedisService) HExists(
// 	ctx context.Context, key, keyname string, field string,
// ) bool {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Error2(ctx, "redis key不存在", zap.Any("key", key), zap.Any("keyname", keyname), err)
// 		return false
// 	}
// 	isExist, err := s.redisClient.HExists(ctx, k, field).Result()
// 	if err != nil {
// 		xlog.Warn2(ctx, "redis HGet failed", err)
// 		return false
// 	}
// 	return isExist
// }
// func (s *RedisService) HExistsV2(
// 	ctx context.Context, key, keyname string, field string,
// ) bool {
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		xlog.Error2(ctx, "redis key不存在", zap.Any("key", key), zap.Any("keyname", keyname), err)
// 		return false
// 	}
// 	isExist := s.redisClient.HExists(ctx, k, field)
// 	/*if err != nil {
// 		xlog.Warn2(ctx,"redis HGet failed", err)
// 		return false
// 	}*/
// 	return isExist.Val()
// }
//
// // 用HScan实现HGetAll
// func (s *RedisService) HGetAll(ctx context.Context, key, keyname string) ([]string, error) {
// 	result := make([]string, 0)
// 	k, _, _, err := s.getCacheInfo(key, keyname, nil)
// 	if err != nil {
// 		return result, err
// 	}
//
// 	scanCount := 1
// 	cursor := uint64(0)
// 	for scanCount <= defaultScanMax {
// 		list, TmpCursor, err := s.redisClient.HScan(ctx, k, cursor, "", defaultScanCount).Result()
// 		if err != nil && err != Nil {
// 			xlog.Warn2(ctx, "redis HScan failed", err)
// 			return result, err
// 		}
// 		if len(list) > 0 {
// 			for i, v := range list {
// 				if i%2 == 0 {
// 					continue
// 				}
// 				result = append(result, v)
// 			}
// 		}
// 		if TmpCursor == 0 {
// 			return result, nil
// 		}
// 		cursor = TmpCursor
// 		xlog.Debug2(
// 			ctx,
// 			"redis HSet 还有数据", zap.String("key", k), zap.Uint64("cursor", cursor), zap.Int("scan次数", scanCount),
// 		)
// 		scanCount++
// 		time.Sleep(time.Millisecond * 5)
// 	}
// 	return result, nil
// }
//
// /*common*/
// func (s *RedisService) getCacheInfo(key, keyname string, value interface{}) (string, int64, []byte, error) {
// 	exp := s.GeyExTimeInt64ByRedisKeyMapForRandomNum(key)
// 	/*exp, ok := RedisKey[key]
// 	if !ok {
// 		return "", 0, []byte{}, RedisKeyNotRegister
// 	}*/
// 	bin := make([]byte, 0)
// 	if value != nil {
// 		marshalBin, err := IJson.Marshal(value)
// 		if err != nil {
// 			ulog.Error("Marshal fail", err)
// 		}
// 		bin = marshalBin
// 	}
// 	return key + keyname, exp, bin, nil
// }
//
// // GeyExTimeByRedisKeyMap 获取name过期时间
// func (s *RedisService) GeyExTimeByRedisKeyMap(name string) time.Duration {
// 	tmp := s.redisKeyMap[name]
// 	return time.Duration(tmp) * time.Second
// }
//
// // GeyExTimeByRedisKeyMapForRandomNum 随机正负百分之5的过期时间防止雪崩
// func (s *RedisService) GeyExTimeByRedisKeyMapForRandomNum(name string) time.Duration {
// 	nameEx := s.redisKeyMap[name]
// 	nameExFloat := float64(nameEx)
// 	tmpExFloat := nameExFloat * s.RandFloat64ByRegulations()
// 	ex := nameExFloat + tmpExFloat
// 	return time.Duration(ex) * time.Second
// }
//
// // GeyExTimeInt64ByRedisKeyMapForRandomNum 随机正负百分之5的过期时间防止雪崩
// func (s *RedisService) GeyExTimeInt64ByRedisKeyMapForRandomNum(name string) int64 {
// 	nameEx := s.redisKeyMap[name]
// 	// 主动设置过期时间为永久
// 	if nameEx == -1 {
// 		return nameEx
// 	}
// 	nameExFloat := float64(nameEx)
// 	randFloat64ByRegulations := s.RandFloat64ByRegulations()
// 	tmpExFloat := nameExFloat * randFloat64ByRegulations
// 	ex := nameExFloat + tmpExFloat
// 	// 当数据小于1时,初始化一秒过期时间,防止永久缓存
// 	if ex < 1 {
// 		return 1
// 	}
// 	return int64(ex)
// }
//
// // s.BytesToString converts byte slice to string.
// func (s *RedisService) BytesToString(b []byte) string {
// 	return *(*string)(unsafe.Pointer(&b))
// }
//
// // StringToBytes converts string to byte slice.
// func (sr RedisService) StringToBytes(s string) []byte {
// 	return *(*[]byte)(unsafe.Pointer(
// 		&struct {
// 			string
// 			Cap int
// 		}{s, len(s)},
// 	))
// }
//
// func (s *RedisService) RandInt64ByRegulations() int64 {
// 	return int64(s.RandIntByMinAndMax(randExampleMin, randExampleMax))
// }
//
// func (s *RedisService) RandIntByRegulations() int64 {
// 	return int64(s.RandIntByMinAndMax(randExampleMin, randExampleMax))
// }
//
// func (s *RedisService) RandIntByMinAndMax(min, max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	if min >= max || min == 0 || max == 0 {
// 		return max
// 	}
// 	return rand.Intn(max-min) + min
// }
//
// func (s *RedisService) RandFloat64ByRegulations() float64 {
// 	// Seed uses the provided seed value to initialize the generator to a deterministic state.
// 	rand.Seed(time.Now().UnixNano())
// 	return s.RandFloatByMinAndMax(randFloatExampleMin, randFloatExampleMax)
// }
//
// func (s *RedisService) RandFloatByMinAndMax(min, max float64) float64 {
// 	return min + rand.Float64()*(max-min)
// }
