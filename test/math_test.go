/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/03 17:42
 */

package test

import (
	"fmt"
	"math/rand"
	"shershon1991/go-utils/src"
	"testing"
	"time"
)

// 保留小数位(四舍五入)
func TestRound(t *testing.T) {
	num := 1.534
	num2 := 1.235
	rst := src.Round(num, 0)
	rst2 := src.Round(num2, 2)
	fmt.Printf("rst T:%T, v:%v \n", rst, rst)
	fmt.Printf("rst2 T:%T, v:%v \n", rst2, rst2)
}

// 保留小数位(四舍五入)
func TestRoundString(t *testing.T) {
	num := 1.534
	num2 := 1.235
	rst := src.RoundString(num, 0)
	rst2 := src.RoundString(num2, 2)
	fmt.Printf("rst T:%T, v:%v \n", rst, rst)
	fmt.Printf("rst2 T:%T, v:%v \n", rst2, rst2)
}

// 生成固定位数的随机数
func TestGenerateFixedDigitRandom(t *testing.T) {
	//设置随机数种子为当前时间的纳秒数
	rand.Seed(time.Now().UnixNano())

	max := int64(1)
	min := int64(1)
	digits := 6
	for i := 0; i < digits; i++ {
		max *= 10
	}
	for i := 0; i < digits-1; i++ {
		min *= 10
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("rst T:%T, v:%v \n", rand.Int63n(max-min)+min, rand.Int63n(max-min)+min)
	}
}
