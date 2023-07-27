/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/27 19:52
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 转换结构体
func TestConvertStruct(t *testing.T) {
	from := Animal{
		Name: "小狗",
		Age:  2,
	}
	var to Person
	_ = src.ConvertStruct(from, &to)
	fmt.Printf("%+v \n", to)
}
