package day2

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"

	"github.com/pterm/pterm"
)

const day = 2
const title = "Red-Nosed Reports"

var dampeners int

// func renderHeatmap(data [][]int) {
// 	var xAxis []string
// 	for i := 0; i < len(data); i++ {
// 		xAxis = append(xAxis, string(i))
// 	}

// 	// Define the labels for the X and Y axes of the heatmap.
// 	headerData := pterm.HeatmapAxis{
// 		XAxis: xAxis,
// 		YAxis: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
// 	}

// 	// Create a heatmap with the defined data and axis labels, and enable RGB colors.
// 	// Then render the heatmap.
// 	pterm.DefaultHeatmap.WithAxisData(headerData).WithData(data).WithEnableRGB().Render()
// }

func removeByIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func isReportSafe(report []int, dampener bool) bool {
	fmt.Println("Report" + pterm.LightMagenta(fmt.Sprint(report)))
	var direction string
	prev := 0

	var failState []bool

	for _, current := range report {
		fail := false
		if prev == 0 {
			// this is the first record - all we can do is record the value
			failState = append(failState, false) // failState will always be false
			prev = current
		} else {
			// subtract current from prev to work out the diff
			diff := prev - current

			// if the diff is 0, or greater than +-3 then bomb out
			if diff == 0 || diff > 3 || diff < -3 {
				fail = true
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
					fail = true
				}
			}

			failState = append(failState, fail)

			prev = current
		}

	}

	fmt.Println("FailState:" + pterm.Red(fmt.Sprint(failState)))

	// how many failures were there?
	failCount, i, failIndex := 0, 0, 0
	for _, value := range failState {
		if value {
			failCount++
			failIndex = i
		}
		i++
	}

	// failure states:
	//  - failCount is 2 or more
	//  - dampener is false and failCount > 0
	if failCount > 1 || (!dampener && failCount > 0) {
		return false
	}

	if failCount == 1 {
		// remove the failed value from the report and try again
		fmt.Println("Iterate")
		dampeners++
		report = removeByIndex(report, failIndex)
		return isReportSafe(report, false)
	}

	// if the for loops gets all the way through, the report is safe
	return true
}

func Part1() {
	var count int

	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)

	// Print a spacer line for better readability.
	pterm.Println()

	lines := base.GetDayDataLines(day)

	for _, s := range lines {
		report := base.StringToIntArray(s)
		if isReportSafe(report, false) {
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
		if isReportSafe(report, true) {
			fmt.Println("Report" + pterm.LightMagenta(fmt.Sprint(report)) + " SAFE")
			// pterm.DefaultBasicText.Println("Report" + pterm.LightMagenta(fmt.Sprint(report)) + " SAFE")
			count++
		} else {
			fmt.Println("Report" + pterm.LightMagenta(fmt.Sprint(report)) + " UNSAFE")
		}

		fmt.Println()
	}

	pterm.DefaultHeader.Println("Number of Reports: " + fmt.Sprint(len(lines)))
	pterm.DefaultHeader.Println("Number of Dampener engagements: " + fmt.Sprint(dampeners))
	pterm.DefaultHeader.Println("Number of " + pterm.LightMagenta("SAFE") + " Reports: " + fmt.Sprint(count))
}
