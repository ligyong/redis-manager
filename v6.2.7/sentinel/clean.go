package sentinel

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Clean(addrs []string, password string) error {
	client := sentinelClients(redis.FailoverOptions{
		MasterName:    "",
		SentinelAddrs: addrs,
		Password:      password,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		fmt.Println("连接失败：", err)
	}

	_, err = client.FlushAll(context.TODO()).Result()
	if err != nil {
		return err
	}
	client.Close()
	/*
		master := c.getMasterNode()
		if master == nil {
			return errors.New("slave node disconnected")
		}
		masterInfo, err := master.Info(context.TODO(), "replication").Result()
		if err != nil {
			return errors.New("master node is offline")
		}

		masterInfoMap := pool.UnmarshalRedisConfig(masterInfo)
		slaveNum, err := strconv.Atoi(masterInfoMap["connected_slaves"])
		if err != nil {
			return errors.New(fmt.Sprintf("unknown error: %v", err))
		}
		if slaveNum+1 != len(c.client) {
			return errors.New("slave node disconnected")
		}
	*/
	return nil
}
