package producer

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"sync"
)

type ProjectConfig struct {
	lock sync.RWMutex
	client   *sls.Client
	producerLogGroupSize int64
	ioLock sync.RWMutex
}


func initProjectConfig(config *ProducerConfig) *ProjectConfig {
	client := &sls.Client{
		Endpoint:        config.Endpoint,
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
		SecurityToken:   config.SecurityToken,
	}
	return &ProjectConfig{
		client: client,
	}
}

func (projectConfig *ProjectConfig) PutNewProjectConfig(config *ProducerConfig) {
	defer projectConfig.lock.Unlock()
	projectConfig.lock.Lock()
	projectConfig.client = &sls.Client{
		Endpoint:        config.Endpoint,
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
		SecurityToken:   config.SecurityToken,
	}
}

func (ProjectConfig *ProjectConfig) getClient() *sls.Client {
	defer ProjectConfig.lock.RUnlock()
	ProjectConfig.lock.RLock()
	return ProjectConfig.client
}
