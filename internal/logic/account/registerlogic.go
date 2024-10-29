package account

import (
	"context"
	"database/sql"
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
	userModel := model.NewUserModel(l.svcCtx.Mysql)
	
	// Check existence with proper error handling
	exists, err := userModel.FindByUsername(l.ctx, req.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		l.Logger.Errorf("failed to query user: %v", err)
		return nil, errors.New("internal server error")
	}
	if exists != nil {
		return nil, errors.New("username already taken")
	}

	// Create user with better error handling
	now := time.Now()
	_, err = userModel.Insert(l.ctx, &model.User{
		Name:      req.Name,
		Password:      req.Password,  // Note: You should hash this password
		RegisterTime:  now,
		LastLoginTime: now,
	})
	if err != nil {
		l.Logger.Errorf("failed to create user: %v", err)
		return nil, errors.New("failed to create user")
	}

	return &types.RegisterResp{}, nil
}
