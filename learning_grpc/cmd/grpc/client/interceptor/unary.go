package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func MyUnaryClientInterceptor1(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 客户端发送信息前的预处理
	fmt.Println("[pre] my unary client interceptor 1", method, req)

	// 执行发送信息
	err := invoker(ctx, method, req, res, cc, opts...)

	// 客户都发送信息后的后处理
	fmt.Println("[post] my unary client interceptor 1", res)
	return err
}
