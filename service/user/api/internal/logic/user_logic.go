package logic

import (
	"Digital_trading_platform/service/user/api/internal/svc"
	"Digital_trading_platform/service/user/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Login(t *types.LoginRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return
}
