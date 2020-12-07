package functools_test

import (
	"testing"

	"github.com/kprav33n/aoc20/functools"
	"github.com/stretchr/testify/require"
)

func TestUnaryCompose1(t *testing.T) {
	require := require.New(t)

	var arg string
	unaryFunc := func(s string) {
		arg = s
	}

	functools.Compose(unaryFunc)("hello")
	require.Equal("hello", arg)
}

func TestBinaryCompose1(t *testing.T) {
	require := require.New(t)

	var arg1 string
	var arg2 int

	binaryFunc := func(s string, i int) {
		arg1 = s
		arg2 = i
	}

	functools.Compose(binaryFunc)("area", 51)
	require.Equal("area", arg1)
	require.Equal(51, arg2)
}

func TestCompose2(t *testing.T) {
	require := require.New(t)

	var arg1 string
	var arg2 int
	var arg3 float32

	firstFunc := func(s string, i int) float32 {
		arg1 = s
		arg2 = i
		return 3.1416
	}

	secondFunc := func(f float32) {
		arg3 = f
	}

	functools.Compose(firstFunc, secondFunc)("area", 51)
	require.Equal("area", arg1)
	require.Equal(51, arg2)
	require.InDelta(3.1416, arg3, 0.0001)
}
