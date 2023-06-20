package sentinel

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Health(addrs []string, password string) error {
	client := sentinelClients(redis.FailoverOptions{
		MasterName:    "",
		SentinelAddrs: addrs,
		Password:      password,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		fmt.Println("连接失败：", err)
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

func sentinelClients(options redis.FailoverOptions) *redis.Client {
	// 创建哨兵模式客户端
	return redis.NewFailoverClient(&options)

	// 测试连接是否正常
	//pong, err := client.Ping(client.Context()).Result()
	//if err != nil {
	//	fmt.Println("连接失败：", err)
	//} else {
	//	fmt.Println("连接成功：", pong)
	//}
	//
	//if err := client.Close(); err != nil {
	//	fmt.Println("关闭连接失败：", err)
	//}
}
