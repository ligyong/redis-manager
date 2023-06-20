package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

type StorageManager interface {
	Upload(filePath, fileName string) error
	Download(filePath, fileName string) error
}

var storage StorageManager = defaultStorage{}

func InitRedisStorage(h StorageManager) {
	storage = h
}

type defaultStorage struct {
}

func (defaultStorage) Upload(filePath, fileName string) error {
	return nil
}

func (defaultStorage) Download(filePath, fileName string) error {
	return nil
}

func GetInnerBody(instanceID, version string, operator RedisInnerOperator, pattern string, body interface{}) []byte {
	innerBody := InnerOperatorRequest{
		InstanceID: instanceID,
		Version:    version,
		Operator:   operator,
		Pattern:    pattern,
		Body:       body,
	}

	bodyByte, err := json.Marshal(innerBody)
	if err != nil {
		return nil
	}

	return bodyByte
}

func RedisNodesSend(addrs []string, instanceID, version, pattern string, operator RedisInnerOperator, body interface{}) error {
	for _, addr := range addrs {
		url := fmt.Sprintf("http://%s/redis/inner", addr)
		c := http.Client{}
		c.Timeout = time.Minute
		_, err := c.Post(url, "application/json", bytes.NewReader(GetInnerBody(instanceID, version, operator, pattern, body)))
		if err != nil {
			return err
		}
	}

	return nil
}

func RedisNodeSend(addr string, instanceID, version, pattern string, operator RedisInnerOperator, data interface{}) error {
	bodyByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c := http.Client{}
	c.Timeout = time.Minute
	_, err = c.Post(addr, "application/json", bytes.NewReader(bodyByte))
	if err != nil {
		return err
	}

	return nil
}

func RedisServiceStart(service string) ([]byte, error) {
	out, err := exec.Command("/usr/bin/bash", "-c", fmt.Sprintf("systemctl start %s", service)).CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}

func RedisServiceStop(service string) ([]byte, error) {
	out, err := exec.Command("/usr/bin/bash", "-c", fmt.Sprintf("systemctl stop %s", service)).CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}

func RedisServiceRestart(service string) ([]byte, error) {
	out, err := RedisServiceStop(service)
	if err != nil {
		return out, err
	}

	time.Sleep(1 * time.Second)

	out, err = RedisServiceStart(service)
	if err != nil {
		return out, err
	}

	return out, nil
}

func RedisServiceDaemonReload() ([]byte, error) {
	out, err := exec.Command("/usr/bin/bash", "-c", "systemctl daemon-reload").CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}

func RedisExec(addr string, instanceID, version, pattern string, operator RedisInnerOperator, data interface{}) error {
	bodyByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s/redis/inner", addr)
	c := http.Client{}
	c.Timeout = time.Minute
	_, err = c.Post(url, "application/json", bytes.NewReader(GetInnerBody(instanceID, version, operator, pattern, bodyByte)))
	if err != nil {
		return err
	}

	return nil
}
