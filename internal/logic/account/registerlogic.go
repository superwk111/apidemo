package account

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"apidemo/internal/common"
	"apidemo/internal/model"
	"apidemo/internal/svc"
	"apidemo/internal/types"

	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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
	
	// Check if user exists
	exists, err := userModel.FindByName(l.ctx, req.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		l.Logger.Errorf("failed to query user: %v", err)
		return nil, common.DatabaseErrorf("failed to query user: %v", err)
	}
	if exists != nil {
		return nil, common.UserAlreadyExistsError()
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("failed to hash password: %v", err)
		return nil, common.DatabaseErrorf("failed to hash password: %v", err)
	}

	// Create user with hashed password
	now := time.Now()
	_, err = userModel.Insert(l.ctx, &model.User{
		Name:          req.Name,
		Password:      string(hashedPassword),  // Store the hashed password
		RegisterTime:  now,
		LastLoginTime: now,
	})
	if err != nil {
		l.Logger.Errorf("failed to create user: %v", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return nil, common.UserAlreadyExistsError()
		}
		return nil, common.DatabaseErrorf("failed to create user: %v", err)
	}

	// Return success response
	return &types.RegisterResp{
		Status: "SUCCESS",
	}, nil
}
