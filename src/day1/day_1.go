package day1

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"slices"
)

func Part1() {
	fmt.Println("Day 1, Part 1: Historian Hysteria")

	file := base.GetDayDataFile(1, 1, false)
	var left, right []int
	var count, sum, total int

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	// sort the data into left & right arrars
	for _, s := range lines {
		arr := base.StringToIntArray(s)

		left = append(left, arr[0])
		right = append(right, arr[1])
	}

	// sort each array low to high
	slices.Sort(left)
	slices.Sort(right)

	for _, l := range left {
		sum = base.Abs(right[count] - l)
		fmt.Println(fmt.Sprintf("%d | L: %d R: %d | %d", count, l, right[count], sum))

		total += sum
		count++
	}

	fmt.Println(fmt.Sprintf(" Total: %d", total))

}

func Part2() {
	fmt.Println("Day 1, Part 2: TBC!!!!")
}
