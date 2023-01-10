package utils

import "strconv"

/***************** 数值类型转字符串 *****************/

//Int 强化int类型
type Int int

//IntProtocal 强化int的扩展方法
type IntProtocal interface {
	Int(def ...int) int
	Int8(def ...int8) int8
	Int16(def ...int16) int16
	Int32(def ...int32) int32
	Int64(def ...int64) int64
}

// ToString int转为string类型
func (v Int) String() string {
	return strconv.FormatInt(int64(v), 10)
}
