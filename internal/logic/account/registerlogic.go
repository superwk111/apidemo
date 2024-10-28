package account

import (
	"context"
	"errors"
	"time"

	"apidemo/internal/model"
	"apidemo/internal/svc"
	"apidemo/internal/types"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	//1. according the username to find the user , if exist return error
	userModel := model.NewUserModel(l.svcCtx.Mysql)
	user, err := userModel.FindByUsername(l.ctx, req.UserName)
	if err != nil {
		l.Logger.Error("inquire user failed", err)
		return nil, err
	}
	if user != nil {
		return nil, errors.New("user already exists")
	}
	//2. if not exist, create a new user, ingested from req
	_, err = userModel.Insert(l.ctx, &model.User{
		Username :          req.UserName,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	//3. return resp
	return &types.RegisterResp{}, err
}
