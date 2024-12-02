package core

import (
	"os"
	"strconv"
)

func ParseInt(number string) int {
	result, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return result
}

func GetEnvText(name string) string {
	return os.Getenv(name)
}

func GetEnvInt(name string) int {
	return ParseInt(os.Getenv(name))
}
