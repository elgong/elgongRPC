# elgongRPC 框架实践
**项目简述**：尝试动手实现基于golang 的 rpc 微服务框架

**项目起始于**：2020.7.8

**目前状态**：更新中...

**项目地址**： [点击可跳转]( https://github.com/users/elgong/projects/1 )


[TOC]

## elgongRPC 框架结构

![框架](https://github.com/elgong/elgongRPC/blob/master/rpc.png)



## elgongRPC  V1.0 特点

- **分层设计**: 尽可能的分层设计，每个层避免出现循环依赖，但可能出现单向的依赖

  - **特殊情况**：比如 `protocol 层`依赖于` codec层`，但是最终调用使用的是`protocol层`对外提供接口， 用户感知不到。 

- **组件尽可能插件化**：由统一的插件管理中心来管理注册和使用

  - **实现思路**：同种类型插件实现类型共同的接口，把自己注册到插件管理中心的map容器中， 调用时使用接口来统一断言调用；
  - 支持自定义**服务发现插件**：可以通过实现`Discovey`接口来自定义
    - 已实现基于配置文件的服务发现插件 [code](https://github.com/elgong/elgongRPC/blob/master/soa/discovey/default_discovey.go)
    - 已实现基于Redis 注册中心的服务发现插件 [code](https://github.com/elgong/elgongRPC/tree/master/soa/register/redisPlugin)
  - 支持自定义**服务注册插件**：可以通过实现`Register`接口来自定义
    - 已实现基于配置文件的服务注册插件   [code](https://github.com/elgong/elgongRPC/tree/master/soa/register/redisPlugin)
    - 已实现基于Redis 注册中心的服务发现插件  [code](https://github.com/elgong/elgongRPC/tree/master/soa/register/redisPlugin)
  - 支持自定义**编解码插件****，可以通过实现 `codec`接口来增加自定义的编解码
    - 已实现基于`gob` 的编解码插件 [code](https://github.com/elgong/elgongRPC/blob/master/codec/codec_gob.go)
    - 已实现基于`msgpack`的编解码插件 [code](https://github.com/elgong/elgongRPC/blob/master/codec/codec_msgpack.go)
  - 支持自定义**protocol协议插件**，可以通过实现 `protocol` 接口来增加
    - 已提供默认插件 [code](https://github.com/elgong/elgongRPC/blob/master/protocol/default_protocol.go)
  - 支持自定义**连接池插件**，可以通过实现 `pool` 接口来实现增加
    - 已提供默认插件 [code](https://github.com/elgong/elgongRPC/blob/master/conn_pool/default_connPool.go)
  - 支持**其他插件**的自定义 ...

- **提供基于Redis的服务注册与服务发现**

  - **实现思路**：
    - 服务端定时注册服务：将服务注册到redis的set中， 对应的key 为 `namespace_`+ 服务名，value 为 ip + port
    - 客户端定时获取服务列表：客户端本地维护一个**服务列表的缓存**, 并通过实时任务去更新该缓存，发现异常服务地址，从redis 中删除

- **连接池管理**： 连接池管理连接，缓解频繁建立和释放连接对服务器资源的浪费

  - **实现思路**：为每个IP 建立一个池子，每个池子的底层利用**双向链表**管理起来；

  - **目前涵盖功能**：（通过`config.yaml`参数配置文件可以选择开启和关闭对应功能）

    1. 线程安全（底层操作链表时加锁实现安全）
    
    2. 惰性初始化 （用时再创建连接）
    
    3. 连接数量控制（支持最大连接数，闲置连接数，最大连接时长控制）
    
    4. 连接健康管理（定时处理掉失效连接）（服务端通过read 返回EOF来判断对方是否掉线； 客户端通过心跳包）
    5. 支持掉线自动重连机制（put 回池子时检查，单独开启线程重连接）
  
- **参数可配置**： 支持默认参数 & yaml 文件配置

    

##  elgongRPC V2.0 将会增加的功能

- 调用链设计（参考grpc）

- 负载均衡插件（两次随机负载均衡+轮询 + 带权）

- log 插件

**还未压测，待V2.0 再压测，现在功能还不完善**

## elgongRPC V1.0 使用DEMO

待补充。。



## commit 约定

- 添加新文件：
  - --new=文件功能
- 修改文件内容：
  - --edited=修改内容
- 其他：
  - --other=其他
- bug
  - --bugfix=bug内容
