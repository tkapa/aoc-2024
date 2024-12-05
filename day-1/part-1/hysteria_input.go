package hysteria

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner.Split(bufio.ScanWords)
	i := 1
	for scanner.Scan() {
		str := scanner.Text()
		num, err := strconv.Atoi(str)
		check(err)
		fmt.Println(scanner.Text())

		if i%2 == 0 {
			list2 = append(list2, num)
		} else {
			list1 = append(list1, num)
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return list1, list2
}
