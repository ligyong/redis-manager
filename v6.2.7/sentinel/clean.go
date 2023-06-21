package sentinel

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ligyong/redis-manager/common"
)

type RedisDataClean struct {
	SentinelNode []string
	Password     string
	MasterName   string
	err          error
}

func (r *RedisDataClean) Do() common.RedisOperatorResult {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    r.MasterName,
		SentinelAddrs: r.SentinelNode,
		Password:      r.Password,
	})

	if _, err := client.FlushAll(context.TODO()).Result(); err != nil {
		r.err = err
	}

	return r
}

func (r *RedisDataClean) Result() error {
	return r.err
}
