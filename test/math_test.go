/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/03 17:42
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

// 保留小数位(四舍五入)
func TestRound(t *testing.T) {
	num := 1.534
	num2 := 1.235
	rst := src.Round(num, 0)
	rst2 := src.Round(num2, 2)
	fmt.Printf("rst T:%T, v:%v \n", rst, rst)
	fmt.Printf("rst2 T:%T, v:%v \n", rst2, rst2)
}

// 保留小数位(四舍五入)
func TestRoundString(t *testing.T) {
	num := 1.534
	num2 := 1.235
	rst := src.RoundString(num, 0)
	rst2 := src.RoundString(num2, 2)
	fmt.Printf("rst T:%T, v:%v \n", rst, rst)
	fmt.Printf("rst2 T:%T, v:%v \n", rst2, rst2)
}
