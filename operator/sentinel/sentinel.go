package sentinel

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ligyong/redis-manager/common"
	"github.com/ligyong/redis-manager/config"
	v6_2_6 "github.com/ligyong/redis-manager/config/v6.2.6"
	"github.com/ligyong/redis-manager/template"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func Route(g *gin.Engine) {
	r := g.Group("sentinel")
	r.POST("/install", Install)
	r.POST("/start", Start)
	r.POST("/restart", Restart)
	r.POST("/delete", Delete)
}

func Install(c *gin.Context) {
	fmt.Println("install operator")
	var options common.RedisInstallOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	if options.MasterAddr == "" {
		fmt.Println("master addr is nil")
		c.JSON(200, errors.New("master addr is nil"))
		return
	}

	err := createRedisWorkDir(options.WorkDir)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	err = createRedisConfig(&options)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	err = template.RedisServiceCreate(options.ServiceName, &template.RedisServiceOptions{
		WorkDir:   path.Join(options.WorkDir, "bin"),
		RedisPath: path.Join(options.WorkDir, "bin/redis-server"),
		RedisConf: path.Join(options.WorkDir, "/config/redis.conf"),
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	err = template.SentinelServiceCreate(options.ServiceName, &template.SentinelServiceOptions{
		SentinelPath: path.Join(options.WorkDir, "bin/redis-sentinel"),
		SentinelConf: path.Join(options.WorkDir, "/config/redis-sentinel.conf"),
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	c.JSON(200, nil)
}

func Start(c *gin.Context) {
	var options common.RedisStartOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err := common.RedisServiceDaemonReload()
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}
	_, err = common.RedisServiceStart(options.ServiceName)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err = common.RedisServiceStart(options.ServiceName + "-sentinel")
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	c.JSON(200, nil)
	return
}

func Restart(c *gin.Context) {
	var options common.RedisRestartOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err := common.RedisServiceRestart(options.ServiceName)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err = common.RedisServiceRestart(options.ServiceName + "-sentinel")
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}
}

func Delete(c *gin.Context) {
	fmt.Println("delete operator")
	var options common.RedisDeleteOptions
	if err := c.ShouldBindJSON(&options); err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err := common.RedisServiceStop(options.ServiceName)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	_, err = common.RedisServiceStop(options.ServiceName + "-sentinel")
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	os.Remove(path.Join("/etc/systemd/system/", fmt.Sprintf("%s.service", options.ServiceName)))
	os.Remove(path.Join("/etc/systemd/system/", fmt.Sprintf("%s-sentinel.service", options.ServiceName)))
	os.RemoveAll(options.WorkDir)
	c.JSON(200, nil)
}

func createRedisConfig(options *common.RedisInstallOptions) error {
	//处理config中，redis的相关路径
	if err := config.CreateRedisConfig("6.2.7", options.Config, path.Join(options.WorkDir, "/config/redis.conf")); err != nil {
		return err
	}

	//生成sentinel配置文件
	conf := fmt.Sprintf(v6_2_6.GetRedisSentinelConfigTemplate(), options.SentinelPort, options.MasterAddr, options.MasterPort, options.SentinelQuorum, options.Password)
	f, err := os.OpenFile(path.Join(options.WorkDir, "/config/redis-sentinel.conf"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

func createRedisWorkDir(dir string) error {
	// 创建目录
	err := os.MkdirAll(path.Join(dir, "config"), 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(dir, "data"), 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(dir, "log"), 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(dir, "backup"), 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Join(dir, "recovery"), 0755)
	if err != nil {
		return err
	}

	return nil
}
