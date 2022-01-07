package client_proxy

// client_proxy can be automatic generate

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testRpc/handler"
)

type RpcStub struct {
	*rpc.Client
}

func NewClient(protocol string, address string) RpcStub {
	log.Println("1. build connetion")
	con, err := net.Dial(protocol, address)
	if err != nil {
		log.Printf("err = %+v\n", err)
		return RpcStub{}
	}

	log.Println("2. use json transform")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(con))

	return RpcStub{client}
}

func (c *RpcStub) Hello(req string, res *string) error {
	log.Println("call function")
	err := c.Call(handler.RpcStuctName+".Show", req, res)
	if err != nil {
		return err
	}

	return nil
}
