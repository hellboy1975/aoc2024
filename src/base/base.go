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
