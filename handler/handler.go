package handler

// handler must define in advance

const (
	RpcServiceName = "RpcService"
)

type RpcService struct{}

func (s *RpcService) Show(req string, res *string) (err error) {
	*res = "i am service. your parameter is : " + req
	return nil
}
