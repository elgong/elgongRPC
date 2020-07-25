// @Title 负载均衡插件  的统一接口
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package loadbalance

const SelectorType = "selector"

// 编解码插件  编解码统一接口
type Selector interface {

	// 返回 IP:PORT 数组
	Select([]string) string
}
