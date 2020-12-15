package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/kprav33n/aoc20/day01"
	"github.com/kprav33n/aoc20/day02"
	"github.com/kprav33n/aoc20/day03"
	"github.com/kprav33n/aoc20/day04"
	"github.com/kprav33n/aoc20/day05"
	"github.com/kprav33n/aoc20/day06"
	"github.com/kprav33n/aoc20/functools"
	"github.com/kprav33n/aoc20/iotools"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dayNNx>\n", os.Args[0])
		os.Exit(1)
	}

	var r runner
	f := reflect.ValueOf(&r).MethodByName(strings.Title(os.Args[1]))
	if !f.IsValid() {
		fmt.Fprintf(os.Stderr, "Unknown puzzle: %s\n", os.Args[1])
		fmt.Fprintf(os.Stderr, "Solutions available for: ")
		t := reflect.TypeOf(&r)
		for i := 0; i < t.NumMethod(); i++ {
			if strings.HasPrefix(t.Method(i).Name, "Day") {
				fmt.Fprintf(os.Stderr, "%s ", strings.ToLower(t.Method(i).Name))
			}
		}
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(2)
	}

	f.Call(nil)
}

type runner struct{}

func (_ *runner) day01(product interface{}) {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		day01.ParseIntLines,
		product,
		fmt.Println,
	)("input/day01.txt")
}

func (r *runner) Day01a() {
	r.day01(day01.ERProduct)
}

func (r *runner) Day01b() {
	r.day01(day01.ERProduct3)
}

func (r *runner) day02(pred func(day02.Entry) bool) {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		day02.ParseLines,
		func(es []day02.Entry) int {
			return day02.Count(pred, es)
		},
		fmt.Println,
	)("input/day02.txt")
}

func (r *runner) Day02a() {
	r.day02(day02.IsValid)
}

func (r *runner) Day02b() {
	r.day02(day02.IsValid2)
}

func (_ *runner) day03(cf func([][]byte) int) {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		day03.ParseMap,
		functools.Unwrap,
		cf,
		fmt.Println,
	)("input/day03.txt")
}

func (r *runner) Day03a() {
	r.day03(func(m [][]byte) int {
		return day03.CountTrees(day03.Slope{Right: 3, Down: 1}, m)
	})
}

func (r *runner) Day03b() {
	r.day03(func(m [][]byte) int {
		return day03.CountTrees(day03.Slope{Right: 1, Down: 1}, m) *
			day03.CountTrees(day03.Slope{Right: 3, Down: 1}, m) *
			day03.CountTrees(day03.Slope{Right: 5, Down: 1}, m) *
			day03.CountTrees(day03.Slope{Right: 7, Down: 1}, m) *
			day03.CountTrees(day03.Slope{Right: 1, Down: 2}, m)
	})
}

func (_ *runner) Day04a() {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		day04.ParsePassports,
		functools.Unwrap,
		day04.CountValid,
		fmt.Println,
	)("input/day04.txt")
}

func (_ *runner) Day05a() {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		func(s string) []string {
			return strings.Split(strings.Trim(s, "\n"), "\n")
		},
		day05.ParsePasses,
		maxIntFromSlice,
		fmt.Println,
	)("input/day05.txt")
}

func (_ *runner) Day06a() {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		bytesToString,
		func(s string) []string {
			return strings.Split(strings.Trim(s, "\n"), "\n\n")
		},
		day06.ParseAnswers,
		sumIntInSlice,
		fmt.Println,
	)("input/day06.txt")
}

// Functions below this are shared between solutions. Extract them to a package
// if apropriate.

func bytesToString(b []byte) string {
	return string(b)
}

func maxIntFromSlice(xs []int) int {
	max := 0
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func sumIntInSlice(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
