package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceOrderLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	OrderRepo repository.OrderRepository
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	OrderRepo := repository.NewOrderRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &PlaceOrderLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		OrderRepo: OrderRepo,
	}
}

func (l *PlaceOrderLogic) PlaceOrder(req *types.PlaceOrderReq) (resp *types.PlaceOrderRes, err error) {
	resp, _ = l.OrderRepo.PlaceOrder(l.ctx, req)
	return
}
