package client_proxy

// client_proxy 客户端代理，在grpc中可以自动生成

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"learnRpc/handler"
)

type RpcStub struct {
	*rpc.Client
}

// Hello 客户端代理 stub
func (c *RpcStub) Hello(req string, res *string) error {
	fmt.Println("2.1 调用远程函数")
	err := c.Call(handler.RpcServiceName+".Show", req, res)
	if err != nil {
		return err
	}

	return nil
}

// NewClient 客户端代理 stub
func NewClient(protocol string, address string) RpcStub {
	fmt.Println("1.1 确定协议和目的端口")
	con, err := net.Dial(protocol, address)
	if err != nil {
		log.Printf("err = %+v\n", err)
		return RpcStub{}
	}

	// 默认是Gob协议
	fmt.Println("1.2 确定传输协议")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(con))

	return RpcStub{client}
}
