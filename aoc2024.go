package main

import (
	// "aoc2024/src/day1"
	// "aoc2024/src/day2"
	// "aoc2024/src/day3"
	// "aoc2024/src/day4"
	"flag"
	"fmt"
	"hellboy1975/aoc2024/src/day1"
)

func main() {
	dayPtr := flag.Int("day", 1, "which day to run")
	partPtr := flag.Int("part", 1, "which part to run")

	flag.Parse()

	fmt.Println("Advent of Code 2024!")

	//TODO: come up with a better method than a giant if statement
	if *dayPtr == 1 {
		if *partPtr == 1 {
			day1.Part1()
		} else {
			day1.Part2()
		}
	}
	// else if *dayPtr == 2 {
	// 	if *partPtr == 1 {
	// 		day2.Part1()
	// 	} else {
	// 		day2.Part2()
	// 	}
	// } else if *dayPtr == 3 {
	// 	if *partPtr == 1 {
	// 		day3.Part1()
	// 	} else {
	// 		day3.Part2()
	// 	}
	// } else if *dayPtr == 4 {
	// 	if *partPtr == 1 {
	// 		day4.Part1()
	// 	} else {
	// 		day4.Part2()
	// 	}
	// }
}
