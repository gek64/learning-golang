package main

import (
	"fmt"
	"learning_grpc/cmd/grpc_client/clients"
	"learning_grpc/cmd/grpc_client/matedatas"
	"log"
)

func main() {
	fmt.Println("一对一 rpc")
	err := clients.StartRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n服务端流 rpc")
	err = clients.StartServerStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n客户端流 rpc")
	err = clients.StartClientStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n双向流 rpc")
	err = clients.StartBiStreamsRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n错误代码测试")
	err = clients.TestRpcErrorCode()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n元数据发送")
	err = matedatas.Login()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n元数据接受")
	err = matedatas.Logout()
	if err != nil {
		log.Panicln(err)
	}
}
