package main

import (
	"fmt"
	"learning_grpc/cmd/grpc/client/client"
	"log"
)

func main() {
	fmt.Println("一对一 rpc")
	err := client.StartRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n服务端流 rpc")
	err = client.StartServerStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n客户端流 rpc")
	err = client.StartClientStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n双向流 rpc")
	err = client.StartBiStreamsRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n错误代码测试")
	err = client.TestRpcErrorCode()
	if err != nil {
		log.Panicln(err)
	}
}
