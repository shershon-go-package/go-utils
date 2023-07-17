/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/07/17 19:41
 */

package test

import (
	"fmt"
	"shershon1991/go-utils/src"
	"testing"
)

// md5加密
func TestGetMd5(t *testing.T) {
	rst := src.GetMd5("hello world!")
	fmt.Println("rst: ", rst)
}
