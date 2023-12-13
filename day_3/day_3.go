package day_3

import (
	"fmt"
	"strconv"

	"advent-of-code-2023/util"
)

func getKey(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func isSymbol(c uint8) bool {
	return !(isNumeric(c) || c == '.')
}

func isNumeric(c uint8) bool {
	return c >= '0' && c <= '9'
}

type numberLocation struct {
	row        int
	startIndex int
	endIndex   int
	number     int
}

func getNumberLocations(lines []string) []numberLocation {
	var numberLocations []numberLocation
	for i, line := range lines {
		startIndex := -1
		for j := range line {
			char := line[j]

			if isNumeric(char) {
				if startIndex == -1 {
					startIndex = j
				}
				continue
			}

			if startIndex != -1 {
				num, err := strconv.Atoi(line[startIndex:j])
				if err != nil {
					panic(err)
				}

				numberLocations = append(numberLocations, numberLocation{
					row:        i,
					startIndex: startIndex,
					endIndex:   j - 1,
					number:     num,
				})

				startIndex = -1
			}

		}

		if startIndex != -1 {
			// We have a number that ends at the end of the line
			num, err := strconv.Atoi(line[startIndex:])
			if err != nil {
				panic(err)
			}
			numberLocations = append(numberLocations, numberLocation{
				row:        i,
				startIndex: startIndex,
				endIndex:   len(line) - 1,
				number:     num,
			})
		}
	}

	return numberLocations
}

func SumPartNumbers(reader util.FileReader) int {
	lines, err := reader("./inputs/day_3.txt")
	if err != nil {
		panic(err)
	}

	numberLocations := getNumberLocations(lines)

	sum := 0

	for _, location := range numberLocations {
		start, end := location.startIndex, location.endIndex
		rowStart, rowEnd := location.row, location.row
		if start != 0 {
			start--
		}

		if end != len(lines[0])-1 {
			end++
		}

		if rowStart != 0 {
			rowStart--
		}

		if rowEnd != len(lines)-1 {
			rowEnd++
		}

		for i := rowStart; i <= rowEnd; i++ {
			for j := start; j <= end; j++ {
				if isSymbol(lines[i][j]) {
					fmt.Println("Adding ", location.number)
					sum += location.number
				}
			}
		}
	}

	return sum
}

func GetGearRatiosSum(reader util.FileReader) int {
	lines, err := reader("./inputs/day_3.txt")
	if err != nil {
		panic(err)
	}

	numberLocations := getNumberLocations(lines)

	starNumbers := map[string][]numberLocation{}

	for _, location := range numberLocations {
		start, end := location.startIndex, location.endIndex
		rowStart, rowEnd := location.row, location.row
		if start != 0 {
			start--
		}

		if end != len(lines[0])-1 {
			end++
		}

		if rowStart != 0 {
			rowStart--
		}

		if rowEnd != len(lines)-1 {
			rowEnd++
		}

		for i := rowStart; i <= rowEnd; i++ {
			for j := start; j <= end; j++ {
				char := lines[i][j]

				if char == '*' {
					key := fmt.Sprintf("%d-%d", i, j)
					currentLocations, exists := starNumbers[key]
					if !exists {
						starNumbers[key] = []numberLocation{location}
					} else {
						currentLocations = append(currentLocations, location)
						starNumbers[key] = currentLocations
					}
				}
			}
		}
	}

	sum := 0

	for _, value := range starNumbers {
		if len(value) == 2 {
			sum += value[0].number * value[1].number
		}
	}

	return sum
}
