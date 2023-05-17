// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"Digital_trading_platform/service/user/rpc/internal/logic"
	"Digital_trading_platform/service/user/rpc/internal/svc"
	"Digital_trading_platform/service/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *user.UserInfoWithPwd) (*user.Result, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.IdRequest) (*user.UserInfo, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserInfo, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) SaveUser(ctx context.Context, in *user.IdRequest) (*user.Result, error) {
	l := logic.NewSaveUserLogic(ctx, s.svcCtx)
	return l.SaveUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *user.UserInfo) (*user.Result, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *user.IdRequest) (*user.Result, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserServer) ChangeUserPwd(ctx context.Context, in *user.PwdRequest) (*user.Result, error) {
	l := logic.NewChangeUserPwdLogic(ctx, s.svcCtx)
	return l.ChangeUserPwd(in)
}

func (s *UserServer) SendCode(ctx context.Context, in *user.SendCode) (*user.Code, error) {
	l := logic.NewSendCodeLogic(ctx, s.svcCtx)
	return l.SendCode(in)
}
