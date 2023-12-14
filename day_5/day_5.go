package day_5

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"advent-of-code-2023/util"
)

func getSeeds(line string) []int {
	numbersString := strings.Split(line, ":")[1][1:]

	numbers := strings.Split(numbersString, " ")
	result := make([]int, 0, len(numbers))

	for _, number := range numbers {
		converted, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		result = append(result, converted)
	}

	return result
}

func getNextType(currentType string, line string) (string, bool) {
	switch line {
	case "seed-to-soil map:":
		return "seed-to-soil", true
	case "soil-to-fertilizer map:":
		return "soil-to-fertilizer", true
	case "fertilizer-to-water map:":
		return "fertilizer-to-water", true
	case "water-to-light map:":
		return "water-to-light", true
	case "light-to-temperature map:":
		return "light-to-temperature", true
	case "temperature-to-humidity map:":
		return "temperature-to-humidity", true
	case "humidity-to-location map:":
		return "humidity-to-location", true
	case "":
		return "", true
	default:
		return currentType, false
	}
}

func getNumber(c uint8) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	} else {
		return -1
	}
}

func parseNumbers(line string) ConversionRange {
	nums := []int{0, 0, 0}

	numIndex := 0

	for i := 0; i < len(line); i++ {
		char := line[i]
		if char == ' ' {
			numIndex++
			continue
		}

		num := getNumber(char)
		if nums[numIndex] == 0 {
			nums[numIndex] = num
		} else {
			nums[numIndex] *= 10
			nums[numIndex] += num
		}
	}

	return ConversionRange{
		destinationRangeStart: nums[0],
		sourceRangeStart:      nums[1],
		offset:                nums[2],
	}
}

type ConversionRange struct {
	destinationRangeStart int
	sourceRangeStart      int
	offset                int
}

func convertValue(value int, r ConversionRange) int {
	difference := value - r.sourceRangeStart
	return r.destinationRangeStart + difference
}

func getConvertedValue(t *util.Tree[ConversionRange], value int) int {
	findFunc := func(seed int) func(ConversionRange) int {
		return func(data ConversionRange) int {
			if seed >= data.sourceRangeStart && seed < data.sourceRangeStart+data.offset {
				return 0
			} else if seed < data.sourceRangeStart {
				return -1
			} else {
				return 1
			}
		}
	}

	option := t.FindFunc(findFunc(value))
	if option.IsAbsent() {
		return value
	} else {
		return convertValue(value, option.MustGet())
	}
}

func GetMinLocation(reader util.FileReader) int {
	lines, err := reader("./inputs/day_5.txt")
	if err != nil {
		panic(err)
	}

	seeds := getSeeds(lines[0])

	compareFunc := func(a ConversionRange, b ConversionRange) int {
		return a.sourceRangeStart - b.sourceRangeStart
	}

	seedToSoilTree := util.NewTree[ConversionRange](compareFunc)
	soilToFertilizerTree := util.NewTree[ConversionRange](compareFunc)
	fertilizerToWaterTree := util.NewTree[ConversionRange](compareFunc)
	waterToLightTree := util.NewTree[ConversionRange](compareFunc)
	lightToTemperatureTree := util.NewTree[ConversionRange](compareFunc)
	temperatureToHumidityTree := util.NewTree[ConversionRange](compareFunc)
	humidityToLocationTree := util.NewTree[ConversionRange](compareFunc)

	lines = lines[1:]

	currentType := ""

	for _, line := range lines {
		var shouldSkip bool
		currentType, shouldSkip = getNextType(currentType, line)

		if shouldSkip {
			continue
		}

		data := parseNumbers(line)

		switch currentType {
		case "seed-to-soil":
			seedToSoilTree.Add(data)
		case "soil-to-fertilizer":
			soilToFertilizerTree.Add(data)
		case "fertilizer-to-water":
			fertilizerToWaterTree.Add(data)
		case "water-to-light":
			waterToLightTree.Add(data)
		case "light-to-temperature":
			lightToTemperatureTree.Add(data)
		case "temperature-to-humidity":
			temperatureToHumidityTree.Add(data)
		case "humidity-to-location":
			humidityToLocationTree.Add(data)
		}
	}

	locations := make([]int, 0, len(seeds))

	for _, seed := range seeds {
		soil := getConvertedValue(seedToSoilTree, seed)
		fertilizer := getConvertedValue(soilToFertilizerTree, soil)
		water := getConvertedValue(fertilizerToWaterTree, fertilizer)
		light := getConvertedValue(waterToLightTree, water)
		temperature := getConvertedValue(lightToTemperatureTree, light)
		humidity := getConvertedValue(temperatureToHumidityTree, temperature)
		location := getConvertedValue(humidityToLocationTree, humidity)

		locations = append(locations, location)

		fmt.Println("Seed = ", seed, " Soil = ", soil, " Fertilizer = ", fertilizer, "Location = ", location)
	}

	min := math.MaxInt

	for _, location := range locations {
		if location < min {
			min = location
		}
	}

	return min
}
