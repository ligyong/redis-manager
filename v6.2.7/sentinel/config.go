package sentinel

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ligyong/redis-manager/common"
)

type RedisConfigSet struct {
	RedisNode []string
	Password  string
	Conf      map[string]string
	err       error
}

func (r *RedisConfigSet) Do() common.RedisOperatorResult {
	f := func(ctx context.Context, client *redis.Client) error {
		pipe := client.TxPipeline()

		for k, v := range r.Conf {
			pipe.ConfigSet(ctx, k, fmt.Sprintf("%v", v))
		}

		pipe.ConfigRewrite(ctx)

		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}

		return nil
	}
	for _, node := range r.RedisNode {
		cli := redis.NewClient(
			&redis.Options{
				Addr:     node,
				Password: r.Password,
			},
		)
		err := f(context.TODO(), cli)
		if err != nil {
			r.err = err
			return r
		}
		cli.Close()
	}

	return r
}

func (r *RedisConfigSet) Result() error {
	return r.err
}
