package sentinel

import "github.com/ligyong/redis-manager/common"

type RedisBackUp struct {
}

func (b *RedisBackUp) Do() common.RedisOperatorResult {
	//send
	return b
}

func (b *RedisBackUp) Result() error {
	return nil
}
