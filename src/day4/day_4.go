package day4

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"slice"
	"strings"

	"github.com/pterm/pterm"
)

const day = 4
const title = "Ceres Search"

const xforward = "XMAS"
const xbackward = "SAMX"

var dirs = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

var green [][]int
var magenta [][]int
var yellow [][]int

var lines [][]byte
var area pterm.AreaPrinter

// determines how many instance of the backwards for forwards string occur in the line
func countInLine(line string) int {
	return strings.Count(line, xforward) + strings.Count(line, xbackward)
}

// iterates over the entire lines array that it's pointed to
func countInLines(lines *[][]byte) int {
	row, count := 0, 0
	for _, line := range *lines {
		l := string(line)
		fmt.Println("	" + fmt.Sprint(row) + ": " + pterm.Magenta(l) + " | " + fmt.Sprint(countInLine(string(l))))
		count += countInLine(l)
		row++
	}
	return count
}

// process the lines array by horizontal lines
func getHorizontal() int {
	count := countInLines(&lines)

	fmt.Println("H count: " + pterm.Green(fmt.Sprint(count)))
	return count
}

// process the lines array by vertical lines
func getVertical() int {
	// transpose to flip the array on it's side, and then process like a horizontal array
	lines = base.Rotate90Degrees(lines)
	count := countInLines(&lines)

	fmt.Println("V count: " + pterm.Green(fmt.Sprint(count)))
	return count
}

// a diagram used for plotting and scheming
// 00 01 02 03 04 05 06 07 08 09
// 10 11 12 13 14 15 16 17 18 19
// 20 21 22 23 24 25 26 27 28 29
// 30 31 32 33 34 35 36 37 38 39
// 40 41 42 43 44 45 46 47 48 49
// 50 51 52 53 54 55 56 57 58 59
// 60 61 62 63 64 65 66 67 68 69

func getDiagonalsRotated() int {
	r := base.Rotate45Degrees(lines)
	row, total := 0, 0

	for _, line := range r {
		line = base.RemoveZeros(line)
		c := countInLine(string(line))
		fmt.Println("	" + fmt.Sprint(row) + ": " + pterm.Magenta(string(line)) + " | " + fmt.Sprint(c))
		total += c
		row++
	}

	fmt.Println("R total: " + pterm.LightBlue(fmt.Sprint(total)))

	return total
}

func getDiagonalsRotatedL() int {
	lines := base.TransposeByte(lines)
	r := base.Rotate45CounterClockwise(lines)

	row, total := 0, 0

	for _, line := range r {
		line = base.RemoveZeros(line)
		c := countInLine(string(line))
		fmt.Println("	" + fmt.Sprint(row) + ": " + pterm.Magenta(string(line)) + " | " + fmt.Sprint(c))
		total += c
		row++
	}

	fmt.Println("R total: " + pterm.LightBlue(fmt.Sprint(total)))

	return total
}

// func checkDir(x, y, xDir, yDir int) bool {

// }

func checkWordFromCell(x, y int) bool {
	r := []int{-1, 0, 1}

	for _, yDir := range r {
		for _, xDir := range r {
			fmt.Println(string(xDir) + ", " + string(yDir))
		}
	}
	return false
}

// prints the martix that we're iterating through.  Makes X's red B)
func printMatrix() string {
	var output string
	x, y := 0, 0

	for _, l := range lines {
		for _, c := range l {
			if string(c) == "X" {
				output += pterm.Red(string(c))
			} else if slice.Index(green, []int{x, y}) {
				output += pterm.Green(string(c))
			} else {
				output += pterm.LightBlue(string(c))
			}
			x++
		}
		y++
		output += "\n"
	}

	return output
}

// scans the 8 cardinal directions from the passed reference to check if it spells XMAS
func scanDirs(x, y int) int {
	count, it := 0, -1

	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)
	primary.Printfln("Root: %d, %d", x, y)

	for _, dir := range dirs {
		// reset variables
		xDir := x
		yDir := y
		word := "X" // we've already found X, so we can seed the word check variable

		fmt.Println("dir	" + fmt.Sprint(dir))
		for i := 0; i < 3; i++ {
			// if the cell to be checked is outside of the bounds of the lines array, skip over this loop
			xDir += dir[0]
			yDir += dir[1]
			it++
			if (xDir < 0) || (xDir >= len(lines)) ||
				(yDir < 0) || (yDir >= len(lines)) {
				continue
			}
			check := string(lines[yDir][xDir])
			fmt.Println("	" + fmt.Sprint(it) + " [" + fmt.Sprintf("%d, %d", yDir, xDir) + "] " + ": " + pterm.Magenta(string(lines[yDir][xDir])))
			if i == 0 && check != "M" {
				break
			} else if i == 1 && check != "A" {
				break
			} else if i == 2 && check != "S" {
				break
			}
			word += check
			fmt.Println("	" + pterm.Red(word))

		}
		// once the loop breaks, see if it makes XMAS
		if word == "XMAS" {
			fmt.Println("	" + pterm.Green("XMAS!!"))
			count++
		}
	}
	return count
}

// scans the matrix for the occurance of a cell with the value X
func scanForX() int {
	//test 5, 5
	// count := scanDirs(4, 1)

	x, y, count := 0, 0, 0

	for _, l := range lines {
		for _, cell := range l {
			if string(cell) == "X" {
				count += scanDirs(x, y)

			}
			x++
		}
		y++
		x = 0
	}
	return count
}

func Part1() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)
	xmasCount := 0
	var err error

	lines, err = base.ReadFileTo2DByteArray(base.GetDayDataFile(day))

	fmt.Println("Matrix columns: " + pterm.LightBlue(len(lines[0])))
	fmt.Println("Matrix rows: " + pterm.LightBlue(len(lines)))

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print two new lines as spacer.
	pterm.Print("\n\n")

	// Start the Area printer from PTerm's DefaultArea, with the Center option.
	// The Area printer allows us to update a specific area of the console output.
	// The returned 'area' object is used to control the area updates.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Update the Area contents with the current time string.
	area.Update(printMatrix())

	area.Stop()

	xmasCount = scanForX()

	fmt.Println("Xmas count: " + pterm.Yellow(fmt.Sprint(xmasCount)))
}

func Part2() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)
}
