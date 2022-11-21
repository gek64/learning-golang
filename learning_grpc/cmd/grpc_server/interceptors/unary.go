package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// MyUnaryServerInterceptor1 一元拦截器(相当于处理接收请求后需要先执行的中间件,例如鉴权、日志等操作)
func MyUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 执行前操作
	log.Println("[pre] 一元拦截器1:", info.FullMethod)

	// 执行服务函数
	res, err := handler(ctx, req)

	// 执行后操作
	log.Println("[post] 一元拦截器1 运行结束")

	return res, err
}
