// @Title 默认TCP连接池插件
// @Description
// @Author  elgong 2020.7.
// @Update  elgong 2020.7.
package conn_pool

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	. "github.com/elgong/elgongRPC/plugin_centre"
)

// Pools 注册进入插件管理中心
func init() {
	pools := DefaultPools{ConnType, "defaultConnPool", sync.Map{}}
	PluginCenter.Register(pools.Type, pools.Name, &pools)
}

// ConnInPool 连接池中连接的封装接口,外部也会用到
type ConnInPool interface {
	PutBack()                      // 放回池子
	SetDead()                      // 设置该连接掉线了
	Send(data []byte) (int, error) // 发送数据
}

// Pools 统一管理的连接池结构体，
// 实现了 pool 接口
type DefaultPools struct {
	Type    PluginType
	Name    PluginName
	poolMap sync.Map
}

// todo: 需要异常补充
func (p *DefaultPools) GetConn(ctx context.Context, address string) (*connInPool, error) {
	select {
	case <-ctx.Done():
		fmt.Println("停止了...")
		return nil, nil
	default:
	}

	// 如果还未创建连接池
	if _, OK := p.poolMap.Load(address); !OK {
		connPoo, err := newConnPool(address)
		if err != nil {
			return nil, errors.New("连接池创建失败")
		}

		p.poolMap.Store(address, connPoo)
	}

	// 获取连接池
	connPoo, isOk := p.poolMap.Load(address)
	// 获取到连接
	if isOk {
		connInpool := connPoo.(connPool).connDeq.pop()
		// 如果为空？ 要新建啊
		if connInpool == nil {
			return nil, errors.New("无连接")
		}
		return connInpool, nil
	}

	return nil, nil
}

func (p *DefaultPools) Close() {
	// 遍历关闭即可，等待实现中
}

// connPool 单ip 的连接池
type connPool struct {
	// 参数配置
	initialCap int // 初始的连接容量
	maxCap     int // 最大的连接容量
	maxIdle    int // 最大的闲置时间
	//idletime    time.Duration
	//maxLifetime time.Duration
	timeOut int // 建立连接超时时间

	// 超时重连
	failReconnect       bool // 是否掉线重连
	failReconnectSecond int  // 重连等待时间
	failReconnectTime   int  // 重连次数

	connDeq *connDeque // 连接池底层的链表

	// 定时器，执行定时管理任务
	ticker       *time.Ticker
	isTickerOpen bool // 是否开启健康管理
	tickerTime   int  // 定时时长
}

// 创建单ip连接池
// 正常情况下，应该不会遇到多个并发建立多个池子冲突的情况
func newConnPool(address string, opts ...ModifyConnOption) (connPool, error) {
	// opt 默认参数
	var opt = defaultConnOptions

	// 修改参数
	for _, o := range opts {
		o(&opt)
	}

	// 创建单ip连接池
	connPoo := connPool{
		initialCap:          opt.initialCap,
		maxCap:              opt.maxCap,
		maxIdle:             opt.maxIdle,
		failReconnect:       opt.failReconnect,
		failReconnectSecond: opt.failReconnectSecond,
		failReconnectTime:   opt.failReconnectTime,
		isTickerOpen:        opt.isTickerOpen,
		tickerTime:          opt.tickerTime,
		timeOut:             opt.timeOut,
	}

	// 创建连接链表
	connDeq := connDeque{address: address, cp: &connPoo}

	// 建立tcp 连接
	conn, err := net.DialTimeout("tcp", address, time.Second*time.Duration(connPoo.timeOut))

	if err != nil {
		return connPoo, err
	}

	// 创建一个连接
	cInPool := connInPool{conn, address, &connPoo, time.Now(), nil, nil, true}
	connDeq.push(cInPool)
	connPoo.connDeq = &connDeq

	// 如果开启了健康管理
	if connPoo.isTickerOpen {
		// 定时时长
		connPoo.ticker = time.NewTicker(time.Second * time.Duration(connPoo.tickerTime))
		// 创建该池子对应的健康管理线程
		go func(connPoo *connPool) {
			for _ = range connPoo.ticker.C {
				// 取出每一个连接，检查是否存活
				n := connPoo.connDeq.getSize()
				fmt.Println("connPool healthy manage: have ", n, "conn in pool when ", time.Now())

				// 如果开启了最大闲置连接数限制
				if connPoo.maxCap > 0 && n > connPoo.maxCap {
					// 先直接丢弃
					for i := 0; i < (n - connPoo.maxCap); i++ {
						connPoo.connDeq.popBottom()
					}
					// 修改当前n的大小
					n = n - connPoo.maxCap
				}

				// 遍历处理 n 次
				for i := 0; i < n && connPoo.connDeq.top != nil; i++ {
					// pop 已经枷锁啦
					// connPoo.connStack.lock.Lock()
					// 保证有连接再 pop
					if connPoo.connDeq.getSize() > 0 {
						cInPool := connPoo.connDeq.popBottom()
						cInPool.Conn.SetWriteDeadline(time.Now().Add(time.Duration(100) * time.Millisecond))

						// 发送心跳包
						_, err := cInPool.Conn.Write([]byte{0xaa, 0xbb})
						if err != nil {
							// 这样会有异常吗，，，，未来再修复吧
							cInPool.Conn.Close()
							continue
						}

						// 最大空闲时间
						if connPoo.maxIdle > 0 && time.Now().Sub(cInPool.updatedtime) >= time.Second*time.Duration(connPoo.maxIdle) {
							// 空闲连接超时啦, 关闭
							cInPool.Conn.Close()
							continue
						}
						// 放回去啦
						connPoo.connDeq.push(*cInPool)
					}
				}
			}
		}(&connPoo)
	}

	return connPoo, nil
}

