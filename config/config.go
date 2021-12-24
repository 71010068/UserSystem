package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"` // 数据库配置配置
}

// 解析yaml配置
func ParseYaml(configYaml *string, cfg **Config) {
	// 读取 yml 配置文件
	cfgYaml, err := ioutil.ReadFile(*configYaml)
	if err != nil {
		log.Fatal(map[string]error{"配置文件读取错误": err})
	}
	// 解析 yaml 配置到 config 中
	if err = yaml.Unmarshal(cfgYaml, cfg); err != nil {
		log.Fatal(map[string]error{"配置文件解析错误": err})
	}
}
