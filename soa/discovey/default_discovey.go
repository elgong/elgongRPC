// @Title  服务发现插件的默认实现
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package discovey

import (
	"github.com/elgong/elgongRPC/config"
	. "github.com/elgong/elgongRPC/plugin_centre"
)

func init() {
	defaultDiscovey := DefaultDiscovey{services2Ip: config.DefalutGlobalConfig.Services}
	PluginCenter.Register(defaultDiscovey.Typ, defaultDiscovey.Name, &defaultDiscovey)
}

type DefaultDiscovey struct {
	Name        PluginName
	Typ         PluginType
	services2Ip map[string][]string
}

func (d *DefaultDiscovey) Get(serviceName string) []string {

	if ret, ok := d.services2Ip[serviceName]; ok {
		return ret
	}
	// 返回空
	return []string{"127.0.0.1:8999"}
}
