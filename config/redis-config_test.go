package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCreateRedisConfig(t *testing.T) {
	configMap := make(map[string]interface{})
	//if err := json.Unmarshal([]byte(c.conf.Config), &configMap); err != nil {
	//	log.Error(err)
	//	return
	//}

	configMap["requirepass"] = "sobey1234" //设置密码
	configMap["cluster-enabled"] = "no"    // 关闭集群模式
	configJson, err := json.Marshal(configMap)
	if err != nil {
		log.Error(err)
		return
	}

	if err := CreateRedisConfig("v6.2.6", string(configJson), "/root/agent/pkg/redis/config/redis.conf"); err != nil {
		log.Error(err)
		return
	}
}
