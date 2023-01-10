package utils

import "strings"

func Check0xIndex(addr string) string {
	if strings.Index(addr, "0x") == -1 {
		return "0x" + addr
	}
	return addr
}

func RemoveOxFromHex(value string) string {
	result := value
	if strings.Index(value, "0x") != -1 {
		result = Substr(value, 2, len(value))
	}
	return result
}
