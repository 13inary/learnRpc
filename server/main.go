package main

import (
	"fmt"
	"learnRpc/handler"
	"learnRpc/server_proxy"
	"log"
)

// main 这里是面向业务的代码
func main() {
	fmt.Println("\n1. 运行监听，注册服务")
	err := server_proxy.NewService("tcp", ":8080", handler.RpcServiceName, &handler.RpcService{})
	if err != nil {
		log.Println(err)
		return
	}
}
