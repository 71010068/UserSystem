// 程序入口
package main

import (
	"flag"
	"fmt"

	"github.com/71010068/UserSystem/config"
)

var (
	configYaml = flag.String("c", "config/config.yaml", "配置文件")
	cfg        *config.Config // yaml 相关 config
)

func init() {
	// 解析 yaml 配置文件
	config.ParseYaml(configYaml, &cfg)
}

func main() {
	fmt.Println(cfg)
}
