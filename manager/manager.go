package manager

import (
	"github.com/ligyong/redis-manager/common"
	_ "github.com/ligyong/redis-manager/operator"
	v6_2_7 "github.com/ligyong/redis-manager/v6.2.7"
)

func NewRedisManagerClient(options *common.RedisManagerOptions) common.RedisManagerClient {
	m := redisManagerClient{clients: map[string]common.RedisClient{}}
	for _, node := range options.Node {
		m.clients[node] = v6_2_7.Redis{Node: node}
	}

	return m
}

type redisManagerClient struct {
	clients map[string]common.RedisClient
}

func (r redisManagerClient) GetClient(node string) common.RedisClient {
	client, ok := r.clients[node]
	if !ok {
		return nil
	}

	return client
}
