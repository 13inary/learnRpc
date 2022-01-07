package server_proxy

// server_proxy can be automatic generate

import (
	"net/rpc"
	"testRpc/handler"
)

type RpcStruct2 interface {
	Show(req string, res *string) error
}

func RegisterServer(s RpcStruct2) error {
	return rpc.RegisterName(handler.RpcStuctName, s)
}
