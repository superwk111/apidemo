package account

import (
	"context"
	"time"

	"apidemo/internal/common"
	"apidemo/internal/model"
	"apidemo/internal/svc"
	"apidemo/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userModel := model.NewUserModel(l.svcCtx.Mysql)

	// Verify user credentials
	user, err := userModel.FindByName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Errorf("failed to find user: %v", err)
		return nil, common.UserNotExistError()
	}

	// Compare hashed passwords
	if !common.CheckPasswordHash(req.Password, user.Password) {
		return nil, common.PasswordErrorError()
	}

	// Generate JWT token
	token, err := l.generateJWT(int64(user.Id))
	if err != nil {
		l.Logger.Errorf("failed to generate token: %v", err)
		return nil, common.NewGenerateTokenError()
	}

	// Return token in response
	return &types.LoginResp{
		Token: token,
	}, nil
}

func (l *LoginLogic) generateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.svcCtx.Config.JWTSecret))
}
