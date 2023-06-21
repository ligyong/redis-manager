package redis_manager

//
//type RedisNodeClient interface {
//}
//
//func New() RedisNodeClient {
//	return RedisSentinelManagerClient{}
//}
//
//type RedisSentinelManagerClient struct {
//	Nodes map[string]transport.RedisClientTransport
//}
//
//func (r *RedisSentinelManagerClient) Install(node string) error {
//	client, ok := r.Nodes[node]
//	if !ok {
//		return errors.New("node is nil")
//	}
//
//	client.Do("Install")
//
//	return nil
//}
