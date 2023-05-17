package logic

import (
	"context"

	"Digital_trading_platform/service/user/rpc/internal/svc"
	"Digital_trading_platform/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.IdRequest) (*user.Result, error) {
	// todo: add your logic here and delete this line

	return &user.Result{}, nil
}
