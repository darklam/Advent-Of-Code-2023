package day_1

import (
	"strings"

	"golang.org/x/exp/maps"

	"advent-of-code-2023/util"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitsArray = maps.Keys(digits)

func getDigitSpelled(s string) int {
	for _, digit := range digitsArray {
		if strings.Contains(s, digit) {
			return digits[digit]
		}
	}
	return -1
}

func getDigit(char uint8) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	} else {
		return -1
	}
}

func SumCalibrationValues(reader util.FileReader) int {
	data, err := reader("./inputs/day_1.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range data {
		l := 0
		r := len(line) - 1

		leftDigit := -1
		rightDigit := -1

		for leftDigit == -1 {
			leftDigit = getDigit(line[l])
			if leftDigit == -1 {
				leftDigit = getDigitSpelled(line[:l+1])
			}
			l++
		}

		for rightDigit == -1 {
			rightDigit = getDigit(line[r])
			if rightDigit == -1 {
				rightDigit = getDigitSpelled(line[r:])
			}
			r--
		}

		sum += leftDigit*10 + rightDigit
	}

	return sum
}
