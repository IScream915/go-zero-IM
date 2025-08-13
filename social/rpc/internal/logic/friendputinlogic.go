package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-IM/social/dao/models"
	"go-zero-IM/social/dao/query/implement"
	"time"

	"go-zero-IM/social/rpc/internal/svc"
	"go-zero-IM/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FriendPutIn 好友业务：请求好友
func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// 判断与目标是否已经是好友
	friendsTable := models.Friends{}.TableName()
	record := &models.Friends{}

	var count int64
	err := l.svcCtx.DB.
		WithContext(l.ctx).
		Table(fmt.Sprintf("%s AS f", friendsTable)).
		Where(fmt.Sprintf("f.`user_id` = '%s' AND f.`friend_uid` = '%s'", in.UserId, in.ReqUid)).
		First(record).
		Count(&count).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 在查询的时候出现了结果为空之外的问题
			return nil, errors.New("查询" + friendsTable + "时出现了问题")
		}
	}
	if count > 0 {
		// 当前与目标已经是好友，无需再添加好友
		return nil, errors.New("无法向已添加过的好友发送好友申请")
	}

	// 是否已经有过未通过的申请记录
	friendReqTable := models.FriendRequests{}.TableName()
	result := &models.FriendRequests{}
	// select *
	// 	from friendReqTable AS fr
	// 	where fr.`user_id` = x
	// 	and fr.`req_id` = y
	err = l.svcCtx.DB.
		WithContext(l.ctx).
		Table(fmt.Sprintf("%s AS fr", friendReqTable)).
		Where(fmt.Sprintf("fr.`user_id` = '%s' AND fr.`req_uid` = '%s'", in.UserId, in.ReqUid)).
		First(result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("查询" + friendReqTable + "时出现了问题" + err.Error())
		}
	} else {
		// 如果查询结果不为空并且处理结果为通过的时候返回错误
		if result.HandleResult == models.ResultUnhandled {
			return nil, errors.New("您已发送过好友申请, 请稍作等待")
		}
	}

	// 创建申请记录
	if txErr := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		txDbTool := implement.NewDbToolHelper[models.FriendRequests](tx)
		if err = txDbTool.InsertSingleRecord(l.ctx, &models.FriendRequests{
			UserID:  in.UserId,
			ReqUID:  in.ReqUid,
			ReqMsg:  in.ReqMsg,
			ReqTime: time.Now(),
		}); err != nil {
			return err
		}
		return nil
	}); txErr != nil {
		return nil, errors.New("创建好友申请时出现了问题" + txErr.Error())
	}

	return &social.FriendPutInResp{}, nil
}
