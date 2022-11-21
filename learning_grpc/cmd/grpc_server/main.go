package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning_grpc/cmd/grpc_server/errors"
	"learning_grpc/cmd/grpc_server/matedatas"
	"learning_grpc/cmd/grpc_server/servers"
	"learning_grpc/pkg/grpc/chat"
	"learning_grpc/pkg/grpc/errorCode"
	"learning_grpc/pkg/grpc/login"
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
		//grpc.UnaryInterceptor(interceptor.MyUnaryServerInterceptor1),
		// 自定义流拦截器(可选)
		//grpc.StreamInterceptor(interceptor.MyStreamServerInterceptor1),

		// 多个拦截器链式启动(可选)
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	// grpc服务注册
	product.RegisterProductServer(s, &servers.Server{})
	user.RegisterUserServer(s, &servers.Server{})
	chat.RegisterChatServer(s, &servers.Server{})
	errorCode.RegisterErrorServer(s, &errors.Server{})
	login.RegisterLoginServer(s, &matedatas.Server{})

	// grpc服务反射(https://github.com/fullstorydev/grpcurl/releases)
	// 向grpc服务器本身获取proto文件信息
	reflection.Register(s)

	// grpc服务使用建立的tcp网络监听器(程序会停留在此处接收网络tcp数据包)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}
