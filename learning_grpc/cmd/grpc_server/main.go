package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// grpc服务反射(https://github.com/fullstorydev/grpcurl/releases)
	// 向grpc服务器本身获取proto文件信息
	reflection.Register(s)

	// grpc服务使用建立的tcp网络监听器(程序会停留在此处接收网络tcp数据包)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
