package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenRandFourDigits 生成随机的 四位非负整数
func GenRandFourDigits() func() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prev := r.Intn(10000)
	return func() string {
		num := r.Intn(10000)
		for prev == num {
			num = r.Intn(10000)
		}
		prev = num
		return fmt.Sprintf("%04d", num)
	}
}
