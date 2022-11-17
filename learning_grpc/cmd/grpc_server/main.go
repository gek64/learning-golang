package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning_grpc/pkg/grpc/chat"
	"learning_grpc/pkg/grpc/error_code"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
	"net"
)

func main() {
	startRpcServer()
}

// 启动RPC服务端
func startRpcServer() {
	// 注册tcp网络监听器
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}

	// grpc服务建立
	s := grpc.NewServer(
		// 自定义一元拦截器(可选)
		grpc.UnaryInterceptor(myUnaryServerInterceptor1),
		// 自定义流拦截器(可选)
		grpc.StreamInterceptor(myStreamServerInterceptor1),

		// 多个拦截器链式启动(可选)
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	// grpc服务注册
	product.RegisterProductServer(s, &server{})
	user.RegisterUserServer(s, &server{})
	chat.RegisterChatServer(s, &serverStream{})
	error_code.RegisterErrorServer(s, &errorCodeServer{})

	// grpc服务反射(https://github.com/fullstorydev/grpcurl/releases)
	// 向grpc服务器本身获取proto文件信息
	reflection.Register(s)

	// grpc服务使用建立的tcp网络监听器(程序会停留在此处接收网络tcp数据包)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
