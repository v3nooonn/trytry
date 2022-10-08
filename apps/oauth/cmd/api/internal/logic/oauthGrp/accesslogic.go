package oauthGrp

import (
	"context"
	"net/http"
	"time"

	"github.com/v3nooonn/trytry-based-on-looklook/apps/customer/cmd/rpc/pb/customer"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/api/internal/svc"
	"github.com/v3nooonn/trytry-based-on-looklook/apps/oauth/cmd/api/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type AccessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessLogic {
	return &AccessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccessLogic) JWTGener(identifier string) (string, error) {
	pvtKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(l.svcCtx.Config.Auth.PrivateKey))
	if err != nil {
		return "", err
	}

	dur := time.Hour * 24 * time.Duration(l.svcCtx.Config.Auth.ExpirationDay)

	// claims := make(jwt.MapClaims)
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(dur).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "v3nooom",
			Id:        identifier,
		},
	)

	return token.SignedString(pvtKey)
}

func (l *AccessLogic) Access(req *types.AccessReq) (resp *types.AccessResp, err error) {
	c, err := l.svcCtx.CustomerRpc.InfoByEmail(l.ctx, &customer.InfoByEmailReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	token, err := l.JWTGener(c.Email)
	if err != nil {
		return nil, err
	}

	resp = &types.AccessResp{
		Response: types.Response{
			Code:    http.StatusOK,
			Message: "ok",
		},
		Data: types.Token{
			Access: token,
		},
	}

	return
}
