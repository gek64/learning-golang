package logic

import (
	"context"

	"learning_gozero/pkg/user/internal/svc"
	"learning_gozero/pkg/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	testUser := user.UserResponse{
		Id:     in.GetId(),
		Name:   "Bob",
		Gender: "male",
	}
	return &testUser, nil
}
