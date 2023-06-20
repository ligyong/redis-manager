package sentinel

import (
	"fmt"
	"github.com/ligyong/redis-manager/common"
)

type RedisInstall struct {
	Options *common.RedisInstallOptions
	Node    string
	err     error
}

func (r *RedisInstall) Do() common.RedisOperatorResult {
	err := common.RedisNodeSend(fmt.Sprintf("http://%s:%d/sentinel/%s", r.Node, 18080, "install"), "", "", "", 0, r.Options)

	r.err = err
	return r
}

func (r *RedisInstall) Result() error {
	return r.err
}

/*
type RedisSentinelInstallOptions struct {
	Nodes      []string
	MasterAddr string
	Workdir    string
	Config     map[string]interface{}
	Password   string
	Role       string

	SentinelPort string
	MasterPort   string
	Quorum       int
}

func Install(options *RedisSentinelInstallOptions) error {
	if options.MasterAddr == "" {
		return errors.New("master addr is nil")
	}
	createRedisConfig(options)
	return nil
}

func createRedisService(options *RedisSentinelInstallOptions) error {
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

func createRedisConfig(options *RedisSentinelInstallOptions) error {
	//生成redis配置文件
	if options.Role != "master" {
		options.Config["slaveof"] = options.MasterAddr
	}

	//处理config中，redis的相关路径
	if err := config.CreateRedisConfig("6.2.7", options.Config, path.Join(options.Workdir, "/config/redis.conf")); err != nil {
		return err
	}

	//生成sentinel配置文件
	conf := fmt.Sprintf(v6_2_6.GetRedisSentinelConfigTemplate(), options.SentinelPort, options.MasterAddr, options.MasterPort, options.Quorum, options.Password)
	f, err := os.OpenFile(path.Join(options.Workdir, "/config/redis-sentinel.conf"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Errorf("更新redis配置文件:%s 失败:%s", "/etc/redis/redis-sentinel.conf", err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(conf)
	if err != nil {
		log.Warnf("写入redis配置到文件:%s失败:%s", "/etc/redis/redis-sentinel.conf", err.Error())
	}

	return nil
}
*/
