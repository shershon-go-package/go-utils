package src

import "github.com/thoas/go-funk"

/**
 * FunkGetInt64
 * @Description: 获取结构体某字段值组成的切片
 * @Author: Shershon
 * @Param data
 * @Param key
 * @Return []int64
 * @Date 2023-07-03 17:40:05
 **/
func FunkGetInt64(data interface{}, key string) []int64 {
	var result []int64
	tmp := funk.Get(data, key)
	if tmp != nil {
		result = append(result, tmp.([]int64)...)
	}
	return result
}

/**
 * FunkGetString
 * @Description: 获取结构体某字段值组成的切片
 * @Author: Shershon
 * @Param data
 * @Param key
 * @Return []string
 * @Date 2023-07-17 20:10:02
 **/
func FunkGetString(data interface{}, key string) []string {
	var result []string
	tmp := funk.Get(data, key)
	if tmp != nil {
		result = append(result, tmp.([]string)...)
	}
	return result
}
