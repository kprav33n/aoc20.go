package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kprav33n/aoc20/day01"
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
	}
}
