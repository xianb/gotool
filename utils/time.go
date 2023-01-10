package utils

import (
	"fmt"
	"time"
)

func ToISO8601(t ...time.Time) string {
	var tt time.Time
	if len(t) > 0 {
		tt = t[0]
	} else {
		tt = time.Now().UTC()
	}

	var tz string
	name, offset := tt.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s", tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), tz)
}

//TimeFormat 格式化时间
//@param format 在go中，为2006-01-02 15:04:05
func TimeFormat(format string, t ...time.Time) string {
	var tt time.Time
	if len(t) > 0 {
		tt = t[0]
	} else {
		tt = time.Now()
	}
	s := tt.Format(format)
	return s
}
