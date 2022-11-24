package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"learning_gozero/pkg/order/internal/config"
	"learning_gozero/pkg/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
