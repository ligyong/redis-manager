package common

import (
	"context"
	"gorm.io/datatypes"
)

var (
	RedisBinPath      = "/bin"
	RedisLogPath      = "/log"
	RedisBackupPath   = "/backup"
	RedisDataPath     = "/data"
	RedisRecoveryPath = "/recovery"
	RedisConfigPath   = "/config"
)

const (
	RedisBackupDefaultDir     = "" //redis默认备份目录
	RedisBackUpSuccessState   = "" //redis备份成功
	RedisBackUpFailState      = "" //redis备份失败
	RedisBackUpRunningState   = "" //redis备份中
	RedisRecoverySuccessState = "" //redis备份恢复成功
	RedisRecoveryFailState    = "" //redis备份恢复失败功
	RedisRecoveryRunningState = "" //redis备份恢复中
)

const (
	SINGLE      = "single"
	REPLICATION = "replication"
	SENTINEL    = "sentinel"
	CLUSTER     = "cluster"
)

const (
	RedisconfigRequirepass = "requirepass"
	RedisconfigMasterauth  = "masterauth"
	RedisconfigPort        = "port"
	RedisconfigSlaveOf     = "slaveof"
)

const (
	RedisConfFileDefaultPath = "/etc/redis/redis.conf"
)

/*******************************************************/
type RedisInnerOperator int

const (
	NodeStart RedisInnerOperator = iota
	NodeRestart
	SentinelStart
	SentinelPasswd
)

type RedisNodeInformation struct {
	InstanceID string            `json:"instanceID" gorm:"column:instanceID"` //实例ID
	Ip         string            `json:"ip"`                                  //节点ip
	Port       string            `json:"port"`                                //redis端口
	Master     string            `json:"master"`                              //master IP
	Config     datatypes.JSONMap `json:"config"`                              //redis配置
	Mode       string            `json:"mode"`                                //redis模式
	State      int               `json:"state"`                               //redis状态
	Password   string            `json:"password"`                            //redis密码
	Version    string            `json:"version"`                             //redis版本
	Role       string            `json:"role"`                                //redis角色
}

type RedisInstallConfig struct {
	Pattern         string                 //redis部署模式
	InstanceID      string                 //实例ID
	Passwd          string                 //redis密码
	NodeInformation []RedisNodeInformation //节点详细信息
	LocalIP         string                 //本地地址
}

type RedisAppInstallRequest struct {
	Nodes      []RedisNode            `json:"nodes"`      //redis节点
	Conf       map[string]interface{} `json:"conf"`       //redis配置
	Password   string                 `json:"password"`   //redis密码
	Pattern    string                 `json:"pattern"`    //redis部署模式
	Version    string                 `json:"version"`    //redis版本
	InstanceID string                 `json:"instanceID"` //实例ID
	LocalIP    string                 `json:"localIP"`
}

type RedisNode struct {
	ServerIP string `json:"serverIp"`
}

type RedisOperatorManager interface {
	New(*RedisInstallConfig) RedisOperatorClient
	NewInner(*RedisInstallConfig) RedisOperatorClient
}

type RedisOperatorClient interface {
	// Install 应用安装
	Install() error
	// ReStart 重启
	ReStart() error
	// Delete 卸载
	Delete() error
	// BackUp 数据备份
	BackUp(string, string) (BackupResult, error)
	// Recovery 数据恢复
	Recovery(uint, string) error
	// UpdatePassword 更新redis密码
	UpdatePassword(string) error
	// UpdateConfig 修改redis配置
	UpdateConfig(ctx context.Context, conf map[string]interface{}) error
	// GetConfig 获取配置
	GetConfig() []RedisNodeInformation
	// Health 健康检测
	Health() error
	// CacheClean 清除缓存
	CacheClean() error
	// InnerHandler 内部操作
	InnerHandler(InnerOperatorRequest) InnerOperatorResponse
	// UpdateRedisState 更新redis状态
	UpdateRedisState(int)
	/*
		BackupList(body common.ReqBody) sobey_err.Response
			RecoveryList(body common.ReqBody) sobey_err.Response
			BackupDelete(body common.ReqBody) sobey_err.Response
	*/
}

type BackupResult struct {
	State    string
	FileName string
	FileSize int64
}

type RedisConfigData interface {
	Get(string) (*RedisInstallConfig, error)
}

type InnerOperatorRequest struct {
	InstanceID string             `json:"instanceID"`
	Version    string             `json:"version"`
	Operator   RedisInnerOperator `json:"operator"`
	Pattern    string             `json:"pattern"` //redis部署模式
	Body       interface{}        `json:"body"`
}

type InnerOperatorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
