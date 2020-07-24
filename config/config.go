// @Title 全局的参数配置
// @Description  函数选项模式
// @Author  elgong 2020.7.24
// @Update  elgong 2020.7.24
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// DefalutGlobalConfig 全局使用的配置参数
var DefalutGlobalConfig *Config

func init() {
	str, _ := os.Getwd()
	fmt.Println(str)
	var err error
	DefalutGlobalConfig, err = Load("../config.yaml")

	if err != nil {
		panic("参数解析异常")
	}

	fmt.Println("参数加载成功:", DefalutGlobalConfig.Name)
}

type Config struct {
	Name string `yaml:"name"`
	Conn struct {
		MaxConn string `yaml:"maxConn"`
		MinConn string `yaml:"minConn"`

		InitialCap int `yaml:"initialCap"`
		MaxCap     int `yaml:"maxCap"`
		MaxIdle    int `yaml:"maxIdle"`
		//idletime:  1,
		//maxLifetime: 2,
		FailReconnect       bool `yaml:"failReconnect"`
		FailReconnectSecond int  `yaml:"failReconnectSecond"`
		FailReconnectTime   int  `yaml:"failReconnectTime"`
		IsTickerOpen        bool `yaml:"isTickerOpen"`
		TickerTime          int  `yaml:"tickerTime"`
	}

	Codec struct {
		Codec string `yaml:"codec-method"`
	}

	Server struct {
		Servicename   string `yaml:"servicename"`
		Servicemethod string `yaml:"servicemethod"`
		Ip            string `yaml:"ip"`
	}
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
