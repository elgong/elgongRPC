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
	err := Load("./config/config.yaml", DefalutGlobalConfig)

	if err != nil {
		panic("参数解析异常")
	}

	fmt.Println("默认参数加载成功")
}

type Config struct {
	Name string `yaml:"name"`
	Conn struct {
		MaxConn string `yaml:"maxConn"`
		MinConn string `yaml:"minConn"`
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
func Load(path string, config *Config) error {
	conf := new(Config)
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("yamlFile.Get err:", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("Unmarshal:", err)
		return err
	}
	config = conf
	return err
}
