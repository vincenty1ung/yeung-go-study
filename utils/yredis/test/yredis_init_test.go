package test

import (
	"context"

	"github.com/vincenty1ung/yeung-go-study/utils/yredis"
)

var (
	ctx      = context.Background()
	redisX   yredis.Redis
	redis    yredis.Redis
	RedisKey = map[string]int64{
		EvaluationPlayerLike:          3600,
		EvaluationPlayerMsgListTop200: 86400 * 7, // 7天有效期(存在增量延长七天的过期时间)

	}
)

const (
	// Evaluation
	EvaluationPlayerLike = "evaluation:playerLike:" // 玩家对动态的点赞

	EvaluationPlayerMsgListTop200 = "evaluation:playerMsgListTop50:" // 评论消息Top50

)

func init() {

	rc, err := yredis.NewClient("192.168.221.129:6379", "", 2)
	yredis.AssertError(err)
	redisX = yredis.NewRedisX(rc)
	redis = yredis.NewRedisN(rc)

}
