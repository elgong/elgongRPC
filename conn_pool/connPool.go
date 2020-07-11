package conn_pool

import (
	"errors"
	"net"
	"sync"
	"time"
)

var Pool = Pools{}


// pool 连接池接口
type pool interface {
	GetConn(address string) (net.Conn, error)
	Close() error
}

//
type Pools struct {
	poolMap sync.Map
}

// todo: 异常补充
func (p *Pools) GetConn(address string) (*connInPool, error){
	// 如果还未创建连接池
	if _, OK := p.poolMap.Load(address); !OK{
		connPoo, err := NewConnPool(address)

		if err != nil {
			return nil, err
		}

		p.poolMap.Store(address, connPoo)
	}

	// 获取连接池
	connPoo, isOk := p.poolMap.Load(address)
	// 获取到连接
	if isOk  {
		connInpool := connPoo.(connPool).connStack.pop()
		// 如果为空？ 要新建啊
		if connInpool == nil{
			return nil, errors.New("无连接")
		}
		return connInpool, nil
	}

	return nil, nil
}

func (p *Pools) Close(){

}

// connPool 单ip 的连接池
type connPool struct {

	connStack    *connDeque
	initialCap  int
	maxCap      int
	maxIdle     int
	idletime    time.Duration
	maxLifetime time.Duration
	cleanerCh   chan struct{}
	failReconnect bool  // 超时重连参数
	failReconnectSecond int

	factory func() (net.Conn, error)
}

//func (c *connPool) get(address string) *connInPool{
//
//}

// 创建单ip连接池
func NewConnPool(address string, opts...ModifyConnOption) (connPool, error){
	// opt 默认参数
	var opt = defaultConnOptions

	// 修改参数
	for _, o := range opts{
		o(&opt)
	}

	// 创建单ip连接池
	connPoo := connPool{
		initialCap: opt.initialCap,
		maxCap: opt.maxCap,
		maxIdle: opt.maxIdle,
		failReconnect: opt.failReconnect,
		failReconnectSecond: opt.failReconnectSecond,
		//idletime: opt.idletime,
		//maxLifetime: nil,
	}

	// 创建连接栈 + 一个连接
	connStack := connDeque{address:address, cp: &connPoo}
	conn, err := net.Dial("tcp", address)


	if err != nil {
		return connPoo, err
	}

	cInPool := connInPool{conn, address, &connPoo, time.Now(), nil, true}
	connStack.push(cInPool)

	connPoo.connStack = &connStack

	return connPoo, nil
}

type connDeque struct {
	top     *connInPool
	bottom  *connInPool  // 底永远为nil
	lock    sync.Mutex
	size    int32
	address string

	cp *connPool
}

func (c *connDeque) push(inPool connInPool){

	if &inPool == nil {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	// 栈中已经存在元素时
	if c.size > 0 {
		inPool.next = c.top
		c.top = &inPool
	} else {
	// 栈中无元素
		c.bottom = nil
		inPool.next = c.bottom
		c.top = &inPool
	}
	c.size++
	// 操作时间
	inPool.updatedtime = time.Now()

}

// todo: 异常补充
func (c *connDeque) pop() *connInPool{

	c.lock.Lock()
	defer c.lock.Unlock()
	// 存在
	if c.size <= 0 {
		conn, err := net.Dial("tcp", c.address)
		if err != nil {
			return nil
		}
		cInPool := connInPool{conn, c.address, c.cp, time.Now(), nil, true}
		// c.push(cInPool)
		return &cInPool
	}
	conn := c.top
	c.top = c.top.next
	c.size--
	conn.updatedtime = time.Now()
	////////////////// 异常待补充
	return conn

}

func (c *connDeque) getSize() int32 {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.size
}


// connInPool 连接的链式封装
type connInPool struct {
	Conn net.Conn
	address string
	cp *connPool  // 关联池子，方便放回
	updatedtime time.Time

	next *connInPool  // 指向下一个
	// pre *connInPool
	isLive bool  // 是否存活，在 PutBack  时检测重新连接或者丢弃
}

// putBack 放回池子
func (c *connInPool) PutBack(){
	// 如果连接失效 && 支持超时重连
	// 开启单独线程处理
	if !c.isLive  {
		// 支持重连
		if c.cp.failReconnect {
			go c.reconnect()
			return
		}
		// 不支持的话，丢点
		return
	}

	// 否则，正常逻辑放池子
	c.cp.connStack.push(*c)
}

func (c *connInPool) SetDead(){
	c.isLive = false
}

// reconnect 超时重连
func (c *connInPool) reconnect(){
	var conn net.Conn
	var err error
	for i := 1; i<5;i++{

		conn, err = net.DialTimeout("tcp", c.address, time.Second * time.Duration(c.cp.failReconnectSecond))
		time.Sleep(time.Duration(time.Second))
	}


	if err != nil {
		return
	}

	c.Conn = conn
	c.cp.connStack.push(*c)
}

