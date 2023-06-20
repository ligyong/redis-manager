package sentinel

import (
	"fmt"
	"github.com/ligyong/redis-manager/common"
)

type RedisRestart struct {
	Options *common.RedisRestartOptions
	Node    string
	err     error
}

func (r *RedisRestart) Do() common.RedisOperatorResult {
	err := common.RedisNodeSend(fmt.Sprintf("http://%s:%d/sentinel/%s", r.Node, 18080, "restart"), "", "", "", 0, r.Options)

	r.err = err
	return r
}

func (r *RedisRestart) Result() error {
	return r.err
}
