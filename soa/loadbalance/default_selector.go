// @Title  负载均衡插件的默认实现
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package loadbalance

import (
	"math/rand"

	. "github.com/elgong/elgongRPC/plugin_centre"
)

// GobCodec 注册进插件管理中心
func init() {
	defaultSelector := DefaultSelector{SelectorType, "defaultSelector"}
	PluginCenter.Register(defaultSelector.Type, defaultSelector.Name, &defaultSelector)
}

// DefaultSelector 默认服务发现
type DefaultSelector struct {
	Type PluginType
	Name PluginName
}

// Select 负载均衡，从[]中随机获取其中一个
func (s *DefaultSelector) Select(serviceList []string) string {
	randNum := rand.Intn(len(serviceList))
	return serviceList[randNum]
}
