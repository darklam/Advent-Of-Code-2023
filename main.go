package main

import (
	"fmt"

	"advent-of-code-2023/day_3"
	"advent-of-code-2023/util"
)

func main() {
	result := day_3.GetGearRatiosSum(util.ReadFile)
	fmt.Println("Result = ", result)
}
