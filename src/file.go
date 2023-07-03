package src

import "os"

/**
 * ExistFile
 * @Description: 判断文件是否存在
 * @Author: Shershon
 * @Param path
 * @Return bool
 * @Date 2023-07-03 17:36:27
 **/
func ExistFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
 * ExistDir
 * @Description: 存在则返回，不存在在创建
 * @Author: Shershon
 * @Param path
 * @Return bool
 * @Date 2023-07-03 17:36:34
 **/
func ExistDir(path string) bool {
	return ExistFile(path)
}

/**
 * CreateDir
 * @Description: 创建目录
 * @Author: Shershon
 * @Param path
 * @Return error
 * @Date 2023-07-03 17:37:08
 **/
func CreateDir(path string) error {
	dirExist := ExistDir(path)
	var err error
	if !dirExist {
		err = os.Mkdir(path, os.ModePerm)
	}
	return err
}
