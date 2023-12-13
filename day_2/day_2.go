package day_2

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code-2023/util"
)

var counts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameId(s string) int {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		panic("didn't parse the game correctly")
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return id
}

func removeWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func isGameValid(rounds []string) bool {
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			// color[1:] because we don't want the starting whitespace
			parts := strings.Split(color[1:], " ")
			numString, colorName := parts[0], removeWhitespace(parts[1])
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}

			if num > counts[colorName] {
				return false
			}
		}
	}

	return true
}

func getMinimumPower(rounds []string) int {
	colorsMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			// color[1:] because we don't want the starting whitespace
			parts := strings.Split(color[1:], " ")
			numString, colorName := parts[0], removeWhitespace(parts[1])
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}

			currentValue := colorsMap[colorName]
			if num > currentValue {
				colorsMap[colorName] = num
			}
		}
	}

	product := 1
	fmt.Println("Rounds = ", rounds, " Minimums: ", colorsMap)
	for _, value := range colorsMap {
		product *= value
	}

	return product
}

func iterateGames(reader util.FileReader, cb func(gameId int, rounds []string)) {
	lines, err := reader("./inputs/day_2.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("didn't parse the game correctly")
		}

		game, data := parts[0], parts[1]
		id := getGameId(game)

		rounds := strings.Split(data, ";")
		if len(rounds) == 0 {
			continue
		}

		cb(id, rounds)
	}
}

// func GetSumOfGameIds(reader util.FileReader) int {
// 	lines, err := reader("./inputs/day_2.txt")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	sum := 0
//
// 	for _, line := range lines {
// 		parts := strings.Split(line, ":")
// 		if len(parts) != 2 {
// 			panic("didn't parse the game correctly")
// 		}
//
// 		game, data := parts[0], parts[1]
// 		id := getGameId(game)
//
// 		rounds := strings.Split(data, ";")
// 		if len(rounds) == 0 {
// 			continue
// 		}
//
// 		if isGameValid(rounds) {
// 			sum += id
// 		}
// 	}
//
// 	return sum
// }

func GetSumOfGameIds(reader util.FileReader) int {
	sum := 0
	iterateGames(reader, func(id int, rounds []string) {
		if isGameValid(rounds) {
			sum += id
		}
	})

	return sum
}

func GetSumOfMinimumPowers(reader util.FileReader) int {
	sum := 0
	iterateGames(reader, func(id int, rounds []string) {
		sum += getMinimumPower(rounds)
	})

	return sum
}
