package errorCode

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning_grpc/pkg/grpc/error_code"
)

type Server struct {
	error_code.UnimplementedErrorServer
}

func (c *Server) ErrorServer(ctx context.Context, in *error_code.ErrorReq) (out *error_code.ErrorResp, err error) {
	// 定义简单错误
	//err = status.Error(codes.Unknown, "unknown error occurred!")

	// 定义带堆栈跟踪的错误
	stat := status.New(codes.Unknown, "unknown error occurred!")
	details, _ := stat.WithDetails(&errdetails.DebugInfo{
		Detail: "自定义的未知错误",
	})
	return &error_code.ErrorResp{Msg: "错误默认返回"}, details.Err()
}
