package main

import (
	"flag"
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"hellboy1975/aoc2024/src/day1"
	"hellboy1975/aoc2024/src/day2"
	"hellboy1975/aoc2024/src/day3"
	"hellboy1975/aoc2024/src/day4"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	dayPtr := flag.Int("day", 1, "which day to run")
	partPtr := flag.Int("part", 1, "which part to run")
	testPtr := flag.Bool("test", false, "use the test data instead of the actual data")
	visualisePtr := flag.Bool("visualise", false, "show any visualisations for the solution")

	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("AOC", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("2024", pterm.FgLightMagenta.ToStyle())).
		Render() // Render the big text to the terminal

	flag.Parse()

	if *testPtr {
		fmt.Println("  Test mode " + pterm.Green("ON"))
		base.SetTestData(true)
	}

	if *visualisePtr {
		fmt.Println("  Visualiser ON")
		base.SetVisualisation(true)
	}

	//TODO: come up with a better method than a giant if statement
	if *dayPtr == 1 {
		if *partPtr == 1 {
			day1.Part1()
		} else {
			day1.Part2()
		}
	} else if *dayPtr == 2 {
		if *partPtr == 1 {
			day2.Part1()
		} else {
			day2.Part2()
		}
	} else if *dayPtr == 3 {
		if *partPtr == 1 {
			day3.Part1()
		} else {
			day3.Part2()
		}
	} else if *dayPtr == 4 {
		if *partPtr == 1 {
			day4.Part1()
		} else {
			day4.Part2()
		}
	} else if *dayPtr == 5 {
		if *partPtr == 1 {
			day5.Part1()
		} else {
			day5.Part2()
		}
	}
}
