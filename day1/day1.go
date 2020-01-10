package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	masses := readInputFromFile("input.txt")
	result := Day1(masses...)
	fmt.Println("Result for Day1: ", result)
	result = Day1Part2(masses...)
	fmt.Println("Result for Day1Part2: ", result)
}

// Day1 receive a list of masses and returns the total fuel needed
func Day1(masses ...int) int {
	total := 0
	for _, val := range masses {
		total += (val/3 - 2)
	}
	return total
}

// Day1Part2 receive a list of masses and returns the total fuel needed accounting with the fuel's mass itself
func Day1Part2(masses ...int) int {
	total := 0
	for _, val := range masses {
		fuelNeeded := (val/3 - 2)
		for fuelNeeded > 0 {
			total += fuelNeeded
			fuelNeeded = (fuelNeeded/3 - 2)
		}
	}
	return total
}

// if an error occur, just panic!
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInputFromFile(filename string) []int {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	lines := strings.Split(string(dat), "\n")
	intlines := []int{}
	for _, val := range lines {
		if strings.Trim(val, " ") != "" {
			i, _ := strconv.Atoi(val)
			intlines = append(intlines, i)
		}
	}
	return intlines

}
