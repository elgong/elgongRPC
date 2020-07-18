// 时间 2020.7.7
package plugin_centre

import (
	"fmt"
	"sync"
)

// 别名
type PluginType string
type PluginName string

// 初始化插件
var PluginCenter = Plugin{
	name: "插件管理",
	pluginMap:  make(map[PluginType]map[PluginName]interface{}),
}

// Plugins 插件中心结构体
type Plugin struct {
	name string
	pluginMap map[PluginType]map[PluginName]interface{}
	lock sync.RWMutex
}

// Register 注册插件到中心
func (p *Plugin) Register(pType PluginType, pName PluginName, plugin interface{}){

	p.lock.Lock()
	defer p.lock.Unlock()
	// 未注册过该插件类型
	if _, OK := p.pluginMap[pType]; !OK{
		p.pluginMap[pType] = make(map[PluginName]interface{})
	}

	// 注册插件
	p.pluginMap[pType][pName] = plugin
}

// Get 从中心获取指定插件
func (p *Plugin) Get(pType PluginType, pName PluginName) interface{}{
	p.lock.RLock()
	defer p.lock.RUnlock()

	if _, OK := p.pluginMap[pType]; !OK{
		panic("该插件类型未注册")
		return nil
	}

	// 找到插件
	if _, OK := p.pluginMap[pType][pName]; !OK{
		panic("该插件名未注册")
		return nil
	}

	return p.pluginMap[pType][pName]
}

// Remove 移除插件
func (p *Plugin) Remove(pType PluginType, pName PluginName){
	p.lock.RLock()
	defer p.lock.RUnlock()

	if _, OK := p.pluginMap[pType]; !OK{
		panic("该插件类型未注册")
		return
	}

	// 找到插件
	if _, OK := p.pluginMap[pType][pName]; !OK{
		panic("该插件名未注册")
		return
	}

	delete(p.pluginMap[pType], pName)

	fmt.Println("插件%s 已经从管理中心移除", pName)
}


