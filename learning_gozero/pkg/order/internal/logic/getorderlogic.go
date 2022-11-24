package logic

import (
	"context"
	"learning_gozero/pkg/user/types/user"

	"learning_gozero/pkg/order/internal/svc"
	"learning_gozero/pkg/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	userResponse, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.OrderReply{
		Id:   userResponse.GetId(),
		Name: userResponse.GetName(),
	}, nil
}
