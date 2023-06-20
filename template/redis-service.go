package template

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

var redisService = `[Unit]
Description=Redis
Documentation=https://redis.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
WorkingDirectory={{.WorkDir}}

ExecStart={{.RedisPath}} {{.RedisConf}}

# Let systemd restart this service always
Restart=always

# Specifies the maximum file descriptor number that can be opened by this process
LimitNOFILE=65536

# Specifies the maximum number of threads this process can create
TasksMax=infinity

# Disable timeout logic and wait until process is stopped
TimeoutStopSec=infinity
SendSIGKILL=yes

[Install]
WantedBy=multi-user.target`

type RedisServiceOptions struct {
	WorkDir   string
	RedisPath string
	RedisConf string
}

func RedisServiceCreate(serviceName string, options *RedisServiceOptions) error {
	//获取redis配置文件模板
	t, err := template.New("redis_config").Parse(redisService)
	if err != nil {
		return err
	}
	// 创建文件
	f, err := os.Create(path.Join("/etc/systemd/system/", fmt.Sprintf("%s.service", serviceName)))
	if err != nil {
		return err
	}
	defer f.Close()

	//将配置解析到模板中
	if err = t.Execute(f, options); err != nil {
		return err
	}

	return nil
}
