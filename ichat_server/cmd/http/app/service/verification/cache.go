package verification

import (
	"context"
	"fmt"
	"pim/app/model/global"
	"time"
)

func CreateVerificationCacheService() *VerificationCacheService {
	return &VerificationCacheService{}
}

type VerificationCacheService struct {
}

var ctx = context.Background()

// 设置验证码
func (v *VerificationCacheService) SetCode(code int, expire time.Duration) bool {
	global.Redis.Set(ctx, fmt.Sprintf("verifcation:code:%d", code), code, expire)

	return true
}
