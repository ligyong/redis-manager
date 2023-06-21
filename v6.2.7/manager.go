package v6_2_7

import (
	"github.com/ligyong/redis-manager/common"
	"github.com/ligyong/redis-manager/v6.2.7/sentinel"
)

type Redis struct {
	Node string
}

func (r Redis) Install(options *common.RedisInstallOptions) common.RedisClientTransport {
	return &sentinel.RedisInstall{
		Options: options,
		Node:    r.Node,
	}
}

func (r Redis) Start(options *common.RedisStartOptions) common.RedisClientTransport {
	return &sentinel.RedisStart{
		Options: options,
		Node:    r.Node,
	}
}

func (r Redis) Restart(options *common.RedisRestartOptions) common.RedisClientTransport {
	return &sentinel.RedisRestart{
		Options: options,
		Node:    r.Node,
	}
}

func (r Redis) Delete(options *common.RedisDeleteOptions) common.RedisClientTransport {
	return &sentinel.RedisDelete{
		Options: options,
		Node:    r.Node,
	}
}

func (Redis) BackUp(options *common.RedisBackUpOptions) common.RedisClientTransport {
	return &sentinel.RedisBackUp{}
}

func (Redis) ConfigGet(options *common.RedisConfigGetOptions) common.RedisClientTransportMap {
	return &sentinel.RedisConfigGet{
		Node:       options.Node,
		MasterName: options.MasterName,
		Password:   options.Password,
	}
}

func (Redis) ConfigSet(options *common.RedisConfigSetOptions) common.RedisClientTransport {
	return &sentinel.RedisConfigSet{
		RedisNode: options.RedisNode,
		Password:  options.Password,
		Conf:      options.Conf,
	}
}

func (Redis) DataClean(options *common.RedisDataCleanOptions) common.RedisClientTransport {
	return &sentinel.RedisDataClean{
		SentinelNode: options.SentinelNode,
		Password:     options.Password,
		MasterName:   options.MasterName,
	}
}

func (Redis) Health(options *common.RedisHealthOptions) common.RedisClientTransport {
	return &sentinel.RedisSentinelHealth{
		RedisNum:     options.RedisNum,
		SentinelNode: options.SentinelNode,
		Password:     options.Password,
		MasterName:   options.MasterName,
	}
}
