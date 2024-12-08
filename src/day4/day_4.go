package day4

import (
	"fmt"
	"hellboy1975/aoc2024/src/base"
	"strings"

	"github.com/pterm/pterm"
)

type cell struct {
	x int
	y int
}

const day = 4
const title = "Ceres Search"

const xforward = "XMAS"
const xbackward = "SAMX"

var lines [][]byte

// determines how many instance of the backwards for forwards string occur in the line
func countInLine(line string) int {
	return strings.Count(line, xforward) + strings.Count(line, xbackward)
}

// iterates over the entire lines array that it's pointed to
func countInLines(lines *[][]byte) int {
	row, count := 0, 0
	for _, line := range *lines {
		l := string(line)
		fmt.Println("	" + fmt.Sprint(row) + ": " + pterm.Magenta(l))
		count += countInLine(l)
		row++
	}
	return count
}

// process the lines array by horizontal lines
func getHorizontal() int {
	count := countInLines(&lines)

	fmt.Println("	H count: " + pterm.Green(fmt.Sprint(count)))
	return count
}

// process the lines array by vertical lines
func getVertical() int {
	// transpose to flip the array on it's side, and then process like a horizontal array
	lines = base.TransposeByte(lines)
	count := countInLines(&lines)

	fmt.Println("	V count: " + pterm.Green(fmt.Sprint(count)))
	return count
}

// func getStringFromCell(cells []cell) string {
// 	for _, c := range cells {
// 	}
// }

func calcDiagonal(col int, row int) []cell {

	if col == 0 && row == 0 {
		var r []cell
		r = append(r, cell{x: 0, y: 0})
		return r
	}

	cells := []cell{}
	off := 0

	for i := row; i <= col+row; i++ {
		c := cell{x: i, y: col - i + row}
		cells = append(cells, c)
		off++
	}

	fmt.Println("	D test: " + pterm.LightBlue(fmt.Sprint(cells)))

	return cells
}

// a diagram used for plotting and scheming
// 00 01 02 03 04 05 06 07 08 09
// 10 11 12 13 14 15 16 17 18 19
// 20 21 22 23 24 25 26 27 28 29
// 30 31 32 33 34 35 36 37 38 39
// 40 41 42 43 44 45 46 47 48 49
// 50 51 52 53 54 55 56 57 58 59
// 60 61 62 63 64 65 66 67 68 69

func getDiagonalsToRight() int {
	var cells [][]cell
	cols := len(lines[0])
	rows := len(lines) // I think it's a square, so rows probably = cols

	for i := 0; i <= cols; i++ {
		cells = append(cells, calcDiagonal(i, 0))
	}

	// I hate this method
	for r := 0; r <= rows; r++ {
		c := calcDiagonal(cols, r)
		if c.x <= cols {
			cells = append(cells, c)
		}

	}

	// fmt.Println("	D test: " + pterm.LightBlue(fmt.Sprint(cells)))

	return 5

}

func Part1() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 1: " + title)
	xmasCount := 0
	var err error

	lines, err = base.ReadFileTo2DByteArray(base.GetDayDataFile(day))

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// step 1 - find instances of `XMAS` and `SAMX` in the horizontal rows
	xmasCount += getHorizontal()

	// then transpose and do the same for vertical
	xmasCount += getVertical()

	xmasCount += getDiagonalsToRight()

	fmt.Println("Xmas count: " + pterm.Yellow(fmt.Sprint(xmasCount)))
}

func Part2() {
	pterm.DefaultHeader.Println("Day " + fmt.Sprint(day) + ", Part 2: " + title)
}
