package logic

import (
	"context"
	"errors"
	"go-zero-IM/social/dao/models"
	"go-zero-IM/user/dao/query/implement"
	"time"

	"go-zero-IM/social/rpc/internal/svc"
	"go-zero-IM/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	dbTool := implement.NewDbToolHelper[models.FriendRequests](l.svcCtx.DB)
	// 获取好友申请
	record, err := dbTool.SearchSingleByField(l.ctx, "id", in.FriendReqId)
	if err != nil {
		return nil, errors.New("查询" + models.FriendRequests{}.TableName() + "时出现了问题" + err.Error())
	}
	if record == nil {
		return nil, errors.New(models.FriendRequests{}.TableName() + "查询结果为空")
	}

	// 判断是否需要更新
	if record.HandleResult != models.ResultUnhandled {
		// 当前好友请求状态已经被处理, 则不能再进行处理
		return nil, errors.New("当前好友请求已处理")
	}

	// 更新t_friend_requests表 和 t_friends表
	if txErr := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		txFRTool := implement.NewDbToolHelper[models.FriendRequests](tx)
		// 更新t_friend_requests表中的状态
		if err = txFRTool.UpdateOneOrMultiFields(l.ctx, "id", in.FriendReqId, map[string]interface{}{
			"handle_result": in.HandleResult,
			"handled_at":    time.Now(),
		}); err != nil {
			return err
		}

		// 当同意好友申请的时候进行好友关系生成
		if in.HandleResult == models.ResultApprove {
			// 在t_friends表中形成好友关系
			txFTool := implement.NewDbToolHelper[models.Friends](tx)
			if err = txFTool.InsertBatchRecords(l.ctx, []models.Friends{
				{
					UserID:    record.UserID,
					FriendUID: record.ReqUID,
				},
				{
					UserID:    record.ReqUID,
					FriendUID: record.UserID,
				},
			}); err != nil {
				return err
			}
		}
		return nil

	}); txErr != nil {
		return nil, errors.New("处理好友申请时出现问题" + txErr.Error())
	}

	return &social.FriendPutInHandleResp{}, nil
}
