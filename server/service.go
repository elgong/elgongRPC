// @Title  服务的结构体
// @Author  elgong 2020.7.18
// @Update  elgong 2020.7.21
package server

import (
	"fmt"
	"reflect"
)

// Service 对应一个服务
type Service struct {
	serviceName string
	typ         reflect.Type
	rcvr        reflect.Value
	methodsMap  map[string]reflect.Method
}

// 调用该服务
func (s *Service) invoke(methodName string, args interface{}, retArgs interface{}) {
	fmt.Println("服务端调用了*****")

	if function, OK := s.methodsMap[methodName]; OK {

		function.Func.Call([]reflect.Value{s.rcvr, reflect.ValueOf(args), reflect.ValueOf(retArgs)})
	}

}
