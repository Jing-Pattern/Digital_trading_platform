package logic

import (
	"context"

	"Digital_trading_platform/service/user/rpc/internal/svc"
	"Digital_trading_platform/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.SendCode) (*user.Code, error) {
	// todo: add your logic here and delete this line

	return &user.Code{}, nil
}
