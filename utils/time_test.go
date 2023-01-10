package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToTimestamp(t *testing.T) {

	count := 60
	year := 2018
	month := 2

	for i := 0; i < count; i++ {
		t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		timestamp := t.Unix()
		fmt.Println(timestamp)

		month++
		if month > 12 {
			month = 1
			year++
		}
	}
}
