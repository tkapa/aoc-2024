package mull

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MullReadFile(file string) string {
	f, err := os.ReadFile(file)
	check(err)
	retStr := string(f)

	return retStr
}
