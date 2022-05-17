package logic

import (
	"context"

	"example-gozero/internal/svc"
	"example-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrdersLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	orderLogic *OrderLogic
}

func NewUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrdersLogic {
	return &UserOrdersLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		orderLogic: NewOrderLogic(ctx, svcCtx),
	}
}

func (l *UserOrdersLogic) UserOrders(req *types.UserOrdersRequest) (resp *types.UserOrdersReply, err error) {
	orders, err := l.orderLogic.ordersByUser(req.ID)
	if err != nil {
		return nil, err
	}

	return &types.UserOrdersReply{
		Orders: orders,
	}, nil
}
