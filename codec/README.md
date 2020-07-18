# codec 包说明

> 这个包是编解码器的定义和实现。
>
> 用户可以自定义实现自己的codec，并通过Register 注册进插件管理中心。但必须遵守以下规则：

## v1.0 插件开发的规范：
插件定义时：
1. 同种类型的插件要实现同一个接口; // 解决取出时的类型断言，这样统一起来就容易使用啦；
2. 插件在实现接口时，附带两个变量，插件类型+插件名；

插件使用时：
3. 使用类型定义了的接口来断言； // 插件中心为了通用使用了 interface{}, 使用公共的就统一啦；


## V1.0 已提供的codec插件类型

- gob
- msgpack

使用时，从注册中心可以获取到对应的插件：

```go
// gob 获取
PluginCenter.Get(CodecType, "gobCodec").(Codec).Encode(data)

// msgpack 获取
PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Encode(data)
```