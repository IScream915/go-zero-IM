package handler

import (
	"context"
	"go-zero-IM/im/ws/internal/svc"
	"go-zero-IM/pkg/ctxData"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
)

type Jwt struct {
	svc    *svc.ServiceContext
	parser *token.TokenParser
	logx.Logger
}

func NewJwt(svc *svc.ServiceContext) *Jwt {
	return &Jwt{
		svc:    svc,
		parser: token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}

func (j *Jwt) Auth(w http.ResponseWriter, r *http.Request) bool {
	tok, err := j.parser.ParseToken(r, j.svc.Config.Jwt.AccessSecret, "")
	if err != nil {
		j.Errorf("parse token err %v ", err)
		return false
	}

	if !tok.Valid {
		return false
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	*r = *r.WithContext(context.WithValue(r.Context(), ctxData.Identify, claims[ctxData.Identify]))

	return true
}

func (j *Jwt) UserId(r *http.Request) string {
	uid, err := ctxData.GetUid(r.Context())
	if err != nil {
		return ""
	}
	return uid
}
