package logic

import (
	"context"

	"Digital_trading_platform/service/user/rpc/internal/svc"
	"Digital_trading_platform/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeUserPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserPwdLogic {
	return &ChangeUserPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeUserPwdLogic) ChangeUserPwd(in *user.PwdRequest) (*user.Result, error) {
	// todo: add your logic here and delete this line

	return &user.Result{}, nil
}
