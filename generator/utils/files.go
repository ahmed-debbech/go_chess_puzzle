package utils

import (
	"os"
)

func IsDirectoryExist(dirt string) bool{
	if _, err := os.Stat(dirt); err != nil {
		return false
	}
	return true
}


func EndItWithSlash(path string) string{
	if path[len(path)-1] == '/' {
		return path
	}
	v := []byte(path)
	v = append(v, '/')
	return string(v)
}