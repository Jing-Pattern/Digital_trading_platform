package logic

import (
	"Digital_trading_platform/service/user/api/internal/types"
	"Digital_trading_platform/service/user/rpc/types/user"
	"context"
	"time"

	"Digital_trading_platform/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(t *types.RegisterRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	ctx, CancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer CancelFunc()
	_, err = l.svcCtx.UserRpc.Register(ctx, &user.UserInfoWithPwd{
		Name:    t.Name,
		Gender:  t.Gender,
		Country: t.Country,
		Mobile:  t.Mobile,
		Code:    t.Password,
	})

	if err != nil {
		logx.Info("logic Register")
		return nil, err
	}
	return &types.Response{}, nil
}
