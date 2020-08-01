// @Title  服务发现插件的默认实现
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package discovey

import (
	"log"

	"github.com/elgong/elgongRPC/common"
	"github.com/elgong/elgongRPC/config"
	. "github.com/elgong/elgongRPC/plugin_centre"
)

func init() {
	defaultDiscovey := DefaultDiscovey{DiscoveyType, "defaultDiscovey", config.DefalutGlobalConfig.Services}
	PluginCenter.Register(defaultDiscovey.Typ, defaultDiscovey.Name, &defaultDiscovey)
}

type DefaultDiscovey struct {
	Typ         PluginType
	Name        PluginName
	services2Ip map[string][]string
}

func (d *DefaultDiscovey) Get(serviceName string) []string {

	if ret, ok := d.services2Ip[serviceName]; ok {
		return ret
	}
	// 返回空
	return []string{"127.0.0.1:8999"}
}

// ReportAndRemove 删除不能用的地址，有点暴力，未来在加重试吧
func (d *DefaultDiscovey) ReportAndRemove(serviceName string, delete string) {
	log.Println("find conn err, report and remove ... ")
	common.StringRemove(d.services2Ip[serviceName], delete)
}
