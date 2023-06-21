package sentinel

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ligyong/redis-manager/common"
	"strconv"
	"strings"
)

type RedisSentinelHealth struct {
	RedisNum     int
	SentinelNode []string
	Password     string
	MasterName   string
	err          error
}

func (r *RedisSentinelHealth) Do() common.RedisOperatorResult {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"172.16.200.30:26379", "172.16.200.31:26379", "172.16.200.32:26379"},
		Password:      "Sobey@1234",
	})

	defer client.Close()

	result, err := client.Info(context.TODO(), "replication").Result()
	if err != nil {
		r.err = err
		return r
	}

	results := strings.Split(result, "\r\n")
	slaveNum := 0
	for _, r := range results {
		if strings.Contains(r, "connected_slaves") {
			slave := strings.Split(r, ":")
			slaveNum, _ = strconv.Atoi(slave[1])
		}
	}

	if slaveNum+1 != r.RedisNum {
		r.err = errors.New(fmt.Sprintf("%d slave is offline", r.RedisNum-1-slaveNum))
	}

	return r
}

func (r *RedisSentinelHealth) Result() error {
	return r.err
}
