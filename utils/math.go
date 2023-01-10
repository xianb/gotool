package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

func StringNumToBigIntWithExp(amount string, exp int32) *big.Int {
	vDecimal, _ := decimal.NewFromString(amount)
	vDecimal = vDecimal.Shift(exp)
	bigInt, ok := new(big.Int).SetString(vDecimal.String(), 10)
	if !ok {
		return big.NewInt(0)
	}
	return bigInt
}

func IntToDecimals(amount int64, decimals int32) decimal.Decimal {
	return decimal.New(amount, 0).Shift(-decimals)
}

func BigIntToDecimals(amount *big.Int, decimals int32) decimal.Decimal {
	if amount == nil {
		return decimal.Zero
	}
	return decimal.NewFromBigInt(amount, 0).Shift(-decimals)
}

func StringValueToBigInt(value string, base int) (*big.Int, error) {
	value = strings.TrimPrefix(value, "0x")
	if value == "" || value == "0" {
		return big.NewInt(0), nil
	}

	bigValue, success := new(big.Int).SetString(value, base)
	if !success {
		return big.NewInt(0), fmt.Errorf("convert value [%v] to bigint failed, check the value and base passed through", value)
	}
	return bigValue, nil
}

func Str2BigInt(value string) *big.Int {
	bigvalue := new(big.Int)
	_, success := bigvalue.SetString(value, 10)
	if !success {
		return big.NewInt(0)
	}
	return bigvalue
}

func Str2BigIntShift(value string, shift int32) *big.Int {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return big.NewInt(0)
	}
	return v.Shift(shift).BigInt()
}

func BytesToDecimals(bits []byte, decimals int32) decimal.Decimal {
	if bits == nil {
		return decimal.Zero
	}
	amount := new(big.Int)
	amount.SetBytes(bits)
	return decimal.NewFromBigInt(amount, 0).Shift(-decimals)
}
