package errorCode

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning_grpc/pkg/grpc/errorCode"
)

type Server struct {
	errorCode.UnimplementedErrorServer
}

func (c *Server) ErrorServer(ctx context.Context, in *errorCode.ErrorReq) (out *errorCode.ErrorResp, err error) {
	// 定义简单错误
	//err = status.Error(codes.Unknown, "unknown error occurred!")

	// 定义带堆栈跟踪的错误
	stat := status.New(codes.Unknown, "unknown error occurred!")
	details, _ := stat.WithDetails(&errdetails.DebugInfo{
		Detail: "自定义的未知错误",
	})
	return &errorCode.ErrorResp{Msg: "错误默认返回"}, details.Err()
}
