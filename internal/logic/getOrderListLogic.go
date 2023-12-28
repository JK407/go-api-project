package logic

import (
	"context"
	"go_project/internal/repository"

	"go_project/internal/svc"
	"go_project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	OrderRepo repository.OrderRepository
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	orderRepo := repository.NewOrderRepository(svcCtx.Gdb, svcCtx.Rdb)
	return &GetOrderListLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		OrderRepo: orderRepo,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListReq) (resp *types.GetOrderListRes, err error) {
	resp, _ = l.OrderRepo.GetOrderList(l.ctx, req)
	return
}
