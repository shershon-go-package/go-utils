/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/17 19:55
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

type User struct {
	Name string
	Age  int64
	Home struct {
		City string
	}
}

// 获取结构体某字段值组成的切片
func TestFunkGetInt64(t *testing.T) {
	userList := []User{
		{
			Name: "张三",
			Age:  20,
			Home: struct {
				City string
			}{"北京"},
		},
		{
			Name: "小明",
			Age:  30,
			Home: struct {
				City string
			}{"南京"},
		},
	}

	// 取一层
	ages := src.FunkGetInt64(userList, "Age")
	fmt.Println("ages:", ages)
}

// 获取结构体某字段值组成的切片
func TestFunkGetString(t *testing.T) {
	userList := []User{
		{
			Name: "张三",
			Home: struct {
				City string
			}{"北京"},
		},
		{
			Name: "小明",
			Home: struct {
				City string
			}{"南京"},
		},
	}

	// 取一层
	names := src.FunkGetString(userList, "Name")
	fmt.Println("names:", names)
	// 取其他层
	homes := src.FunkGetString(userList, "Home.City")
	fmt.Println("homes:", homes)
}
