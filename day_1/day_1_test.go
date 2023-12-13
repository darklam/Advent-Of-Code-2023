package day_1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumCalibrationValues(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	reader := func(_ string) ([]string, error) {
		return lines, nil
	}

	result := SumCalibrationValues(reader)

	require.Equal(t, 142, result)
}

func TestSumCalibrationValuesWithSpelledDigits(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	reader := func(_ string) ([]string, error) {
		return lines, nil
	}

	result := SumCalibrationValues(reader)

	require.Equal(t, 281, result)
}
