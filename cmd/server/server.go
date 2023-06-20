package main

import (
	"github.com/ligyong/redis-manager/common"
	"github.com/ligyong/redis-manager/manager"
)

func main() {
	manager.NewRedisManagerClient(&common.RedisManagerOptions{
		Node:     []string{"172.16.200.30"},
		Pattern:  "sentinel",
		Password: "",
	})

	select {}
}
