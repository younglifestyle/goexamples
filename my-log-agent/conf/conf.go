package conf

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"goexamples/my-log-agent/httpstream"
	"sync"
)

var (
	// Conf conf
	Conf = &Config{}
	lock = new(sync.RWMutex)
)

type Config struct {
	// httpstream
	HttpAddr httpstream.Config `json:"httpstream"`
}

// 以后改为etcd获取配置
func ParseConfig(cfg string) error {
	var c Config

	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("config file %s is nonexistent", cfg)
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read config file %s fail %s", cfg, err)
	}

	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse config file %s fail %s", cfg, err)
	}

	lock.Lock()
	defer lock.Unlock()

	Conf = &c

	return nil
}
