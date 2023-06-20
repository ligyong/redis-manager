package transport

/*
type RedisClientTransport interface {
	Do(string) error
}

func NewClient(node string) RedisClientTransport {
	return &HttpTransport{addrs: node}
}

type HttpTransport struct {
	addrs string
}

func (t *HttpTransport) Do(operator string) error {
	url := fmt.Sprintf("http://%s/redis/inner", t.addrs)
	c := http.Client{}
	c.Timeout = time.Minute
	_, err := c.Post(url, "application/json", bytes.NewReader(GetInnerBody(instanceID, version, operator, pattern, bodyByte)))
	if err != nil {
		return err
	}

	return nil
}

*/
