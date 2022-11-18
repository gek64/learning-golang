package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"learning_grpc/pkg/grpc/login"
)

func (s *Server) Login(ctx context.Context, in *login.LoginReq) (out *login.LoginResp, err error) {
	// 通过上下文的ctx来传递元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, value := range md {
			fmt.Println(key, value)
		}
	}
	return &login.LoginResp{Msg: "ok"}, nil
}
