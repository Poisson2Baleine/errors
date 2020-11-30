package errors

import "errors"

var (
	ServerError           = errors.New("服务器内部错误")
	ParamError            = errors.New("参数错误")
	FrequentlyError       = errors.New("访问太过频繁")
	BusyError             = errors.New("服务器繁忙")
	PermissionDeniedError = errors.New("没有权限")
	NotFoundError         = errors.New("资源不存在")
	TimeOutError          = errors.New("请求超时")
	RpcCallError          = errors.New("RPC调用失败")
	RedisConnError        = errors.New("REDIS连接失败")
	MysqlConnError        = errors.New("MYSQL连接失败")
	EtcdConnError         = errors.New("ETCD连接失败")
	unknownError          = errors.New("未知错误")
)

func GetErrorCode(err error) int {
	switch err {
	case ServerError:
		return 50000
	case ParamError:
		return 40000
	case FrequentlyError:
		return 40010
	case BusyError:
		return 50010
	case PermissionDeniedError:
		return 40001
	case NotFoundError:
		return 40004
	case TimeOutError:
		return 50002
	case RpcCallError:
		return 50003
	case RedisConnError:
		return 56379
	case MysqlConnError:
		return 53306
	case EtcdConnError:
		return 52379
	default:
		return 55555
	}
}
