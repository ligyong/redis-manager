package sentinel

import (
	"fmt"
	"github.com/ligyong/redis-manager/common"
)

type RedisStart struct {
	Options *common.RedisStartOptions
	Node    string
	err     error
}

func (r *RedisStart) Do() common.RedisOperatorResult {
	err := common.RedisNodeSend(fmt.Sprintf("http://%s:%d/sentinel/%s", r.Node, 18080, "start"), "", "", "", 0, r.Options)

	r.err = err
	return r
}

func (r *RedisStart) Result() error {
	return r.err
}

/*
type RedisInstallOptions struct {
	Node       []string
	MasterAddr string
	Workdir    string
	Config     map[string]interface{}
	Version    string
	Password   string
	Role       string
}

func RedisStart(options RedisInstallOptions) error {
	//todo 配置修改
	//todo 选主、启动redis
	for i, node := range options.Node {
		body := options
		if i == 0 {
			body.Role = "master"
			body.MasterAddr = node
		}
		if err := common.RedisNodeSend(node, "", options.Version, "sentinel", 0, body); err != nil {
			return err
		}
	}

	return nil
}

type RedisSentinelStartOptions struct {
	Node         []string
	SentinelPort string
	MasterAddr   string
	MasterPort   string
	Quorum       int
	Password     string
	Version      string
}

func RedisSentinelStart(options RedisSentinelStartOptions) error {
	//todo 启动redis-sentinel
	err := common.RedisNodesSend(options.Node, "", options.Version, "sentinel", 1, options) //todo带上body "sentinelStart"
	if err != nil {
		return err
	}
	return nil
}

func InnerStart(options RedisInstallOptions) error {
	if options.Role != "master" {
		options.Config["slaveof"] = options.MasterAddr
	}

	//处理config中，redis的相关路径
	if err := config.CreateRedisConfig(options.Version, options.Config, options.Workdir); err != nil {
		return err
	}

	_, err := common.RedisServiceDaemonReload()
	if err != nil {
		return err
	}
	_, err = common.RedisServiceStart("redis")
	if err != nil {
		return err
	}

	return nil
}

type redisInnerSentinelStart struct {
	SentinelPort string
	MasterAddr   string
	MasterPort   string
	Quorum       int
	Password     string
}

func InnerSentinelStart(options RedisSentinelStartOptions) error {
	conf := fmt.Sprintf(v6_2_6.GetRedisSentinelConfigTemplate(), options.SentinelPort, options.MasterAddr, options.MasterPort, options.Quorum, options.Password)
	f, err := os.OpenFile("/etc/redis/redis-sentinel.conf", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Errorf("更新redis配置文件:%s 失败:%s", "/etc/redis/redis-sentinel.conf", err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(conf)
	if err != nil {
		log.Warnf("写入redis配置到文件:%s失败:%s", "/etc/redis/redis-sentinel.conf", err.Error())
	}

	_, err = common.RedisServiceDaemonReload()
	if err != nil {
		return err
	}
	_, err = common.RedisServiceStart("redis-sentinel")
	if err != nil {
		return err
	}

	return nil
}

*/
