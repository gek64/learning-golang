package main

import (
	"fmt"

	"learning_gozero/pkg/user/internal/config"
	"learning_gozero/pkg/user/internal/server"
	"learning_gozero/pkg/user/internal/svc"
	"learning_gozero/pkg/user/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var c config.Config
	conf.MustLoad("configs/user.json", &c)

	// 服务的配置文件
	ctx := svc.NewServiceContext(c)

	// 使用zrpc来注册grpc服务,会自动调用zrpc内部的上下文来控制服务(取消)
	s := zrpc.MustNewServer(c.RpcServerConf,
		// 注册grpc服务的函数
		func(grpcServer *grpc.Server) {
			// 用户grpc服务
			user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))
			// 反射服务
			reflection.Register(grpcServer)
		})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
