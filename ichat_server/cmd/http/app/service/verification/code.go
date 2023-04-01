package verification

import (
	"pim/app/utils/rand"
	"pim/app/utils/response"
	"time"
)

func CreateVerificationService() *VerificationService {
	return &VerificationService{}
}

type VerificationService struct {
	Code int
}

// 获取验证码
func (v *VerificationService) GetCode() *response.Response {
	//生成4位数字的验证码
	code := rand.GenCode()

	v.Code = code

	CreateVerificationCacheService().SetCode(code, time.Minute*2)

	return response.Ok.WithData(&v)
}
