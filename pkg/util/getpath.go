package util

import "os"

func GetPath() string {
	curDir, _ := os.Getwd()

	return curDir + "/"
}
