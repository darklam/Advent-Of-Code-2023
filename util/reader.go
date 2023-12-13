package util

import (
	"bufio"
	"os"
)

type FileReader func(string) ([]string, error)

func ReadFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var result []string

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	return result, nil
}