// connDeque 与连接池绑定的底层栈实现
type connDeque struct {
	top     *connInPool // 顶
	bottom  *connInPool // 底
	lock    sync.Mutex
	size    int
	address string
	cp      *connPool // 指向所属的pool
}

// push 放入conn 线程安全
func (c *connDeque) push(inPool connInPool) {

	if &inPool == nil {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	// 栈中已经存在元素时
	if c.size > 0 {
		inPool.next = c.top // 新元素next指向当前的栈顶
		inPool.pre = nil    // 新元素pre 指向 nil
		c.top.pre = &inPool // 旧栈顶pre 指向 新元素
		c.top = &inPool
	} else {
		// 栈中无元素时, 先设置栈底元素
		c.bottom = &connInPool{}
		c.bottom.next = nil    // 栈底next指空
		c.bottom.pre = &inPool // 栈底pre指向新元素
		inPool.next = c.bottom // 新元素next指向栈底
		inPool.pre = nil       // 新元素pre
		c.top = &inPool        // top更新
	}
	c.size++
	// 操作时间
	// inPool.updatedtime = time.Now()

}

// todo: 异常补充
func (c *connDeque) pop() *connInPool {

	c.lock.Lock()
	defer c.lock.Unlock()
	// 不存在
	if c.size <= 0 {
		conn, err := net.Dial("tcp", c.address)
		if err != nil {
			return nil
		}
		cInPool := connInPool{conn, c.address, c.cp, time.Now(), nil, nil, true}
		// c.push(cInPool)
		return &cInPool
	}

	// 将当前的栈顶返回,先保留起来
	conn := c.top
	// 栈顶向下移动
	c.top = c.top.next
	c.top.pre = nil

	// 要返回的元素,清空指向的东西
	conn.pre = nil
	conn.next = nil

	c.size--
	return conn
}

// 队列前端取出  相当于pop
func (c *connDeque) poll() *connInPool {
	return c.pop()
}

// 队列后端取出
func (c *connDeque) popBottom() *connInPool {
	c.lock.Lock()
	defer c.lock.Unlock()
	// 不存在
	if c.size <= 0 {
		return nil
	}

	// 取到倒数二节点
	conn := c.bottom.pre
	// 从链表中移除 倒数二节点
	c.bottom.pre = conn.pre

	if conn.pre != nil {
		conn.pre.next = c.bottom
	}

	conn.next = nil
	conn.pre = nil

	c.size--
	// conn.updatedtime = time.Now()
	////////////////// 异常待补充
	return conn
}

// getSize 得到大小
func (c *connDeque) getSize() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.size
}

// connInPool 连接的链式封装- 连接池的底层原子结构
// implenment ConnInPool 接口
type connInPool struct {
	Conn        net.Conn
	address     string
	cp          *connPool // 关联池子，方便放回
	updatedtime time.Time

	next   *connInPool // 指向下一个
	pre    *connInPool
	isLive bool // 是否存活，在 PutBack  时检测重新连接或者丢弃
}

// putBack 放回池子
func (c *connInPool) PutBack() {

	// 最大连接数限制
	if c.cp.maxCap > 0 && c.cp.connDeq.size >= c.cp.maxCap {
		c.Conn.Close()
	}
	// 如果连接失效 && 支持超时重连
	// 开启单独线程处理
	if !c.isLive {
		// 支持重连
		if c.cp.failReconnect {
			go c.reconnect()
			return
		}
		// 不支持的话，断开连接
		c.Conn.Close()
		return
	}
	// 否则，正常逻辑放池子
	c.cp.connDeq.push(*c)
}

// SetDead 发现连接失效时，修改连接状态，在putback 放回是统一处理
func (c *connInPool) SetDead() {
	c.isLive = false
}

// reconnect 超时重连
func (c *connInPool) reconnect() {
	var conn net.Conn
	var err error
	for i := 1; i < c.cp.failReconnectTime; i++ {

		conn, err = net.DialTimeout("tcp", c.address, time.Second*time.Duration(c.cp.failReconnectSecond))
		time.Sleep(time.Duration(time.Second))
		// 连接成功后，直接退出
		if err == nil {
			break
		}
	}

	if err != nil {
		return
	}

	c.Conn = conn
	c.cp.connDeq.push(*c)
}

// Send 把数据包循环发送出去
func (c *connInPool) Send(data []byte) (int, error) {
	var sendDataLen = 0
	var err error

	// 循环发送数据，包可能过大，so循环
	for sendDataLen = 0; sendDataLen < len(data); {
		n, err := c.Conn.Write(data[sendDataLen:])
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		sendDataLen += n
	}

	// 更新时间设为当前
	c.updatedtime = time.Now()

	return sendDataLen, err
}
