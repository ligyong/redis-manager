package config

import (
	"encoding/json"
	"errors"
	"fmt"
	v6_2_6 "github.com/ligyong/redis-manager/config/v6.2.6"
	"os"
	"text/template"
)

// 根据模板和数据生成redis的配置文件
func CreateRedisConfig(version string, configMap map[string]interface{}, filePath string) error {
	configJson, err := json.Marshal(configMap)
	if err != nil {
		return err
	}

	switch version {
	case "6.2.6", "6.2.7", "6.2.8":
		return v6_2_6.CreateRedisConfig(string(configJson), filePath)
	}
	return nil
}

func redisConfigVerification(filePath string) error {
	return nil
}

// todo 端口
func getConfigDefault(version, conf string) (interface{}, error) {
	switch version {
	case "6.2.6", "6.2.7", "6.2.8":
		return v6_2_6.NewDefaultConfigStruct(conf)
	}

	return "", errors.New(fmt.Sprintf("redis version %s,is not support", version))
}

func getConfigTemplateByVersion(version string) string {
	switch version {
	case "6.2.6", "6.2.7", "6.2.8":
		return v6_2_6.GetRedisConfigTemplate()
	}

	return ""
}

// 根据模板和数据生成redis的配置文件
func CreateRedisSentinelConfig(version string, configMap map[string]interface{}, filePath string) error {
	configJson, err := json.Marshal(configMap)
	if err != nil {
		return err
	}

	//根据传入的配置生成相关结构，并初步校验数据有效性
	configDefault, err := getSentinelConfigDefault(version, string(configJson))
	if err != nil {
		return err
	}
	//获取redis配置文件模板
	redisConfigTemplate := getSentinelConfigTemplateByVersion(version)
	t, err := template.New("redis_config").Parse(redisConfigTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	//将配置解析到模板中
	if err = t.Execute(f, configDefault); err != nil {
		return err
	}
	//对配置文件进行校验
	if err := redisConfigVerification(filePath); err != nil {
		return err
	}

	return nil
}

func getSentinelConfigTemplateByVersion(version string) string {
	switch version {
	case "6.2.6", "6.2.7", "6.2.8":
		return v6_2_6.GetRedisSentinelConfigTemplate()
	}

	return ""
}

func getSentinelConfigDefault(version, conf string) (interface{}, error) {
	switch version {
	case "6.2.6", "6.2.7", "6.2.8":
		return v6_2_6.NewDefaultConfigStruct(conf)
	}

	return "", errors.New(fmt.Sprintf("redis version %s,is not support", version))
}

func GetSentinelConfigStruct(version, conf string) (interface{}, error) {
	return getConfigDefault(version, conf)
}

func GetSentinelConfigMap(version string, configMap map[string]interface{}) (map[string]interface{}, error) {
	conf, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}

	config, err := getConfigDefault(version, string(conf))
	if err != nil {
		return nil, err
	}

	configJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	newConfigMap := make(map[string]interface{})
	err = json.Unmarshal(configJson, &newConfigMap)
	if err != nil {
		return nil, err
	}

	return newConfigMap, nil
}

func GetSentinelConfigJson(version, conf string) (string, error) {
	config, err := getConfigDefault(version, conf)
	if err != nil {
		return "", err
	}

	configJson, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(configJson), nil
}
