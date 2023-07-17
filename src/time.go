package src

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type CompareDateResult struct {
	Status      bool // true 为赋值成功
	Day         int
	Hour        int
	Minute      int
	Second      float64
	Desc        string
	TotalSecond float64
}

/**
 * GetDateFormat
 * @Description: 获取日期时间字符串的实际格式
 * @Author: Shershon
 * @Param date
 * @Return string
 * @Date 2023-07-03 17:35:40
 **/
func GetDateFormat(date string) string {
	if _, err := time.Parse("2006-01-02", date); err == nil {
		return "2006-01-02"
	} else if _, err := time.Parse("2006-01-02 15:04:05", date); err == nil {
		return "2006-01-02 15:04:05"
	} else if _, err := time.Parse("2006-01-02 15:04:05.000", date); err == nil {
		return "2006-01-02 15:04:05.000"
	}
	return ""
}

/**
 * DateSub
 * @Description: 时间相减
 * @Author: Shershon
 * @Param maxTime
 * @Param minTime
 * @Return CompareDateResult
 * @Return error
 * @Date 2023-07-03 17:35:46
 **/
func DateSub(maxTime, minTime string) (CompareDateResult, error) {
	var result CompareDateResult
	maxDate, minDate, err := parseCompareTime(maxTime, minTime)
	if err != nil {
		return result, err
	}
	// diff = 1h10m1s
	result.TotalSecond = maxDate.Sub(minDate).Seconds()
	diff := maxDate.Sub(minDate).String()
	fieldsFunc := strings.FieldsFunc(diff, func(r rune) bool {
		return unicode.IsLetter(r)
	})
	result.Desc = diff
	if len(fieldsFunc) == 3 {
		result.Status = true
		// 总小时
		totalHours, _ := strconv.Atoi(fieldsFunc[0])
		day := totalHours / 24
		// 天
		result.Day = day
		// 小时
		result.Hour = totalHours - (day * 24)
		// 分钟
		result.Minute, _ = strconv.Atoi(fieldsFunc[1])
		// 秒
		second, _ := strconv.ParseFloat(fieldsFunc[2], 64)
		result.Second = second
	}
	return result, nil
}

/**
 * DateSub2
 * @Description: 时间相减
 * @Author: Shershon
 * @Param maxTime
 * @Param minTime
 * @Return string XX时XX分XX秒
 * @Return error
 * @Date 2023-07-17 10:15:15
 **/
func DateSub2(maxTime, minTime string) (string, error) {
	// 将两个时间类型转换成 time.Time 类型
	maxDate, minDate, err := parseCompareTime(maxTime, minTime)
	if err != nil {
		return "", err
	}
	// 计算两个时间相差的时分秒
	diff := maxDate.Sub(minDate)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) - hours*60
	seconds := int(diff.Seconds()) - hours*3600 - minutes*60

	return fmt.Sprintf("%d小时%d分%d秒", hours, minutes, seconds), nil
}

/**
 * IsDiffDay
 * @Description: 判断时间差是否满足 >= day
 * @Author: Shershon
 * @Param max
 * @Param min
 * @Param day
 * @Return bool
 * @Return error
 * @Date 2023-07-03 17:35:55
 **/
func IsDiffDay(max string, min string, day int) (bool, error) {
	maxDate, minDate, err := parseCompareTime(max, min)
	if err != nil {
		return false, err
	}
	hours := maxDate.Sub(minDate).Hours()
	if hours/float64(24) >= float64(day) {
		return true, nil
	}
	return false, nil
}

/**
 * parseCompareTime
 * @Description: 解析时间
 * @Author: Shershon
 * @Param max
 * @Param min
 * @Return time.Time
 * @Return time.Time
 * @Return error
 * @Date 2023-07-03 17:36:01
 **/
func parseCompareTime(max string, min string) (time.Time, time.Time, error) {
	if len(max) != len(min) {
		return time.Time{}, time.Time{}, errors.New("两个时间格式需要一致")
	}
	format := GetDateFormat(max)
	minDate, err1 := time.ParseInLocation(format, min, time.Local)
	maxDate, err2 := time.ParseInLocation(format, max, time.Local)
	if err1 != nil || err2 != nil {
		return time.Time{}, time.Time{}, errors.New("时间解析失败")
	}
	if maxDate.After(minDate) {
		return maxDate, minDate, nil
	}
	return minDate, maxDate, nil
}

/**
 * FormatLocalStrDate
 * @Description: 字符串时间格式化
 * @Author: Shershon
 * @Param date
 * @Param format
 * @Return string
 * @Return error
 * @Date 2023-07-03 17:36:13
 **/
func FormatLocalStrDate(date string, format string) (string, error) {
	dateFormat := GetDateFormat(date)
	location, err := time.ParseInLocation(dateFormat, date, time.Local)
	if err != nil {
		return "", err
	}
	retDate := location.Format(format)
	return retDate, nil
}
