package main

import (
	"log"
	"testRpc/client_proxy"
)

func main() {
	client := client_proxy.NewClient("tcp", "localhost:8080")

	var res string
	err := client.Hello("rpc", &res)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	log.Printf("res = %+v\n", res)
}
