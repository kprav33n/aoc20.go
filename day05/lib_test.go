package day05_test

import (
	"testing"

	"github.com/kprav33n/aoc20/day05"
	"github.com/stretchr/testify/require"
)

func TestParseEntry(t *testing.T) {
	assert := require.New(t)
	assert.Equal(567, day05.ParsePass("BFFFBBFRRR"))
	assert.Equal(119, day05.ParsePass("FFFBBBFRRR"))
	assert.Equal(820, day05.ParsePass("BBFFBBFRLL"))
}
