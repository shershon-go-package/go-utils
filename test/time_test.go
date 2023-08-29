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
	maxDate := "2023-07-17 18:58:44"
	minDate := "2023-07-13 14:55:26"
	rst, _ := src.DateSub(maxDate, minDate)   // 返回包含类似 1h10m10s 字段的结构体
	rst2, _ := src.DateSub2(maxDate, minDate) // 返回 XX时XX分XX秒 格式的字符串
	fmt.Printf("rst.Desc: T:%T, v:%+v \n", rst.Desc, rst.Desc)
	fmt.Printf("rst2: T:%T, v:%+v \n", rst2, rst2)
}

// 判断时间差是否满足 >= day
func TestDurationGEDay(t *testing.T) {
	maxDate := "2023-07-17 18:58:44"
	minDate := "2023-07-15 14:55:26"
	rst, _ := src.IsDiffDay(maxDate, minDate, 3)
	fmt.Printf("rst: T:%T, v:%+v \n", rst, rst)
}

// 字符串时间格式化
func TestFormatLocalStrDate(t *testing.T) {
	date := "2023-07-17 19:31:30"
	format := "2006-01-02 15:04"
	rst, _ := src.FormatLocalStrDate(date, format)
	fmt.Printf("rst: T:%T, v:%+v \n", rst, rst)

	date2 := "2023-07-17"
	format2 := "2006-01-02 15:04:05"
	rst2, _ := src.FormatLocalStrDate(date2, format2)
	fmt.Printf("rst2: T:%T, v:%+v \n", rst2, rst2)
}

// 输出日期之间的所有日期（按天切分日期）
func TestSplidDate(t *testing.T) {
	rst := src.SplitDate("2023-07-01", "2023-07-31", "2006-01-02")
	fmt.Println(rst)
}

// 比较两个时间字符串的大小
func TestCompareTime(t *testing.T) {
	time1 := "2023-09-01 00:00:00"
	time2 := "2023-09-01 12:00:00"
	rst, _ := src.CompareTime(time1, time2)
	if rst > 0 {
		fmt.Println("time1 > time2")
	} else if rst == 0 {
		fmt.Println("time1 = time2")
	} else {
		fmt.Println("time1 < time2")
	}

	time11 := "2023-09-02 00:00:00"
	time22 := "2023-09-01 12:00:00"
	rst2, _ := src.CompareTime(time11, time22)
	if rst2 > 0 {
		fmt.Println("time11 > time22")
	} else if rst2 == 0 {
		fmt.Println("time11 = time22")
	} else {
		fmt.Println("time11 < time22")
	}
}
