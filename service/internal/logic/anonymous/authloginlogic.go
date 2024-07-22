package anonymous

import (
	"ai-gallery/service/internal/logic/basic"
	"context"

	pkg "ai-gallery/pkg/jwt"
	"ai-gallery/service/internal/errors"
	"ai-gallery/service/internal/model"
	"ai-gallery/service/internal/svc"
	"ai-gallery/service/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginLogic {
	return &AuthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLoginLogic) AuthLogin(req *types.AuthLoginRequest) (resp *types.AuthLoginResponse, err error) {
	user, err := basic.ValidateAccount(l.ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	JWTClaims := pkg.Claims{
		ID:       user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	if user.Role == model.ADMIN {
		JWTClaims.Audience = jwt.ClaimStrings{pkg.AudienceAdmin}
	} else {
		JWTClaims.Audience = jwt.ClaimStrings{pkg.AudienceUser}
	}
	token, err := pkg.GetJwtToken(
		l.svcCtx.Config.JWT.JwtKey,
		l.svcCtx.Config.JWT.JwtIssuer,
		l.svcCtx.Config.JWT.JwtExpire,
		JWTClaims)
	if err != nil {
		return nil, errors.ErrAccountTokenGen(l.ctx)
	}
	return &types.AuthLoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}, nil
}
