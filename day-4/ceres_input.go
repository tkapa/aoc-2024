package ceres

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CeresReadFile(file string) [][]string {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	horizontalInputList := make([][]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var text = scanner.Text()
		horizontalInputList = append(horizontalInputList, strings.Split(text, ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return horizontalInputList
}
