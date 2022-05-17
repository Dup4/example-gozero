package logic

import (
	"context"

	"example-gozero/internal/svc"
	"example-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	productLogic *ProductLogic
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		productLogic: NewProductLogic(ctx, svcCtx),
	}
}

const (
	success = 0
	failure = -1
)

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateRequest) (*types.OrderCreateReply, error) {
	product, err := l.productLogic.productByID(req.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Count > 0 {
		return &types.OrderCreateReply{Code: success}, nil
	}

	return &types.OrderCreateReply{Code: failure}, nil
}
