package day_3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

var input2 = []string{
	"467..114..",
	"...*...4*2",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestSumPartNumbers(t *testing.T) {
	reader := func(_ string) ([]string, error) {
		return input, nil
	}
	result := SumPartNumbers(reader)

	require.Equal(t, 4361, result)
}

func TestSumPartNumbers2(t *testing.T) {
	reader := func(_ string) ([]string, error) {
		return input2, nil
	}
	result := SumPartNumbers(reader)

	require.Equal(t, 5114, result)
}

func TestGetGearRatiosSum(t *testing.T) {
	reader := func(_ string) ([]string, error) {
		return input, nil
	}

	result := GetGearRatiosSum(reader)

	require.Equal(t, 467835, result)
}
