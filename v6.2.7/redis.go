package v6_2_7

/*
var redisOperatorMap map[string]common.RedisOperatorManager

func init() {
	go newInnerServer()
}

func newInnerServer() {
	r := gin.New()
	r.POST("/install", operator.Install)
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
}

func RedisInnerOperatorHandler(c *gin.Context) {
	var body common.InnerOperatorRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusOK, common.InnerOperatorResponse{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	switch body.Operator {
	case common.NodeStart:
		err = sentinel.InnerStart(body.Body.(sentinel.RedisInstallOptions))
	case common.SentinelStart:
		err = sentinel.InnerSentinelStart(body.Body.(sentinel.RedisSentinelStartOptions))
	}

	if err != nil {
		c.JSON(http.StatusOK, common.InnerOperatorResponse{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.InnerOperatorResponse{
		Code: 0,
	})
	return
}


*/
