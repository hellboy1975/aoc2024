package day2

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"

	"github.com/pterm/pterm"
)

const day = 2
const title = "Red-Nosed Reports"

// this function makes a copy of the passed array, removes a value and returns it
func removeByIndex(s []int, index int) []int {

	cpy := make([]int, len(s))
	copy(cpy, s)
	return append(cpy[:index], cpy[index+1:]...)
}

// checks to see if the passed report is safe
func isReportSafe(report []int) bool {
	var direction string
	prev := 0

	for _, current := range report {
		if prev == 0 {
			// this is the first record - all we can do is record the value
			prev = current
		} else {
			// subtract current from prev to work out the diff
			diff := prev - current

			// if the diff is 0, or greater than +-3 then bomb out
			if diff == 0 || diff > 3 || diff < -3 {
				return false
			}

			// if direction hasn't been set
			if direction == "" {
				if diff > 0 {
					direction = "+"
				} else {
					direction = "-"
				}
			} else {
				// else if it has been set, if the diff is opposite, reject the record
				if (direction == "+" && diff < 0) || (direction == "-" && diff > 0) {
					return false
				}
			}

			prev = current
		}

	}

	// if the for loops gets all the way through, the report is safe
	return true
}

// if the dampeners are engaged, then we check reports - 1 array, returning when the first safe one is found
func dampenerCheck(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isReportSafe(removeByIndex(report, i)) {
			return true
		}
	}

	fmt.Println("Report  " + pterm.LightMagenta(fmt.Sprint(report)) + " UNSAFE")
	return false
}

func Part1() {
	var count int

	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)

	// Print a spacer line for better readability.
	pterm.Println()

	lines := base.GetDayDataLines(day)

	for _, s := range lines {
		report := base.StringToIntArray(s)
		if isReportSafe(report) {
			count++
		}
	}

	pterm.DefaultHeader.Println("Number of safe reports: " + fmt.Sprint(count))
}

func Part2() {
	count := 0
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)

	lines := base.GetDayDataLines(day)

	for _, s := range lines {
		report := base.StringToIntArray(s)
		if isReportSafe(report) {
			count++
		} else {
			if dampenerCheck(report) {
				count++
			}
		}
	}

	pterm.DefaultHeader.Println("Number of Reports: " + fmt.Sprint(len(lines)))
	pterm.DefaultHeader.Println("Number of " + pterm.LightMagenta("SAFE") + " Reports: " + fmt.Sprint(count))
}
