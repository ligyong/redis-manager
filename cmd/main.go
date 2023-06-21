package main

import (
	"fmt"
	"github.com/ligyong/redis-manager/common"
	"github.com/ligyong/redis-manager/config"
	"github.com/ligyong/redis-manager/manager"
	"path"
)

func main() {

	m := manager.NewRedisManagerClient(&common.RedisManagerOptions{
		Node:     []string{"172.16.200.30", "172.16.200.31", "172.16.200.32"},
		Pattern:  "sentinel",
		Password: "",
	})

	err := Install(m)
	if err != nil {
		return
	}
	err = Start(m)
	if err != nil {
		return
	}

	err = Restart(m)
	if err != nil {
		return
	}

	//err := Delete(m)
	//if err != nil {
	//	return
	//}
	return
}

func Install(m common.RedisManagerClient) error {
	configMap := make(map[string]interface{})
	configMap[common.RedisconfigRequirepass] = "123456"
	configMap[common.RedisconfigMasterauth] = "123456"

	configMap, err := config.GetSentinelConfigMap("6.2.7", configMap)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//todo 对配置文件的修改放在上层进行操作
	configMap["logfile"] = path.Join("/sobey/moudle/redis/1", configMap["logfile"].(string))
	configMap["dir"] = path.Join("/sobey/moudle/redis/1", configMap["dir"].(string))
	//master节点
	err = m.GetClient("172.16.200.30").Install(&common.RedisInstallOptions{
		MasterAddr:     "172.16.200.30",
		MasterPort:     "6379",
		Role:           "master",
		Config:         configMap,
		WorkDir:        "/sobey/moudle/redis/1",
		SentinelPort:   "26379",
		SentinelQuorum: 2,
		Password:       "123456",
		ServiceName:    "redis-1",
	}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//slave节点
	slaveConfigMap := configMap
	slaveConfigMap["slaveof"] = "172.16.200.30 6379"
	err = m.GetClient("172.16.200.31").Install(&common.RedisInstallOptions{
		MasterAddr:     "172.16.200.30",
		MasterPort:     "6379",
		Role:           "slave",
		Config:         configMap,
		WorkDir:        "/sobey/moudle/redis/1",
		SentinelPort:   "26379",
		SentinelQuorum: 2,
		Password:       "123456",
		ServiceName:    "redis-1",
	}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.32").Install(&common.RedisInstallOptions{
		MasterAddr:     "172.16.200.30",
		MasterPort:     "6379",
		Role:           "slave",
		Config:         configMap,
		WorkDir:        "/sobey/moudle/redis/1",
		SentinelPort:   "26379",
		SentinelQuorum: 2,
		Password:       "123456",
		ServiceName:    "redis-1",
	}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Start(m common.RedisManagerClient) error {
	err := m.GetClient("172.16.200.30").Start(&common.RedisStartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.31").Start(&common.RedisStartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.32").Start(&common.RedisStartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Restart(m common.RedisManagerClient) error {
	err := m.GetClient("172.16.200.30").Restart(&common.RedisRestartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.31").Restart(&common.RedisRestartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.32").Restart(&common.RedisRestartOptions{ServiceName: "redis-1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

	return nil
}

func Delete(m common.RedisManagerClient) error {
	err := m.GetClient("172.16.200.30").Delete(&common.RedisDeleteOptions{ServiceName: "redis-1", WorkDir: "/sobey/moudle/redis/1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.31").Delete(&common.RedisDeleteOptions{ServiceName: "redis-1", WorkDir: "/sobey/moudle/redis/1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = m.GetClient("172.16.200.32").Delete(&common.RedisDeleteOptions{ServiceName: "redis-1", WorkDir: "/sobey/moudle/redis/1"}).Do().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
