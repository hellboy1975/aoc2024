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

const testDataFile = "data/day_%d_%d.txt"

// gets the data file for the requested day and part
func GetDayDataFile(day, part int) string {
	file := fmt.Sprintf(testDataFile, day, part)
	fmt.Println("  Data file: " + file)

	return file
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
	nums = strings.TrimSpace(nums)             // remove trailing and leading whitespace
	nums = strings.ReplaceAll(nums, "  ", " ") // remove double spaces
	for _, num := range strings.Split(nums, " ") {
		n, _ := strconv.Atoi(num)
		arr = append(arr, n)
	}
	return arr
}
