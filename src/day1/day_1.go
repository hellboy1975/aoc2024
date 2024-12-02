package day1

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"slices"

	"github.com/pterm/pterm"
)

func Part1() {
	fmt.Println("Day 1, Part 1: Historian Hysteria")

	file := base.GetDayDataFile(1, 1)
	multi := pterm.DefaultMultiPrinter
	var left, right []int
	var count, sum, total int

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	// sort the data into left & right arrays
	for _, s := range lines {
		arr := base.StringToIntArray(s)

		left = append(left, arr[0])
		right = append(right, arr[1])
	}

	// sort each array low to high
	slices.Sort(left)
	slices.Sort(right)

	multi.Start()
	for _, l := range left {
		sum = base.Abs(right[count] - l)

		if base.UseVisualisation {
			pb1, _ := pterm.DefaultProgressbar.WithTotal(5000).WithWriter(multi.NewWriter()).Start(fmt.Sprintf("Line %d", count))
			pb1.Add(sum)
		}

		total += sum
		count++
	}
	multi.Stop()

	fmt.Println(fmt.Sprintf(" Total: %d", total))

}

func Part2() {
	fmt.Println("Day 1, Part 2: Historian Hysteria")

	file := base.GetDayDataFile(1, 1)
	var left, right []int
	var count, total int

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	// sort the data into left & right arrays
	for _, s := range lines {
		arr := base.StringToIntArray(s)

		left = append(left, arr[0])
		right = append(right, arr[1])
	}

	for _, l := range left {
		total += l * base.CountInArray(l, right)
		count++
	}

	fmt.Println(fmt.Sprintf(" Total: %d", total))
}
