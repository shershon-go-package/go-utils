package src

import (
	"errors"
)

/**
 * ConvertStruct
 * @Description: 转换结构体
 * @Author: Shershon
 * @Param from
 * @Param to
 * @Return error
 * @Date 2023-07-03 17:41:43
 **/
func ConvertStruct(from interface{}, to interface{}) error {
	if to == nil {
		return errors.New("variable to can not be nil")
	}
	json := Json()
	toString, err := json.MarshalToString(from)
	if err != nil {
		return err
	}
	return json.UnmarshalFromString(toString, to)
}
