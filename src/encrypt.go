package src

import (
	"crypto/md5"
	"fmt"
)

/**
 * GetMd5
 * @Description: 生成MD5密码
 * @Author: Shershon
 * @Param str
 * @Return string
 * @Date 2023-07-03 17:36:21
 **/
func GetMd5(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
