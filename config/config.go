// @Title  全局配置项
// @Author  elgong 2020.7.23
// @Update  elgong 2020.7.23
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ip      string `yaml:"IP"`
	service string `yaml:"name"`
	//service struct {
	//	ip string `yaml: "ip"`
	//}
}

func main() {
	content, _ := ioutil.ReadFile("./elgongRPC.yaml")
	env := &Config{}
	err := yaml.Unmarshal(content, env)

	fmt.Println(err, env)

}
