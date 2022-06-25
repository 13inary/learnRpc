package handler

// handler 服务端 的远程函数列表

const (
	RpcServiceName = "RpcService"
)

// RpcService 添加远程函数 的业务逻辑
type RpcService struct{}

func (s *RpcService) Show(req string, res *string) (err error) {
	*res = "i am service. your parameter is : " + req
	return nil
}
