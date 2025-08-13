package ctxData

import (
	"context"
	"errors"
)

func GetUid(ctx context.Context) (string, error) {
	if uid, ok := ctx.Value(Identify).(string); ok {
		return uid, nil
	} else {
		return "", errors.New("无法从ctx中获取用户的uid信息")
	}
}
