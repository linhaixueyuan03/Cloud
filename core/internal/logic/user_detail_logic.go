package logic

import (
	"Cloud/core/model"
	"context"
	"errors"

	"Cloud/core/internal/svc"
	"Cloud/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.LoginDetailRequest) (resp *types.LoginDetailReply, err error) {
	resp = &types.LoginDetailReply{}
	ub := new(model.UserBasic)
	has, err := model.Engine.Where("identity=?", req.Idetity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户没有创建")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email

	return
}
