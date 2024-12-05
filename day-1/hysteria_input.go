package hysteria

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func HysteriaReadFile(file string) ([]int, []int) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		left, err := strconv.Atoi(numbers[0])
		check(err)
		list1 = append(list1, left)

		right, err := strconv.Atoi(numbers[1])
		check(err)
		list2 = append(list2, right)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return list1, list2
}
