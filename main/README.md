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
	
#### 特别注意的点
1. 编码的数据中有空接口类型，传递时赋值的空接口为：基本类型（int、float、string）、切片时，可以不进行注册。
2. 编码的数据中有空接口类型，传递时赋值的空接口为：map、struct时，必须进行注册。

```go
gob.Register(map[int]string{}) //TODO：进行了注册
gob.Register(Msg{}) //TODO：进行了注册
```

## net包的基本使用
