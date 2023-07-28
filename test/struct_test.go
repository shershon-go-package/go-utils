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
	from := make([]Animal, 0)
	from = append(from, Animal{
		Name: "小狗",
		Age:  2,
	})
	to := make([]Person, 0)
	_ = src.ConvertStruct(from, &to)
	fmt.Printf("T:%T, v:%+v \n", to, to)
}
