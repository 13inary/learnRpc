package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testRpc/handler"
	"testRpc/server_proxy"
)

func tcpRun() {
	log.Println("1. new server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("err = %+v\n", err)
		return
	}

	log.Println("2. register handler")
	err = server_proxy.RegisterServer(&handler.RpcStruct{})
	if err != nil {
		log.Printf("err = %+v\n", err)
		return
	}

	for {
		log.Println("3. listen request")
		con, err := listener.Accept()
		if err != nil {
			log.Printf("err = %+v\n", err)
			return
		}

		log.Println("4. run server by json transform")
		//rpc.ServeConn(con)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(con))
	}

}

func main() {
	tcpRun()
}

func httpRun() {
	log.Println("register handler")
	err := rpc.RegisterName("HelloRpc", &handler.RpcStruct{})
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
