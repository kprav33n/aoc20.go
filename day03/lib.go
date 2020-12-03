package day03

import "strings"

func ParseMap(s string) ([][]byte, error) {
	var result [][]byte
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		result = append(result, []byte(line))
	}

	return result, nil
}

type Slope struct {
	Right int
	Down  int
}

func CountTrees(s Slope, m [][]byte) int {
	row, col := 0, 0
	numRows := len(m)
	numCols := len(m[0])
	numTrees := 0

	for row < numRows-1 {
		col = (col + s.Right) % numCols
		row += s.Down
		if m[row][col] == '#' {
			numTrees++
		}
	}

	return numTrees
}
