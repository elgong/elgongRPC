// @Title  服务发现 插件的 统一接口
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package discovey

const DiscoveyType = "discovey"

type Discovey interface {
	Get(serviceName string) []string
}
