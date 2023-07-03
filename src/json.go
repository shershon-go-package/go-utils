package src

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/thoas/go-funk"
	"strings"
	"time"
)

var jsont = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	// 自适应类型
	extra.RegisterFuzzyDecoders()
}

func Json() jsoniter.API {
	return jsont
}

type DateTime time.Time
type Date time.Time

/**
 * MarshalJSON
 * @Description: 处理时间类型
 * @Author: Shershon
 * @Receiver d
 * @Return []byte
 * @Return error
 * @Date 2023-07-03 17:40:19
 **/
func (d DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02 15:04:05"))
	if strings.Contains(stamp, "000") {
		stamp = `""`
	}
	return []byte(stamp), nil
}

/**
 * UnmarshalJSON
 * @Description: 反序列化
 * @Author: Shershon
 * @Receiver d
 * @Param b
 * @Return error
 * @Date 2023-07-03 17:40:28
 **/
func (d *DateTime) UnmarshalJSON(b []byte) error {
	val := string(b)
	var t time.Time
	var err error
	if funk.ContainsString([]string{`""`, `null`, `"null"`}, val) || strings.Contains(val, "000") {
		t, _ = time.Parse(`2006-01-02 15:04:05`, "0000-00-00 00:00:00")
	} else {
		t, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	}
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

/**
 * MarshalJSON
 * @Description: 序列化
 * @Author: Shershon
 * @Receiver d
 * @Return []byte
 * @Return error
 * @Date 2023-07-03 17:40:34
 **/
func (d Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02"))
	if strings.Contains(stamp, "000") {
		stamp = `""`
	}
	return []byte(stamp), nil
}

/**
 * UnmarshalJSON
 * @Description: 反序列化
 * @Author: Shershon
 * @Receiver d
 * @Param b
 * @Return error
 * @Date 2023-07-03 17:41:00
 **/
func (d *Date) UnmarshalJSON(b []byte) error {
	val := string(b)
	var t time.Time
	var err error
	if funk.ContainsString([]string{`""`, `null`, `"null"`}, val) || strings.Contains(val, "000") {
		t, _ = time.Parse(`2006-01-02`, "0000-00-00")
	} else {
		t, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	}
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (d *Date) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = Date(t)
		return nil
	}
	// 转成字符串
	date2 := fmt.Sprintf("%s", src)
	if strings.Contains(date2, "000") {
		*d = Date(time.Time{})
		return nil
	}
	format := GetDateFormat(date2)
	location, err := time.ParseInLocation(format, date2, time.Local)
	if err == nil {
		*d = Date(location)
		return nil
	}
	return fmt.Errorf("can not convert %v to wmutil.Date", date2)
}

// 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (d *DateTime) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = DateTime(t)
		return nil
	}
	// 转成字符串
	// 转成字符串
	date2 := fmt.Sprintf("%s", src)
	if strings.Contains(date2, "000") {
		*d = DateTime(time.Time{})
		return nil
	}
	format := GetDateFormat(date2)
	location, err := time.ParseInLocation(format, date2, time.Local)
	if err == nil {
		*d = DateTime(location)
		return nil
	}
	return fmt.Errorf("can not convert %v to wmutil.DateTime", date2)
}
