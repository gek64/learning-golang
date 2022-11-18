package server

import (
	"learning_grpc/pkg/grpc/chat"
	"learning_grpc/pkg/grpc/login"
	"learning_grpc/pkg/grpc/product"
	"learning_grpc/pkg/grpc/user"
)

// Server 实现grpc服务
type Server struct {
	product.UnimplementedProductServer
	user.UnimplementedUserServer
	chat.UnimplementedChatServer
	login.UnimplementedLoginServer
}
