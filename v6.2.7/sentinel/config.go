package sentinel

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func ConfigSet(addrs []string, password string, conf map[string]interface{}) error {
	client := sentinelClients(redis.FailoverOptions{
		MasterName:    "",
		SentinelAddrs: addrs,
		Password:      password,
	})
	/*
		f := func(ctx context.Context, client *redis.Client) error {
			pipe := client.TxPipeline()

			for k, v := range conf {
				pipe.ConfigSet(ctx, k, fmt.Sprintf("%v", v))
			}

			pipe.ConfigRewrite(ctx)

			if _, err := pipe.Exec(ctx); err != nil {
				return err
			}

			return nil
		}


	*/
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	//client.ConfigSet(context.TODO(), "", "")

	return nil
}
