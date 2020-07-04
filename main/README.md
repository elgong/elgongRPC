# main 包功能说明

> 这个包是学习使用各个组件功能的，不属于这个工程范围，未来可移除


## rpc 
> go 自带的net.rpc 使用



## codec

`go_gob.go` 

> 测试go 自带的gob 编解码功能

### gob 使用步骤
#### 编码
- enc := gob.NewEncoder(&buf)  // 创建一个编码器，并指定缓存位置
- err := enc.Encode(&sendMsg)  // 把具体数据编码后放入缓冲区

#### 解码
- enc := gob.NewEncoder(&buf)   // 创建一个解码器，并指定缓存位置
- err := enc.Encode(&msg)       // 把序列解码为 msg
	

