// @Title  服务注册 的统一接口
// @Author  elgong 2020.7.29
// @Update  elgong 2020.7.29
package register

const RegisterType = "register"

// Register 服务注册接口
type Register interface {
	// Ip 包含端口号
	Register(serviceName string, Ip string) error
}
