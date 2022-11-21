package matedatas

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learning_grpc/pkg/grpc/login"
)

func (s *Server) Logout(ctx context.Context, in *login.LoginReq) (out *login.LoginResp, err error) {
	// 通过上下文的ctx来传递元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, value := range md {
			fmt.Println(key, value)
		}
	}

	// 设置返回的上下文标志头部
	headerMD := metadata.New(map[string]string{
		"type": "logout",
		"from": "server",
		"in":   "header",
	})
	err = grpc.SetHeader(ctx, headerMD)
	if err != nil {
		return nil, err
	}

	// 设置返回的上下文标志尾部
	trailerMD := metadata.New(map[string]string{
		"type": "logout",
		"from": "server",
		"in":   "trailer",
	})
	err = grpc.SetTrailer(ctx, trailerMD)
	if err != nil {
		return nil, err
	}

	return &login.LoginResp{Msg: "ok"}, nil
}
