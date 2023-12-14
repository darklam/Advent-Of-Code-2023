package main

import (
	"fmt"

	"advent-of-code-2023/day_5"
	"advent-of-code-2023/util"
)

func main() {
	result := day_5.GetMinLocation(util.ReadFile)
	fmt.Println("Result = ", result)
}
