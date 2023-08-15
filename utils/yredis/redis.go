package yredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	// RedisBase

	// NewContext(ctx context.Context) (Redis, error)
	WithTimeout(ctx context.Context, timeout time.Duration) Redis
	WithContext(ctx context.Context) Redis
	Clone(ctx context.Context) Redis
	Do(ctx context.Context, args ...interface{}) *redis.Cmd
	DoContext(ctx context.Context, args ...interface{}) *redis.Cmd
	Command(ctx context.Context) *redis.CommandsInfoCmd
	ClientGetName(ctx context.Context) *redis.StringCmd
	Echo(ctx context.Context, message interface{}) *redis.StringCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Quit(ctx context.Context) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Unlink(ctx context.Context, keys ...string) *redis.IntCmd
	Dump(ctx context.Context, key string) *redis.StringCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd
	Move(ctx context.Context, key string, db int) *redis.BoolCmd
	ObjectRefCount(ctx context.Context, key string) *redis.IntCmd
	ObjectEncoding(ctx context.Context, key string) *redis.StringCmd
	ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd
	Persist(ctx context.Context, key string) *redis.BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	PTTL(ctx context.Context, key string) *redis.DurationCmd
	RandomKey(ctx context.Context) *redis.StringCmd
	Rename(ctx context.Context, key, newkey string) *redis.StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd
	SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd
	Touch(ctx context.Context, keys ...string) *redis.IntCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd
	Type(ctx context.Context, key string) *redis.StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	Append(ctx context.Context, key, value string) *redis.IntCmd
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd
	BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd
	GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd
	StrLen(ctx context.Context, key string) *redis.IntCmd
	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	HExists(ctx context.Context, key, field string) *redis.BoolCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd
	HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	HLen(ctx context.Context, key string) *redis.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd
	HVals(ctx context.Context, key string) *redis.StringSliceCmd
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd
	LIndex(ctx context.Context, key string, index int64) *redis.StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd
	RPop(ctx context.Context, key string) *redis.StringCmd
	RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SCard(ctx context.Context, key string) *redis.IntCmd
	SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd
	SMembers(ctx context.Context, key string) *redis.StringSliceCmd
	SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd
	SPop(ctx context.Context, key string) *redis.StringCmd
	SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRandMember(ctx context.Context, key string) *redis.StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
	XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd
	XLen(ctx context.Context, stream string) *redis.IntCmd
	XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd
	XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd
	XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd
	XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd
	XPending(ctx context.Context, stream, group string) *redis.XPendingCmd
	XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd
	XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd
	XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd
	XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd
	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZCard(ctx context.Context, key string) *redis.IntCmd
	ZCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRank(ctx context.Context, key, member string) *redis.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *redis.IntCmd
	ZScore(ctx context.Context, key, member string) *redis.FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd
	PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd
	PFCount(ctx context.Context, keys ...string) *redis.IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd
	BgRewriteAOF(ctx context.Context) *redis.StatusCmd
	BgSave(ctx context.Context) *redis.StatusCmd
	ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd
	ClientList(ctx context.Context) *redis.StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd
	ClientID(ctx context.Context) *redis.IntCmd
	ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd
	ConfigResetStat(ctx context.Context) *redis.StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd
	ConfigRewrite(ctx context.Context) *redis.StatusCmd
	DBSize(ctx context.Context) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
	FlushAllAsync(ctx context.Context) *redis.StatusCmd
	FlushDB(ctx context.Context) *redis.StatusCmd
	FlushDBAsync(ctx context.Context) *redis.StatusCmd
	Info(ctx context.Context, section ...string) *redis.StringCmd
	LastSave(ctx context.Context) *redis.IntCmd
	Save(ctx context.Context) *redis.StatusCmd
	Shutdown(ctx context.Context) *redis.StatusCmd
	ShutdownSave(ctx context.Context) *redis.StatusCmd
	ShutdownNoSave(ctx context.Context) *redis.StatusCmd
	SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd
	Time(ctx context.Context) *redis.TimeCmd
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd
	ScriptFlush(ctx context.Context) *redis.StatusCmd
	ScriptKill(ctx context.Context) *redis.StatusCmd
	ScriptLoad(ctx context.Context, script string) *redis.StringCmd
	DebugObject(ctx context.Context, key string) *redis.StringCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd
	PubSubNumPat(ctx context.Context) *redis.IntCmd
	ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd
	ClusterNodes(ctx context.Context) *redis.StringCmd
	ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterResetSoft(ctx context.Context) *redis.StatusCmd
	ClusterResetHard(ctx context.Context) *redis.StatusCmd
	ClusterInfo(ctx context.Context) *redis.StringCmd
	ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd
	ClusterSaveConfig(ctx context.Context) *redis.StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd
	ClusterFailover(ctx context.Context) *redis.StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd
	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd
	GeoRadius(
		ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery,
	) *redis.GeoLocationCmd
	GeoRadiusStore(
		ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery,
	) *redis.IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd
	ReadOnly(ctx context.Context) *redis.StatusCmd
	ReadWrite(ctx context.Context) *redis.StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd
}

