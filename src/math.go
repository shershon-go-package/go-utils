package src

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

/**
 * Round
 * @Description: 保留小数位(四舍五入)
 * @Author: Shershon
 * @Param num
 * @Param decimal
 * @Return float64
 * @Date 2023-07-03 19:10:27
 **/
func Round(num float64, decimal int) float64 {
	if decimal == 0 {
		return math.Round(num)
	}
	// 转10的次方
	pow := math.Pow10(decimal)
	return math.Floor(num*pow+0.5) / pow
}

/**
 * RoundString
 * @Description: 保留小数位(四舍五入)
 * @Author: Shershon
 * @Param num
 * @Param decimal
 * @Return string
 * @Date 2023-07-03 17:41:18
 **/
func RoundString(num float64, decimal int) string {
	f := Round(num, decimal)
	return strconv.FormatFloat(f, 'f', decimal, 64)
}

/**
 * GenerateFixedDigitRandom
 * @Description: 生成固定位数的随机数
 * @Author: Shershon
 * @Param digits
 * @Return int64
 * @Date 2023-08-03 11:29:21
 **/
func GenerateFixedDigitRandom(digits int) int64 {
	//设置随机数种子为当前时间的纳秒数
	rand.Seed(time.Now().UnixNano())

	max := int64(1)
	min := int64(1)
	for i := 0; i < digits; i++ {
		max *= 10
	}
	for i := 0; i < digits-1; i++ {
		min *= 10
	}

	return rand.Int63n(max-min) + min
}
