package random

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	length := 10

	p := NewRandom(length, true, false)
	for i := 0; i < 100; i++ {
		fmt.Println("-->>", p.New())

	}
}
