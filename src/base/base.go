package base

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const fullDataFile = "data/day_%d.txt"
const testDataFile = "data/day_%d_test.txt"

var UseTestData bool = false
var UseVisualisation bool = false

func SetTestData(x bool) {
	UseTestData = x
}

func SetVisualisation(x bool) {
	UseVisualisation = x
}

// combines the functionality of GetDayDataFile and reading the lines to a string array
func GetDayDataLines(day int) []string {
	file := GetDayDataFile(day)

	lines, err := ReadLines(file)
	if err != nil {
		panic(err)
	}

	return lines
}

// gets the days data as a string
func GetDayDataString(day int) string {
	b, err := os.ReadFile(GetDayDataFile(day))
	if err != nil {
		fmt.Print(err)
	}
	// b will be a []byte, so convert to a string
	return string(b)
}

// gets the data file for the requested day and part
func GetDayDataFile(day int) string {
	fn := fullDataFile
	if UseTestData {
		fn = testDataFile
	}
	file := fmt.Sprintf(fn, day)
	fmt.Println("  Data file: " + file)

	return file
}

func CountInArray(needle int, haystack []int) int {
	var count int
	for _, row := range haystack {
		if row == needle {
			count++
		}
	}
	return count
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Read a whole file into the memory and store it as array of lines
func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

// removes any duplicate integers from the passed slice
func RemoveDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// takes a string of space separated numbers, and converts this to an array of ints
func StringToIntArray(nums string) (arr []int) {
	nums = strings.TrimSpace(nums)              // remove trailing and leading whitespace
	nums = strings.ReplaceAll(nums, "   ", " ") // remove triple spaces
	return StringToIntArraySplit(nums, " ")
}

// same as StringToIntArray, but instead of doing some bullshit with spaces you just pass a delimiter
func StringToIntArraySplit(nums string, d string) (arr []int) {
	for _, num := range strings.Split(nums, d) {
		n, _ := strconv.Atoi(num)
		arr = append(arr, n)
	}
	return arr
}

func ReadFileTo2DByteArray(filePath string) ([][]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]byte
	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// script for transposing slice from (n x m) to (m x n) for golang.  From Gemini
func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// transposes a byte array.  Credit:  https://stackoverflow.com/questions/63257822/fast-transpose-byte-matrix-byte-in-golang-assembly
func TransposeByte(M [][]byte) [][]byte {
	m := len(M)
	n := len(M[0])

	// transposed matrix
	T := make([][]byte, n)
	for i := 0; i < n; i++ {
		T[i] = make([]byte, m)
	}

	var row []byte // a row in T
	for i := 0; i < n; i++ {
		row = T[i]
		for j := 0; j < m; j++ {
			row[j] = M[j][i]
		}
	}

	return T
}

// takes a 2D array, and rotates it 45 degrees
// for example, it will take a square, and make it a diamond
// 1 2 3
// 1 2 3
// 1 2 3
//
// to
//
//	 1
//	1 2
//
// 1 2 3
//
//	2 3
//	 3
func Rotate45Degrees(matrix [][]byte) [][]byte {
	n := len(matrix)
	newSize := 2*n - 1
	rotated := make([][]byte, newSize)
	for i := range rotated {
		rotated[i] = make([]byte, newSize)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newI := i + j
			newJ := j - i + n - 1
			rotated[newI][newJ] = matrix[i][j]
		}
	}

	return rotated
}

func Rotate90Degrees(matrix [][]byte) [][]byte {
	n := len(matrix)
	rotated := make([][]byte, n)
	for i := range rotated {
		rotated[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-1-i] = matrix[i][j]
		}
	}

	return rotated
}

func Rotate45CounterClockwise(matrix [][]byte) [][]byte {
	n := len(matrix)
	newSize := 2*n - 1
	rotated := make([][]byte, newSize)
	for i := range rotated {
		rotated[i] = make([]byte, newSize)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newI := i + j
			newJ := n - 1 - (i - j)
			rotated[newI][newJ] = matrix[i][j]
		}
	}

	return rotated
}

func RemoveZeros(arr []byte) []byte {
	// Initialize a counter to track non-zero elements
	count := 0

	// Iterate over the array, shifting non-zero elements to the beginning
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	// Create a new slice with the correct length (number of non-zero elements)
	result := make([]byte, count)
	copy(result, arr[:count])

	return result
}
