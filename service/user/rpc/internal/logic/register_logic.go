package logic

import (
	"Digital_trading_platform/common"
	"Digital_trading_platform/service/user/model"
	"Digital_trading_platform/service/user/rpc/internal/svc"
	"Digital_trading_platform/service/user/rpc/types/user"
	"context"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 注册 rpc 实现逻辑
func (l *RegisterLogic) Register(in *user.UserInfoWithPwd) (*user.Result, error) {
	// todo: add your logic here and delete this line
	c := l.ctx
	_, err := l.svcCtx.UserModel.FindOneByMobile(c, in.Mobile)
	switch err {
	case nil:
		return nil, status.Error(100, "该用户已存在")
	case model.ErrNotFound:
		newUser := &model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: string(common.PasswordCry(in.Code)),
		}
		res, err := l.svcCtx.UserModel.Insert(c, newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUser.Id, err = res.LastInsertId()
		return &user.Result{}, err
	default:
		return nil, err
	}
}
