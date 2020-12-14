package main

import (
	"fmt"
	"io/ioutil"
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

type runner struct{}

func (_ *runner) day01f(product interface{}) {
	functools.Compose(
		iotools.ReadStdinOrFile,
		functools.Unwrap,
		func(b []byte) string {
			return string(b)
		},
		day01.ParseIntLines,
		product,
		fmt.Println,
	)("input/day01.txt")
}

func (r *runner) Day01a() {
	r.day01f(day01.ERProduct)
}

func (r *runner) Day01b() {
	r.day01f(day01.ERProduct3)
}

func (_ *runner) Day02a() {
	data, err := ioutil.ReadFile("input/day02.txt")
	if err != nil {
		panic(err)
	}

	entries := day02.ParseLines(string(data))
	count := day02.Count(day02.IsValid, entries)
	fmt.Println(count)
}

func (_ *runner) Day02b() {
	data, err := ioutil.ReadFile("input/day02.txt")
	if err != nil {
		panic(err)
	}

	entries := day02.ParseLines(string(data))
	count := day02.Count(day02.IsValid2, entries)
	fmt.Println(count)
}

func (_ *runner) Day03a() {
	data, err := ioutil.ReadFile("input/day03.txt")
	if err != nil {
		panic(err)
	}

	m, err := day03.ParseMap(string(data))
	if err != nil {
		panic(err)
	}

	count := day03.CountTrees(day03.Slope{3, 1}, m)
	fmt.Println(count)
}

func (_ *runner) Day03b() {
	data, err := ioutil.ReadFile("input/day03.txt")
	if err != nil {
		panic(err)
	}

	m, err := day03.ParseMap(string(data))
	if err != nil {
		panic(err)
	}

	count := day03.CountTrees(day03.Slope{1, 1}, m) *
		day03.CountTrees(day03.Slope{3, 1}, m) *
		day03.CountTrees(day03.Slope{5, 1}, m) *
		day03.CountTrees(day03.Slope{7, 1}, m) *
		day03.CountTrees(day03.Slope{1, 2}, m)
	fmt.Println(count)
}

func (_ *runner) Day04a() {
	data, err := ioutil.ReadFile("input/day04.txt")
	if err != nil {
		panic(err)
	}

	passports, err := day04.ParsePassports(string(data))
	if err != nil {
		panic(err)
	}

	count := day04.CountValid(passports)
	fmt.Println(count)
}

func (_ *runner) Day05a() {
	data, err := ioutil.ReadFile("input/day05.txt")
	if err != nil {
		panic(err)
	}

	passes := strings.Split(strings.Trim(string(data), "\n"), "\n")
	ids := day05.ParsePasses(passes)

	max := 0
	for _, i := range ids {
		if i > max {
			max = i
		}
	}

	fmt.Println(max)
}

func (_ *runner) Day06a() {
	data, err := ioutil.ReadFile("input/day06.txt")
	if err != nil {
		panic(err)
	}

	groups := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	counts := day06.ParseAnswers(groups)

	sum := 0
	for _, c := range counts {
		sum += c
	}

	fmt.Println(sum)
}

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
