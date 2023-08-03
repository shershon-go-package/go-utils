/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/08/03 15:37
 */

package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 生成指定位数的随机字母
func TestRandomStr(t *testing.T) {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	digit := 5
	for i := 0; i < 10; i++ {
		bytes := make([]byte, digit)
		for i := 0; i < digit; i++ {
			b := rand.Intn(26) + 65
			bytes[i] = byte(b)
		}
		fmt.Println(string(bytes))
	}
}
