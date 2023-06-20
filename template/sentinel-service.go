package template

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

var sentinelService = `[Unit]
Description=Redis Sentinel
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
ExecStart={{.SentinelPath}} {{.SentinelConf}}
Restart=always

[Install]
WantedBy=multi-user.target`

type SentinelServiceOptions struct {
	SentinelPath string
	SentinelConf string
}

func SentinelServiceCreate(serviceName string, options *SentinelServiceOptions) error {
	//获取redis配置文件模板
	t, err := template.New("redis_config").Parse(sentinelService)
	if err != nil {
		return err
	}
	// 创建文件
	f, err := os.Create(path.Join("/etc/systemd/system/", fmt.Sprintf("%s-sentinel.service", serviceName)))
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
