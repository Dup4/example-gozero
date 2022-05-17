package logic

import (
	"context"

	"example-gozero/internal/svc"
	"example-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductLogic) Product(req *types.ProductRequest) (*types.ProductReply, error) {
	return l.productByID(req.ID)
}

func (l *ProductLogic) productByID(id int64) (*types.ProductReply, error) {
	return &types.ProductReply{
		ID:    id,
		Name:  "apple watch 3",
		Price: 3333.33,
		Count: 99,
	}, nil
}
