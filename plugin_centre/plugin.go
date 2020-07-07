// 时间 2020.7.7
package plugin_centre

import (
	"fmt"
	"sync"
)

// 类型别名
type pluginType string
type pluginName string

// Plugins 插件中心结构体
type Plugins struct {
	name string
	pluginMap map[pluginType]map[pluginName]interface{}
	lock sync.RWMutex
}

// Register 注册插件到中心
func (p *Plugins) Register(pType pluginType, pName pluginName, plugin interface{}){

	p.lock.Lock()
	defer p.lock.Unlock()
	// 未注册过该插件类型
	if _, OK := p.pluginMap[pType]; !OK{
		p.pluginMap[pType] = make(map[pluginName]interface{})
	}

	// 注册插件
	p.pluginMap[pType][pName] = plugin
}

// Get 从中心获取指定插件
func (p *Plugins) Get(pType pluginType, pName pluginName) interface{}{
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
func (p *Plugins) Remove(pType pluginType, pName pluginName){
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

// 初始化插件
var plugins = Plugins{
	name: "插件管理",
	pluginMap:  make(map[pluginType]map[pluginName]interface{}),
}