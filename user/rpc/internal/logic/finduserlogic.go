package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"user/dao/models"
	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrSQL = errors.New("SQL语句查询错误")
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var results *[]*user.UserEntity

	// 构造查询子句
	condition := make([]string, 0)
	if in.Phone != "" {
		condition = append(condition, fmt.Sprintf("u.`phone` = '%s'", in.Phone))
	}
	if in.Nickname != "" {
		condition = append(condition, fmt.Sprintf("u.`nickname` = '%s'", in.Nickname))
	}
	if len(in.Ids) > 0 {
		ids := make([]string, len(in.Ids))
		for i, id := range in.Ids {
			ids[i] = fmt.Sprintf("'%s'", id)
		}
		condition = append(condition, fmt.Sprintf("u.`id` IN (%s)", strings.Join(ids, ",")))
	}

	userTable := models.User{}.TableName()
	if err := l.svcCtx.DB.
		Table(fmt.Sprintf("%s AS u", userTable)).
		WithContext(l.ctx).
		Select(fmt.Sprintf("u.`id`, u.`nickname`, u.`phone`, u.`status`, u.`sex`")).
		Where(strings.Join(condition, " AND ")).
		Scan(&results).Error; err != nil {
		return nil, ErrSQL
	}

	return &user.FindUserResp{
		User: *results,
	}, nil
}
