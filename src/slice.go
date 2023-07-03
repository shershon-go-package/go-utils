package src

import (
	"context"
	"github.com/cockroachdb/errors"
	"strconv"
	"strings"
)

/**
 * SliceInt64Join
 * @Description: int64切片拼接
 * @Author: Shershon
 * @Param ctx
 * @Param data
 * @Param sep
 * @Return string
 * @Date 2023-07-03 17:41:25
 **/
func SliceInt64Join(ctx context.Context, data []int64, sep string) string {
	var strSlice []string
	for _, datum := range data {
		strSlice = append(strSlice, strconv.FormatInt(datum, 10))
	}
	return strings.Join(strSlice, sep)
}

/**
 * SliceIntJoin
 * @Description: int切片拼接
 * @Author: Shershon
 * @Param ctx
 * @Param data
 * @Param sep
 * @Return string
 * @Date 2023-07-03 17:41:30
 **/
func SliceIntJoin(ctx context.Context, data []int, sep string) string {
	var strSlice []string
	for _, datum := range data {
		strSlice = append(strSlice, strconv.Itoa(datum))
	}
	return strings.Join(strSlice, sep)
}

/**
 * Split2Int64
 * @Description: 切割字符串，转成[]int64
 * @Author: Shershon
 * @Param ctx
 * @Param data
 * @Param sep
 * @Return []int64
 * @Return error
 * @Date 2023-07-03 17:41:36
 **/
func Split2Int64(ctx context.Context, data, sep string) ([]int64, error) {
	var result []int64
	split := strings.Split(data, sep)
	if len(split) == 0 {
		return result, errors.New("Split error")
	}
	for _, s := range split {
		parseInt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, parseInt)
	}
	return result, nil
}
