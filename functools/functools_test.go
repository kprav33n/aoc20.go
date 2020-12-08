package functools_test

import (
	"errors"
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

func TestUnaryUnwrapError(t *testing.T) {
	require := require.New(t)

	f := func() (int, error) {
		return 14, errors.New("forced error")
	}
	require.Panics(func() { functools.Unwrap(f()) })
}

func TestUnaryUnwrapOK(t *testing.T) {
	require := require.New(t)

	f := func() (int, error) {
		return 14, nil
	}
	result := functools.Unwrap(f())
	require.Len(result, 1)
	require.Equal(result[0], 14)
}

func TestBinaryUnwrapError(t *testing.T) {
	require := require.New(t)

	f := func() (string, int, error) {
		return "emc", 2, errors.New("forced error")
	}
	require.Panics(func() { functools.Unwrap(f()) })
}

func TestBinaryUnwrapOK(t *testing.T) {
	require := require.New(t)

	f := func() (string, int, error) {
		return "emc", 2, nil
	}
	result := functools.Unwrap(f())
	require.Len(result, 2)
	require.Equal(result[0], "emc")
	require.Equal(result[1], 2)
}

func TestUnwrapComposeError(t *testing.T) {
	require := require.New(t)

	firstFunc := func() (string, error) {
		return "", errors.New("forced error")
	}

	secondFunc := func(s string) {
	}

	require.PanicsWithError("forced error", func() {
		functools.Compose(firstFunc, functools.Unwrap, secondFunc)()
	})
}

func TestUnwrapComposeOK(t *testing.T) {
	require := require.New(t)

	firstFunc := func() (string, error) {
		return "hello", nil
	}

	var arg string
	secondFunc := func(s string) {
		arg = s
	}

	functools.Compose(firstFunc, functools.Unwrap, secondFunc)()
	require.Equal("hello", arg)
}
