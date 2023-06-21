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
	ConfigGet(options *RedisConfigGetOptions) RedisClientTransportMap
	ConfigSet(options *RedisConfigSetOptions) RedisClientTransport
	DataClean(options *RedisDataCleanOptions) RedisClientTransport
	Health(options *RedisHealthOptions) RedisClientTransport
}

type RedisClientTransport interface {
	Do() RedisOperatorResult
}

type RedisClientTransportMap interface {
	Do() RedisOperatorResultMap
}

type RedisOperatorResult interface {
	Result() error
}

type RedisOperatorResultMap interface {
	Result() (map[string]string, error)
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

type RedisConfigGetOptions struct {
	Node       []string
	MasterName string
	Password   string
}

type RedisConfigSetOptions struct {
	RedisNode []string
	Password  string
	Conf      map[string]string
}

type RedisDataCleanOptions struct {
	SentinelNode []string
	Password     string
	MasterName   string
}

type RedisHealthOptions struct {
	RedisNum     int
	SentinelNode []string
	Password     string
	MasterName   string
}
