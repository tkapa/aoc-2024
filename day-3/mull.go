package mull

import (
	"regexp"
	"strconv"
	"strings"
)

func MullGetMulTotalSimple(input string) int {

	muls := MullFindValidMuls(input)
	total := MullMultMuls(muls)

	return total
}

func MullGetMulTotalDoDont(input string) int {

	dont := "don't()"
	doSlices := make([]string, 0)
	sliceAtDo := strings.Split(input, "do()")

	for i := range sliceAtDo {

		if !strings.Contains(sliceAtDo[i], dont) {
			doSlices = append(doSlices, sliceAtDo[i])
			continue
		}

		doSlice := strings.Split(sliceAtDo[i], dont)[0]
		doSlices = append(doSlices, doSlice)
	}

	total := 0
	for i := range doSlices {
		total += MullGetMulTotalSimple(doSlices[i])
	}

	return total
}

func MullFindValidMuls(input string) [][]int {
	// New 2d slice to track a list of int pairs to multiply
	muls := make([][]int, 0)

	r, _ := regexp.Compile(`mul\((\d+,\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	for x := range matches {
		strPair := matches[x][1]
		numPair := strings.Split(strPair, ",")
		val1, _ := strconv.Atoi(numPair[0])
		val2, _ := strconv.Atoi(numPair[1])

		numArr := []int{val1, val2}

		muls = append(muls, numArr)
	}

	return muls
}

func MullMultMuls(muls [][]int) int {
	total := 0

	for x := range muls {
		total += muls[x][0] * muls[x][1]
	}

	return total
}
