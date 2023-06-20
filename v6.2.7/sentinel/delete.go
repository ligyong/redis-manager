package sentinel

import (
	"fmt"
	"github.com/ligyong/redis-manager/common"
)

type RedisDelete struct {
	Options *common.RedisDeleteOptions
	Node    string
	err     error
}

func (r *RedisDelete) Do() common.RedisOperatorResult {
	err := common.RedisNodeSend(fmt.Sprintf("http://%s:%d/sentinel/%s", r.Node, 18080, "delete"), "", "", "", 0, r.Options)

	r.err = err
	return r
}

func (r *RedisDelete) Result() error {
	return nil
}
