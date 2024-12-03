package day2

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"

	"github.com/pterm/pterm"
)

const day = 2
const title = "Red-Nosed Reports"

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

func isReportSafe(report []int) bool {
	var count int
	pterm.DefaultBasicText.Println("Report" + pterm.LightMagenta(fmt.Sprint(report)))

	for _, level := range report {
		count += level
	}
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
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)

	// Print a spacer line for better readability.
	pterm.Println()
}
