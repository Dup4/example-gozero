package logic

import (
	"context"

	"example-gozero/internal/svc"
	"example-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.OrderRequest) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *OrderLogic) ordersByUser(uid int64) ([]*types.OrderReply, error) {
	if uid == 123 {
		// It should actually be queried from database or cache
		return []*types.OrderReply{
			{
				ID:       "236802838635",
				State:    1,
				CreateAt: "2022-5-12 22:59:59",
			},
			{
				ID:       "236802838636",
				State:    1,
				CreateAt: "2022-5-10 20:59:59",
			},
		}, nil
	}

	return nil, nil
}
