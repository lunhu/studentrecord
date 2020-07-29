package common

import (
	"sync"

	"github.com/robfig/config"
)

const (
	profile = "config.ini"
)

var (
	//Config 配置文件工具
	configParser *config.Config
	configOnce sync.Once
)

func GetConfig() *config.Config {
	configOnce.Do(func(){
		var err error
		configParser, err = config.ReadDefault(profile)
		if err != nil {
			panic("获取配置文件 config.ini 失败: "+err.Error())
		}
	})
	return configParser
}



