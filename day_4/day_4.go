package day_4

import (
	"strings"

	"advent-of-code-2023/util"
)

func getNumber(char uint8) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	} else {
		return -1
	}
}

func parseNumbers(numsString string) ([]int, []int, error) {
	var winning []int
	var numbers []int

	parsingWinning := true

	currentNum := -1

	for i := 0; i < len(numsString); i++ {
		char := numsString[i]
		if char == '|' {
			parsingWinning = false
			continue
		} else if char == ' ' && currentNum != -1 {
			if parsingWinning {
				winning = append(winning, currentNum)
			} else {
				numbers = append(numbers, currentNum)
			}

			currentNum = -1
		} else if char != ' ' {
			number := getNumber(char)
			if currentNum == -1 {
				currentNum = number
			} else {
				currentNum *= 10
				currentNum += number
			}
		}
	}

	if currentNum != -1 {
		if parsingWinning {
			winning = append(winning, currentNum)
		} else {
			numbers = append(numbers, currentNum)
		}
	}

	return winning, numbers, nil
}

func getNumbers(lines []string) ([][]int, [][]int, error) {
	resultWinning := make([][]int, 0, len(lines))
	resultNumbers := make([][]int, 0, len(lines))

	for _, line := range lines {
		numbersPart := strings.Split(line, ":")[1]

		// it's arr[1:] because the first character is always a space
		winningNumbers, numbers, err := parseNumbers(numbersPart)
		if err != nil {
			return nil, nil, err
		}

		resultWinning = append(resultWinning, winningNumbers)
		resultNumbers = append(resultNumbers, numbers)
	}

	return resultWinning, resultNumbers, nil
}

func GetTotalPoints(reader util.FileReader) int {
	lines, err := reader("./inputs/day_4.txt")
	if err != nil {
		panic(err)
	}

	points := 0

	winning, numbers, err := getNumbers(lines)
	if err != nil {
		panic(err)
	}

	for i, nums := range numbers {
		winningMap := map[int]bool{}

		for _, w := range winning[i] {
			winningMap[w] = true
		}

		currentPoints := 0

		for _, n := range nums {
			_, isWinning := winningMap[n]

			if isWinning {
				if currentPoints == 0 {
					currentPoints = 1
				} else {
					currentPoints *= 2
				}
			}
		}

		points += currentPoints
	}

	return points
}

func GetScratchcardCount(reader util.FileReader) int {
	lines, err := reader("./inputs/day_4.txt")
	if err != nil {
		panic(err)
	}

	cardsCount := map[int]int{}

	winning, numbers, err := getNumbers(lines)

	for i := range numbers {
		cardsCount[i+1] = 0
	}

	for i, nums := range numbers {
		// +1 because we are at this card, so we have it at least once
		cardCount := cardsCount[i+1] + 1
		w := winning[i]

		winningMap := map[int]bool{}

		for _, winningNum := range w {
			winningMap[winningNum] = true
		}

		winCount := 0

		for _, num := range nums {
			_, wins := winningMap[num]
			if wins {
				winCount++
			}
		}

		if winCount != 0 {
			endIndex := winCount + i
			if endIndex > len(numbers)-1 {
				endIndex = len(numbers) - 1
			}

			for j := i + 1; j <= endIndex; j++ {
				currentCount := cardsCount[j+1]
				currentCount += cardCount

				cardsCount[j+1] = currentCount
			}
		}
	}

	// We haven't counted the original cards yet so we start from len(numbers)
	sum := len(numbers)

	for _, count := range cardsCount {
		sum += count
	}

	return sum
}
