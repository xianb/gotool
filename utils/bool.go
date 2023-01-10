package utils

func BoolToUInt(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func UIntToBool(u uint64) bool {
	if u == 1 {
		return true
	}
	return false
}
