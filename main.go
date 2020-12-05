package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kprav33n/aoc20/day01"
	"github.com/kprav33n/aoc20/day02"
	"github.com/kprav33n/aoc20/day03"
	"github.com/kprav33n/aoc20/day04"
)

func day01a() {
	data, err := ioutil.ReadFile("input/day01.txt")
	if err != nil {
		panic(err)
	}

	values := day01.ParseIntLines(string(data))
	product := day01.ERProduct(values)
	fmt.Println(product)

}

func day01b() {
	data, err := ioutil.ReadFile("input/day01.txt")
	if err != nil {
		panic(err)
	}

	values := day01.ParseIntLines(string(data))
	product := day01.ERProduct3(values)
	fmt.Println(product)

}

func day02a() {
	data, err := ioutil.ReadFile("input/day02.txt")
	if err != nil {
		panic(err)
	}

	entries := day02.ParseLines(string(data))
	count := day02.Count(day02.IsValid, entries)
	fmt.Println(count)
}

func day02b() {
	data, err := ioutil.ReadFile("input/day02.txt")
	if err != nil {
		panic(err)
	}

	entries := day02.ParseLines(string(data))
	count := day02.Count(day02.IsValid2, entries)
	fmt.Println(count)
}

func day03a() {
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

func day03b() {
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

func day04a() {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dayNNx>\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "day01a":
		day01a()

	case "day01b":
		day01b()

	case "day02a":
		day02a()

	case "day02b":
		day02b()

	case "day03a":
		day03a()

	case "day03b":
		day03b()

	case "day04a":
		day04a()
	}
}
