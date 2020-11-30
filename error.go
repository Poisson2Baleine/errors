package errors

import "errors"

var (
	serverError           = errors.New("服务器内部错误")
	paramError            = errors.New("参数错误")
	frequentlyError       = errors.New("访问太过频繁")
	busyError             = errors.New("服务器繁忙")
	permissionDeniedError = errors.New("没有权限")
	notFoundError         = errors.New("资源不存在")
	timeOutError          = errors.New("请求超时")
	rpcCallError          = errors.New("RPC调用失败")
	redisConnError        = errors.New("REDIS连接失败")
	mysqlConnError        = errors.New("MYSQL连接失败")
	etcdConnError         = errors.New("ETCD连接失败")
	unknownError          = errors.New("未知错误")
)

func GetErrorCode(err error) int {
	switch err {
	case serverError:
		return 50000
	case paramError:
		return 40000
	case frequentlyError:
		return 40010
	case busyError:
		return 50010
	case permissionDeniedError:
		return 40001
	case notFoundError:
		return 40004
	case timeOutError:
		return 50002
	case rpcCallError:
		return 50003
	case redisConnError:
		return 56379
	case mysqlConnError:
		return 53306
	case etcdConnError:
		return 52379
	default:
		return 55555
	}
}
