package day01

import (
	"strconv"
	"strings"
)

func ParseIntLines(content string) []int {
	var result []int
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result = append(result, value)
	}
	return result
}

func ERProduct(input []int) int {
	for i := range input {
		for j := range input {
			if i != j && input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}

	return 0
}

func ERProduct3(input []int) int {
	for i := range input {
		for j := range input {
			for k := range input {
				if i != j && j != k && input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return 0
}
