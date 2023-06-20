package operator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ligyong/redis-manager/operator/sentinel"
	"net"
	"net/http"
	"time"
)

func init() {
	go func() {
		r := gin.New()
		sentinel.Route(r)
		addr := fmt.Sprintf(":%d", 18080)
		server := &http.Server{Addr: addr, Handler: r, ReadHeaderTimeout: time.Minute}
		ln, err := net.Listen("tcp4", addr)
		if err != nil {
			fmt.Printf("启动服务失败：%s", err.Error())
			panic(err)
		}
		// 只监听ipv4
		if s, ok := ln.(*net.TCPListener); !ok {
			panic(fmt.Sprintf("启动server服务失败:%v", ok))
		} else {
			if err := server.Serve(s); err != nil {
				fmt.Printf("启动服务失败：%s", err.Error())
			}
		}
	}()
}
