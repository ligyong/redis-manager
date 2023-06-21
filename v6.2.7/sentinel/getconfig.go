package sentinel

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ligyong/redis-manager/common"
)

type RedisConfigGet struct {
	Node       []string
	MasterName string
	Password   string
	result     map[string]string
	err        error
}

func (r *RedisConfigGet) Do() common.RedisOperatorResultMap {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    r.MasterName,
		SentinelAddrs: r.Node,
		Password:      r.Password,
	})

	result, err := client.ConfigGet(context.TODO(), "*").Result()
	if err != nil {
		r.err = err
		return r
	}

	configMap := make(map[string]string)
	for i := 0; i < len(result); i = i + 2 {
		key := fmt.Sprintf("%v", result[i])
		value := fmt.Sprintf("%v", result[i+1])
		configMap[key] = value
	}

	return r
}

func (r *RedisConfigGet) Result() (map[string]string, error) {
	return r.result, r.err
}
