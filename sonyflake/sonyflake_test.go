package sonyflake

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("-->>", GenUUid())
}
