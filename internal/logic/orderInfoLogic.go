package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	OrderRepo repository.OrderRepository
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	orderRepo := repository.NewOrderRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &OrderInfoLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		OrderRepo: orderRepo,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoRes, err error) {
	resp, _ = l.OrderRepo.OrderInfo(l.ctx, req)
	return
}
