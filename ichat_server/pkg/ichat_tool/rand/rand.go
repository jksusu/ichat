package rand

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"ichat/pkg/ichat_tool/md5"
	"math/rand"
	"time"
)

// GenSlot 生成密码盐6位长度
func GenSlot() string {
	number := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000)
	return fmt.Sprintf("%d%s", number, "a")
}

func GenUid() string {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	return fmt.Sprintf("%d", n)
}

// GenToken 生成token
func GenToken() string {
	return md5.MD5(uuid.NewV1().String())
}

// GenCode 四位验证码
func GenCode() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
}
