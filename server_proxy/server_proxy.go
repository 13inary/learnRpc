package server_proxy

// server_proxy 服务端代理，在grpc中可以自动生成

import (
	"fmt"
	"io"
	"learnRpc/handler"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// RpcServicer 服务函数列表
type RpcServicer interface {
	Show(req string, res *string) error
}

// NewService 服务端代理 stub
func NewService(protocol string, address string, serviceName string, services RpcServicer) error {
	fmt.Println("1.1 新建服务端")
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Printf("err = %+v\n", err)
		return err
	}

	fmt.Println("1.2 注册服务（远程函数）")
	err = rpc.RegisterName(serviceName, services)
	if err != nil {
		log.Printf("err = %+v\n", err)
		return err
	}

	for {
		fmt.Println("1.3 ...持续监听链接...")
		con, err := listener.Accept()
		if err != nil {
			log.Printf("err = %+v\n", err)
			return err
		}

		log.Println("1.3.1 运行服务")
		//rpc.ServeConn(con)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(con))
	}
}

// httpRun 使用http协议的rpc
func httpRun() {
	log.Println("register handler")
	err := rpc.RegisterName("HelloRpc", &handler.RpcService{})
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	log.Println("define server")
	http.HandleFunc("/httprpc", func(w http.ResponseWriter, r *http.Request) {
		var con io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(con))
	})

	log.Println("listen port")
	http.ListenAndServe(":8080", nil)
}
