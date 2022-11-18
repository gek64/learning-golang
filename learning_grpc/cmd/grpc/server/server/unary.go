package server

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning_grpc/internal/mod"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
)

// GetProduct 实现GetProduct服务逻辑
func (s *Server) GetProduct(ctx context.Context, in *product.ProductReq) (out *product.ProductResp, err error) {
	p := product.ProductResp{
		Id:    in.GetId(),
		Name:  "测试用商品",
		SKU:   "A1000",
		Price: 100,
	}
	return &p, ctx.Err()
}

// GetUser 实现GetUser服务逻辑
func (s *Server) GetUser(ctx context.Context, in *user.UserReq) (out *user.UserResp, err error) {
	// 新建时间
	newTime := mod.GetRandomTime()
	// 填充返回结构体
	u := user.UserResp{
		Id:       in.GetId(),
		Name:     "bob",
		Birthday: timestamppb.New(newTime),
	}
	return &u, ctx.Err()
}
