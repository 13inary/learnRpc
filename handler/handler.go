package handler

// handler must define in advance

const (
	RpcStuctName = "RpcStruct"
)

type RpcStruct struct{}

func (s *RpcStruct) Show(req string, res *string) (err error) {
	*res = "rpc : " + req
	return nil
}
