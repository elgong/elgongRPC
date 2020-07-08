package interceptor

import (
	"context"
	"fmt"
)

// 要拦截的方法，起个别名
type CallFunc func(ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}) error


// 拦截器
type Interceptor func(ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}, call CallFunc) error

// 定义一个拦截器数组类型
type Interceptors []Interceptor

func (i Interceptors) GetInterceptors(ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{},callFunc CallFunc) error{

	if len(i) <= 1 {
		// 如果没有拦截器, 则直接调用被拦截的方法
		if len(i) == 0 {
			return callFunc(ctx, serviceName, methodName, reqBody, rspBody)
		}
		// 只有一个拦截器
		return i[0](ctx, serviceName, methodName, reqBody, rspBody, callFunc)
	} else {
		// 多个拦截器时， 需要借助递归来实现
		return func(ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}) error {

			var index = 0

			return i.getFirst(index, ctx, serviceName, methodName, reqBody, rspBody, callFunc)
			// return nil
		}(ctx, serviceName, methodName, reqBody, rspBody)
	}
}

func (i Interceptors)getFirst(index int, ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}, callFunc CallFunc) error{

	if len(i) <= index {
		return callFunc(ctx, serviceName, methodName, reqBody, rspBody)
	}

	return i.getFirst(index + 1, ctx, serviceName, methodName, reqBody, rspBody, callFunc)
}

func (i Interceptors) Register(interceptor Interceptor){
	fmt.Println("注册了插件")

	if interceptor == nil {
		return
	}
	 i = append(i, interceptor)

	 fmt.Println(len(i))
}



var Interceptorss *Interceptors
func init(){
	Interceptorss = &Interceptors{}
	Interceptorss.Register(nil)
}



