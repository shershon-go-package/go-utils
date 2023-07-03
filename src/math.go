package src

import (
	"math"
	"strconv"
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