// type redisX interface {
//	RedisBase
//
//	NewContext(ctx context.Context) (redisX, error)
//	WithTimeout(ctx context.Context, timeout time.Duration) redisX
//	WithContext(ctx context.Context) redisX
//	Clone(ctx context.Context) redisX
//	Do(ctx context.Context, args ...interface{}) interface{}
//	DoContext(ctx context.Context, args ...interface{}) interface{}
//	Command(ctx context.Context) map[string]*redis.CommandInfo
//	ClientGetName(ctx context.Context) string
//	Echo(ctx context.Context, message interface{}) string
//	Ping(ctx context.Context) string
//	Quit(ctx context.Context) string
//	Del(ctx context.Context, keys ...string) int64
//	Unlink(ctx context.Context, keys ...string) int64
//	Dump(ctx context.Context, key string) string
//	Exists(ctx context.Context, keys ...string) int64
//	Expire(ctx context.Context, key string, expiration time.Duration) bool
//	ExpireAt(ctx context.Context, key string, tm time.Time) bool
//	Keys(ctx context.Context, pattern string) []string
//	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) string
//	Move(ctx context.Context, key string, db int) bool
//	ObjectRefCount(ctx context.Context, key string) int64
//	ObjectEncoding(ctx context.Context, key string) string
//	ObjectIdleTime(ctx context.Context, key string) time.Duration
//	Persist(ctx context.Context, key string) bool
//	PExpire(ctx context.Context, key string, expiration time.Duration) bool
//	PExpireAt(ctx context.Context, key string, tm time.Time) bool
//	PTTL(ctx context.Context, key string) time.Duration
//	RandomKey(ctx context.Context) string
//	Rename(ctx context.Context, key, newkey string) string
//	RenameNX(ctx context.Context, key, newkey string) bool
//	Restore(ctx context.Context, key string, ttl time.Duration, value string) string
//	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) string
//	Sort(ctx context.Context, key string, sort *redis.Sort) []string
//	SortStore(ctx context.Context, key, store string, sort *redis.Sort) int64
//	SortInterfaces(ctx context.Context, key string, sort *redis.Sort) []interface{}
//	Touch(ctx context.Context, keys ...string) int64
//	TTL(ctx context.Context, key string) time.Duration
//	Type(ctx context.Context, key string) string
//	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64)
//	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
//	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
//	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64)
//	Append(ctx context.Context, key, value string) int64
//	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) int64
//	BitOpAnd(ctx context.Context, destKey string, keys ...string) int64
//	BitOpOr(ctx context.Context, destKey string, keys ...string) int64
//	BitOpXor(ctx context.Context, destKey string, keys ...string) int64
//	BitOpNot(ctx context.Context, destKey string, key string) int64
//	BitPos(ctx context.Context, key string, bit int64, pos ...int64) int64
//	BitField(ctx context.Context, key string, args ...interface{}) []int64
//	Decr(ctx context.Context, key string) int64
//	DecrBy(ctx context.Context, key string, decrement int64) int64
//	Get(ctx context.Context, key string) string
//	GetBit(ctx context.Context, key string, offset int64) int64
//	GetRange(ctx context.Context, key string, start, end int64) string
//	GetSet(ctx context.Context, key string, value interface{}) string
//	Incr(ctx context.Context, key string) int64
//	IncrBy(ctx context.Context, key string, value int64) int64
//	IncrByFloat(ctx context.Context, key string, value float64) float64
//	MGet(ctx context.Context, keys ...string) []interface{}
//	MSet(ctx context.Context, values ...interface{}) string
//	MSetNX(ctx context.Context, values ...interface{}) bool
//	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) string
//	SetBit(ctx context.Context, key string, offset int64, value int) int64
//	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
//	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
//	SetRange(ctx context.Context, key string, offset int64, value string) int64
//	StrLen(ctx context.Context, key string) int64
//	HDel(ctx context.Context, key string, fields ...string) int64
//	HExists(ctx context.Context, key, field string) bool
//	HGet(ctx context.Context, key, field string) string
//	HGetAll(ctx context.Context, key string) map[string]string
//	HIncrBy(ctx context.Context, key, field string, incr int64) int64
//	HIncrByFloat(ctx context.Context, key, field string, incr float64) float64
//	HKeys(ctx context.Context, key string) []string
//	HLen(ctx context.Context, key string) int64
//	HMGet(ctx context.Context, key string, fields ...string) []interface{}
//	HSet(ctx context.Context, key string, values ...interface{}) int64
//	HMSet(ctx context.Context, key string, values ...interface{}) bool
//	HSetNX(ctx context.Context, key, field string, value interface{}) bool
//	HVals(ctx context.Context, key string) []string
//	BLPop(ctx context.Context, timeout time.Duration, keys ...string) []string
//	BRPop(ctx context.Context, timeout time.Duration, keys ...string) []string
//	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) string
//	LIndex(ctx context.Context, key string, index int64) string
//	LInsert(ctx context.Context, key, op string, pivot, value interface{}) int64
//	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) int64
//	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) int64
//	LLen(ctx context.Context, key string) int64
//	LPop(ctx context.Context, key string) string
//	LPush(ctx context.Context, key string, values ...interface{}) int64
//	LPushX(ctx context.Context, key string, values ...interface{}) int64
//	LRange(ctx context.Context, key string, start, stop int64) []string
//	LRem(ctx context.Context, key string, count int64, value interface{}) int64
//	LSet(ctx context.Context, key string, index int64, value interface{}) string
//	LTrim(ctx context.Context, key string, start, stop int64) string
//	RPop(ctx context.Context, key string) string
//	RPopLPush(ctx context.Context, source, destination string) string
//	RPush(ctx context.Context, key string, values ...interface{}) int64
//	RPushX(ctx context.Context, key string, values ...interface{}) int64
//	SAdd(ctx context.Context, key string, members ...interface{}) int64
//	SCard(ctx context.Context, key string) int64
//	SDiff(ctx context.Context, keys ...string) []string
//	SDiffStore(ctx context.Context, destination string, keys ...string) int64
//	SInter(ctx context.Context, keys ...string) []string
//	SInterStore(ctx context.Context, destination string, keys ...string) int64
//	SIsMember(ctx context.Context, key string, member interface{}) bool
//	SMembers(ctx context.Context, key string) []string
//	SMembersMap(ctx context.Context, key string) map[string]struct{}
//	SMove(ctx context.Context, source, destination string, member interface{}) bool
//	SPop(ctx context.Context, key string) string
//	SPopN(ctx context.Context, key string, count int64) []string
//	SRandMember(ctx context.Context, key string) string
//	SRandMemberN(ctx context.Context, key string, count int64) []string
//	SRem(ctx context.Context, key string, members ...interface{}) int64
//	SUnion(ctx context.Context, keys ...string) []string
//	SUnionStore(ctx context.Context, destination string, keys ...string) int64
//	XAdd(ctx context.Context, a *redis.XAddArgs) string
//	XDel(ctx context.Context, stream string, ids ...string) int64
//	XLen(ctx context.Context, stream string) int64
//	XRange(ctx context.Context, stream, start, stop string) []redis.XMessage
//	XRangeN(ctx context.Context, stream, start, stop string, count int64) []redis.XMessage
//	XRevRange(ctx context.Context, stream string, start, stop string) []redis.XMessage
//	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) []redis.XMessage
//	XRead(ctx context.Context, a *redis.XReadArgs) []redis.XStream
//	XReadStreams(ctx context.Context, streams ...string) []redis.XStream
//	XGroupCreate(ctx context.Context, stream, group, start string) string
//	XGroupCreateMkStream(ctx context.Context, stream, group, start string) string
//	XGroupSetID(ctx context.Context, stream, group, start string) string
//	XGroupDestroy(ctx context.Context, stream, group string) int64
//	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) int64
//	XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) []redis.XStream
//	XAck(ctx context.Context, stream, group string, ids ...string) int64
//	XPending(ctx context.Context, stream, group string) *redis.XPending
//	XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) []redis.XPendingExt
//	XClaim(ctx context.Context, a *redis.XClaimArgs) []redis.XMessage
//	XClaimJustID(ctx context.Context, a *redis.XClaimArgs) []string
//	XTrim(ctx context.Context, key string, maxLen int64) int64
//	XTrimApprox(ctx context.Context, key string, maxLen int64) int64
//	XInfoGroups(ctx context.Context, key string) []redis.XInfoGroup
//	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKey
//	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKey
//	ZAdd(ctx context.Context, key string, members ...*redis.Z) int64
//	ZAddNX(ctx context.Context, key string, members ...*redis.Z) int64
//	ZAddXX(ctx context.Context, key string, members ...*redis.Z) int64
//	ZAddCh(ctx context.Context, key string, members ...*redis.Z) int64   // Deprecated: Use
//	ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) int64 // Deprecated: Use
//	ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) int64 // Deprecated: Use
//	ZIncr(ctx context.Context, key string, member *redis.Z) float64
//	ZIncrNX(ctx context.Context, key string, member *redis.Z) float64
//	ZIncrXX(ctx context.Context, key string, member *redis.Z) float64
//	ZCard(ctx context.Context, key string) int64
//	ZCount(ctx context.Context, key, min, max string) int64
//	ZLexCount(ctx context.Context, key, min, max string) int64
//	ZIncrBy(ctx context.Context, key string, increment float64, member string) float64
//	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) int64
//	ZPopMax(ctx context.Context, key string, count ...int64) []redis.Z
//	ZPopMin(ctx context.Context, key string, count ...int64) []redis.Z
//	ZRange(ctx context.Context, key string, start, stop int64) []string
//	ZRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z
//	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) []string
//	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) []string
//	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) []redis.Z
//	ZRank(ctx context.Context, key, member string) int64
//	ZRem(ctx context.Context, key string, members ...interface{}) int64
//	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) int64
//	ZRemRangeByScore(ctx context.Context, key, min, max string) int64
//	ZRemRangeByLex(ctx context.Context, key, min, max string) int64
//	ZRevRange(ctx context.Context, key string, start, stop int64) []string
//	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) []redis.Z
//	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) []string
//	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) []string
//	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) []redis.Z
//	ZRevRank(ctx context.Context, key, member string) int64
//	ZScore(ctx context.Context, key, member string) float64
//	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) int64
//	/*PFAdd(ctx context.Context, key string, els ...interface{}) int64
//	PFCount(ctx context.Context, keys ...string) int64
//	PFMerge(ctx context.Context, dest string, keys ...string) string
//	BgRewriteAOF(ctx context.Context) string
//	BgSave(ctx context.Context) string
//	ClientKill(ctx context.Context, ipPort string) string
//	ClientKillByFilter(ctx context.Context, keys ...string) int64
//	ClientList(ctx context.Context) string
//	ClientPause(ctx context.Context, dur time.Duration) bool
//	ClientID(ctx context.Context) int64
//	ConfigGet(ctx context.Context, parameter string) []interface{}
//	ConfigResetStat(ctx context.Context) string
//	ConfigSet(ctx context.Context, parameter, value string) string
//	ConfigRewrite(ctx context.Context) string
//	DBSize(ctx context.Context) int64
//	FlushAll(ctx context.Context) string
//	FlushAllAsync(ctx context.Context) string
//	FlushDB(ctx context.Context) string
//	FlushDBAsync(ctx context.Context) string
//	Info(ctx context.Context, section ...string) string
//	LastSave(ctx context.Context) int64
//	Save(ctx context.Context) string
//	Shutdown(ctx context.Context) string
//	ShutdownSave(ctx context.Context) string
//	ShutdownNoSave(ctx context.Context) string
//	SlaveOf(ctx context.Context, host, port string) string
//	Time(ctx context.Context) *redis.TimeCmd
//	Eval(ctx context.Context, script string, keys []string, args ...interface{}) interface{}
//	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) interface{}
//	ScriptExists(ctx context.Context, hashes ...string) []bool
//	ScriptFlush(ctx context.Context) string
//	ScriptKill(ctx context.Context) string
//	ScriptLoad(ctx context.Context, script string) string
//	DebugObject(ctx context.Context, key string) string
//	Publish(ctx context.Context, channel string, message interface{}) int64
//	PubSubChannels(ctx context.Context, pattern string) []string
//	PubSubNumSub(ctx context.Context, channels ...string) map[string]int64
//	PubSubNumPat(ctx context.Context) int64
//	ClusterSlots(ctx context.Context) []redis.ClusterSlot
//	ClusterNodes(ctx context.Context) string
//	ClusterMeet(ctx context.Context, host, port string) string
//	ClusterForget(ctx context.Context, nodeID string) string
//	ClusterReplicate(ctx context.Context, nodeID string) string
//	ClusterResetSoft(ctx context.Context) string
//	ClusterResetHard(ctx context.Context) string
//	ClusterInfo(ctx context.Context) string
//	ClusterKeySlot(ctx context.Context, key string) int64
//	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) []string
//	ClusterCountFailureReports(ctx context.Context, nodeID string) int64
//	ClusterCountKeysInSlot(ctx context.Context, slot int) int64
//	ClusterDelSlots(ctx context.Context, slots ...int) string
//	ClusterDelSlotsRange(ctx context.Context, min, max int) string
//	ClusterSaveConfig(ctx context.Context) string
//	ClusterSlaves(ctx context.Context, nodeID string) []string
//	ClusterFailover(ctx context.Context) string
//	ClusterAddSlots(ctx context.Context, slots ...int) string
//	ClusterAddSlotsRange(ctx context.Context, min, max int) string
//	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) int64
//	GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd
//	GeoRadius(
//		ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery,
//	) []redis.GeoLocation
//	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) int64
//	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) []redis.GeoLocation
//	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) int64
//	GeoDist(ctx context.Context, key string, member1, member2, unit string) float64
//	GeoHash(ctx context.Context, key string, members ...string) []string
//	ReadOnly(ctx context.Context) string
//	ReadWrite(ctx context.Context) string
//	MemoryUsage(ctx context.Context, key string, samples ...int) int64*/
// }
