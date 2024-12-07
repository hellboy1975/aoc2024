package day3

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"regexp"
	"strings"

	"github.com/pterm/pterm"
)

const day = 3
const title = "Mull It Over"

// will split the mull into its two digits and multiply them
func processMul(mul string) int {
	m := base.StringToIntArraySplit(mul, ",")
	return m[0] * m[1]
}

func getMulsFromCode(code string) int {
	re := regexp.MustCompile(`mul\(([0-9,]+)\)`)
	multi := 0

	muls := re.FindAllStringSubmatch(code, -1)

	for _, mul := range muls {
		fmt.Println(" raw: " + pterm.Green(fmt.Sprint(mul)))
		multi += processMul(mul[1])
	}

	return multi
}

func Part1() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)

	code := base.GetDayDataString(day)
	multi := getMulsFromCode(code)

	fmt.Println("Multi count: " + pterm.Yellow(fmt.Sprint(multi)))
}

func Part2() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)

	code := base.GetDayDataString(day)
	multi, n := 0, 0

	// can we try removing any strings that are between a don't() and do() marker?
	for _, mul := range strings.Split(code, "don't()") {
		if n == 0 || n%2 == 0 {
			slab := ""
			fmt.Println(" raw slab: " + fmt.Sprint(n) + " | " + pterm.LightMagenta(fmt.Sprint(mul)))
			if n == 0 {
				slab = mul
			} else {
				s := strings.SplitAfterN(mul, "do()", 2)
				if len(s) > 1 {
					slab = s[1]
				}
			}
			fmt.Println(" act slab: " + fmt.Sprint(n) + " | " + pterm.Yellow(fmt.Sprint(slab)))
			multi += getMulsFromCode(slab) // in theory it should be index 1

		}
		n++
	}

	fmt.Println("Multi count: " + pterm.Yellow(fmt.Sprint(multi)))
}
