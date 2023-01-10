package sonyflake

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var serial *sonyflake.Sonyflake

func init() {
	serial = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenUUid() string {
	sid, err := serial.NextID()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", sid)
}
