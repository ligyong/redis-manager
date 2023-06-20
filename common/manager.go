package common

type RedisManagerClient interface {
	GetClient(node string) RedisClient
}

type RedisManagerOptions struct {
	Node     []string
	Pattern  string
	Password string
}

type RedisClient interface {
	Install(options *RedisInstallOptions) RedisClientTransport
	Start(options *RedisStartOptions) RedisClientTransport
	Restart(options *RedisRestartOptions) RedisClientTransport
	Delete(options *RedisDeleteOptions) RedisClientTransport
	BackUp(options *RedisBackUpOptions) RedisClientTransport
}

type RedisClientTransport interface {
	Do() RedisOperatorResult
}

type RedisOperatorResult interface {
	Result() error
}

type RedisInstallOptions struct {
	MasterAddr     string                 `json:"masterAddr"` //master 地址
	MasterPort     string                 `json:"masterPort"`
	Role           string                 `json:"role"`
	Config         map[string]interface{} `json:"config"`
	WorkDir        string                 `json:"workDir"`
	SentinelPort   string                 `json:"sentinelPort"`
	SentinelQuorum int                    `json:"sentinelQuorum"`
	Password       string                 `json:"password"`
	ServiceName    string                 `json:"serviceName"`
}

type RedisStartOptions struct {
	ServiceName string `json:"serviceName"`
}

type RedisRestartOptions struct {
	ServiceName string `json:"serviceName"`
}

type RedisDeleteOptions struct {
	ServiceName string `json:"serviceName"`
	WorkDir     string `json:"workDir"`
}

type RedisBackUpOptions struct {
}
