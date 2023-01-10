package utils

import (
	"os"
	"strconv"
)

func GetEnvBool(name string) bool {
	v := os.Getenv(name)
	if v == "true" || v == "1" {
		return true
	}
	return false
}

func GetEnvString(name string, def ...string) string {
	value := os.Getenv(name)
	if value == "" && len(def) > 0 {
		return def[0]
	}
	return value
}

func GetEnvInt(name string, def ...int) int {
	value := os.Getenv(name)
	if value == "" && len(def) > 0 {
		return def[0]
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}
