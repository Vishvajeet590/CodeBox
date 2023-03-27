package utils

import "os"

func IsFilePresent(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}
