/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/17 19:45
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

// 判断文件是否存在
func TestExistFile(t *testing.T) {
	path := "./math_test.go"
	rst := src.ExistFile(path)
	fmt.Printf("rst -> T:%T, v:%v \n", rst, rst)
}

// 判断目录是否存在：存在则返回，不存在则创建目录
func TestCreateDir(t *testing.T) {
	path := "./test2"
	rst := src.CreateDir(path)
	fmt.Printf("rst -> T:%T, v:%v \n", rst, rst)
}
