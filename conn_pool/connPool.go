package conn_pool

import (
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
		return connPoo.(connPool).connStack.pop(), nil
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
		//idletime: opt.idletime,
		//maxLifetime: nil,
	}

	// 创建连接栈 + 一个连接
	connStack := connDeque{}
	conn, err := net.Dial("tcp", address)


	if err != nil {
		return connPoo, err
	}

	cInPool := connInPool{conn, address, &connPoo, time.Now(), nil}
	connStack.push(cInPool)

	connPoo.connStack = &connStack

	return connPoo, nil
}

type connDeque struct {
	top     *connInPool
	bottom  *connInPool  // 底永远为nil
	lock    sync.Mutex
	size    int32
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
	if c.size > 0 {
		conn := c.top
		c.top = c.top.next
		c.size--
		conn.updatedtime = time.Now()
		////////////////// 异常待补充
		return conn
	}

	return nil
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
}

// putBack 放回
func (c *connInPool) PutBack(){
	c.cp.connStack.push(*c)
}

