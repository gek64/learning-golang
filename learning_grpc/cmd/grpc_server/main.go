package main

import (
	"google.golang.org/grpc"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
	"net"
)

func main() {
	// 注册tcp网络监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}

	// grpc服务建立
	s := grpc.NewServer()

	// grpc服务注册
	product.RegisterProductServer(s, &server{})
	user.RegisterUserServer(s, &server{})

	// grpc服务使用建立的tcp网络监听器(程序会停留在此处接收网络tcp数据包)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
