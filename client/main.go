package main

import (
	"fmt"
	"log"

	"learnRpc/client_proxy"
)

// main 这里是面向业务的代码
func main() {
	fmt.Println("\n1. 新建客户端")
	client := client_proxy.NewClient("tcp", "localhost:8080")

	fmt.Println("\n2. 使用客户端 做远程调用")
	var res string
	err := client.Hello("i am client", &res)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	fmt.Println("\n3. 处理返回值")
	log.Printf("res = %+v\n", res)
}
