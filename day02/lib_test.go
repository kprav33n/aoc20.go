package day02_test

import (
	"testing"

	"github.com/kprav33n/aoc20/day02"
	"github.com/stretchr/testify/require"
)

func TestParseEntry(t *testing.T) {
	assert := require.New(t)
	entry, err := day02.ParseEntry("1-3 a: abcde")
	assert.NoError(err)
	assert.Equal(1, entry.Min)
	assert.Equal(3, entry.Max)
	assert.Equal(byte('a'), entry.Char)
	assert.Equal([]byte("abcde"), entry.Password)
}

func TestIsValid(t *testing.T) {
	assert := require.New(t)

	entry := day02.Entry{
		Min:      1,
		Max:      3,
		Char:     'a',
		Password: []byte("abcde"),
	}
	assert.True(day02.IsValid(entry))

	entry = day02.Entry{
		Min:      1,
		Max:      3,
		Char:     'b',
		Password: []byte("cdefg"),
	}
	assert.False(day02.IsValid(entry))
}

func TestIsValid2(t *testing.T) {
	assert := require.New(t)

	entry := day02.Entry{
		Min:      1,
		Max:      3,
		Char:     'a',
		Password: []byte("abcde"),
	}
	assert.True(day02.IsValid2(entry))

	entry = day02.Entry{
		Min:      1,
		Max:      3,
		Char:     'b',
		Password: []byte("cdefg"),
	}
	assert.False(day02.IsValid2(entry))
}
