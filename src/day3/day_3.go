package day3

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"regexp"

	"github.com/pterm/pterm"
)

const day = 3
const title = "Mull It Over"

// will split the mull into its two digits and multiply them
func processMul(mul string) int {
	m := base.StringToIntArraySplit(mul, ",")
	return m[0] * m[1]
}

func Part1() {

	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)
	mcount, multi := 0, 0
	code := base.GetDayDataString(day)
	re := regexp.MustCompile(`mul\(([0-9,]+)\)`)

	muls := re.FindAllStringSubmatch(code, -1)

	for _, mul := range muls {
		fmt.Println("  " + pterm.LightMagenta(fmt.Sprint(mcount)) + " raw: " + pterm.Green(fmt.Sprint(mul)))
		multi += processMul(mul[1])
		mcount++
	}
	fmt.Println("Multi count: " + pterm.Yellow(fmt.Sprint(multi)))
}

func Part2() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)
}
