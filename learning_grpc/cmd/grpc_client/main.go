package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("一对一 rpc")
	err := startRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n服务端流 rpc")
	err = startServerStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n客户端流 rpc")
	err = startClientStreamRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n双向流 rpc")
	err = startBiStreamsRpc()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("\n错误代码测试")
	err = testRpcErrorCode()
	if err != nil {
		log.Panicln(err)
	}
}
