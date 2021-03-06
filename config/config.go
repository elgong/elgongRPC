// @Title 全局的参数配置
// @Description  函数选项模式
// @Author  elgong 2020.7.24
// @Update  elgong 2020.7.24
package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// DefalutGlobalConfig 全局使用的配置参数
var DefalutGlobalConfig *Config

func init() {

	var err error
	DefalutGlobalConfig, err = Load("./config.yaml")
	fmt.Println(DefalutGlobalConfig)
	if err != nil {
		panic("参数解析异常")
	}
	fmt.Println("参数加载成功:", DefalutGlobalConfig.Name)
}

type Config struct {

	// 服务端和客户端公共的配值
	Name string `yaml:"name"`

	CodecPlugin    string `yaml:"codecPlugin"`
	ProtocolPlugin string `yaml:"protocolPlugin"`

	// 客户端的配置
	SelectorPlugin string `yaml:"selectorPlugin"`
	DiscoveyPlugin string `yaml:"discoveyPlugin"`
	ConnPlugin     string `yaml:"connPlugin"`

	Conn struct {
		MaxConn string `yaml:"maxConn"`
		MinConn string `yaml:"minConn"`

		InitialCap int `yaml:"initialCap"`
		MaxCap     int `yaml:"maxCap"`
		MaxIdle    int `yaml:"maxIdle"`

		TimeOut int `yaml:"timeout"`

		FailReconnect       bool `yaml:"failReconnect"`
		FailReconnectSecond int  `yaml:"failReconnectSecond"`
		FailReconnectTime   int  `yaml:"failReconnectTime"`
		IsTickerOpen        bool `yaml:"isTickerOpen"`
		TickerTime          int  `yaml:"tickerTime"`
	}

	// 客户端手动指定的 服务名 ：【地址】
	Services map[string][]string `yaml:"services"`

	// 服务端的配置
	Server struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}

	RegisterPlugin string `yaml:"registerPlugin"`
}

// Load 解析配置参数
func Load(path string) (*Config, error) {
	conf := new(Config)
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("yamlFile.Get err:", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("Unmarshal:", err)
		return nil, err
	}

	return conf, err
}
