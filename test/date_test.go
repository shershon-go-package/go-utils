/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/03 19:32
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

// 获取日期时间字符串的实际格式
func TestGetDateFormat(t *testing.T) {
	date := "2023-07-03"
	date2 := "2023-07-03 19:36:50.100"
	rst := src.GetDateFormat(date)
	rst2 := src.GetDateFormat(date2)
	fmt.Println("rst:", rst)
	fmt.Println("rst:", rst2)
}

// 时间相减
func TestDateSub(t *testing.T) {
	maxDate := "2023-07-03 18:30:03.555"
	minDate := "2023-07-03 17:20:02.111"
	rst, _ := src.DateSub(maxDate, minDate)
	fmt.Printf("rst: T:%T, v:%+v \n", rst, rst)
}
