package utils

import (
	"fmt"
	"math/rand"
)

// GenRandFourDigits 生成随机的 四位非负整数
func GenRandFourDigits() string {
	num := rand.Intn(10000)
	return fmt.Sprintf("%04d", num)
}
