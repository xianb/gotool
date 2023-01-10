package utils

import (
	"testing"
)

func TestStringNumToBigIntWithDecimals(t *testing.T) {
	num := StringNumToBigIntWithExp("3.223", 8)
	t.Logf("num : %v\n", num)
	t.Logf("num int64 : %v\n", num.Int64())
	newnum := BigIntToDecimals(num, 8)
	t.Logf("newnum : %v\n", newnum.String())
}
