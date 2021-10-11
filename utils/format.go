package utils

import (
	"github.com/Summrry-top/gin-glory-api/global"
	"strconv"

	// "math"
	"strings"
	"time"
)

// 格式化时间 2006-01-02 15:04:05
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 计算总叶数
func TotalPage(total int64, pageNum int64) int64 {
	n := total / pageNum
	if total%pageNum > 0 {
		n += 1
	}
	return n
	// return int64(math.Ceil(float64(total) / float64(pageNum)))
}

// id拼接为字符串 1 => %|1|%
func Join(id int) string {
	return strings.Join(global.Array, strconv.Itoa(id))
}

// 字符串数组拼接为字符串 [1,2,3] => |1|2|3|
func Joins(array []string) string {
	return strings.Join(global.Array, strings.Join(array, global.VerticalBar))
}

// !(sortId,tagId)字符串不能为空  字符串转字符串数组 |1|2|3| => [1,2,3]
func ToArray(str string) []string {
	str = str[1 : len(str)-1]
	return strings.Split(str, global.VerticalBar)
}
