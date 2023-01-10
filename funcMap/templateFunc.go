package funcMap

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"
)

// ToInt64 类型转换，获得int64
func ToInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		re, err = strconv.ParseInt(v.(string), 10, 64)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	case int64:
		re = v.(int64)
	case int32:
		re = v.(int64)
	case int:
		re = int64(v.(int))
	default:
		err = fmt.Errorf("不能转换")
	}
	return
}

func IndexAddOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index + 1
}

func IndexDecrOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index - 1
}

func TemplateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

var FuncMap = template.FuncMap{
	"IndexAddOne":  IndexAddOne,
	"IndexDecrOne": IndexDecrOne,
	"TemplateTime": TemplateTime,
	"StrToUpper":   StrToUpper,
}
