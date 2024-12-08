package day4

import (
	"reflect"
	"testing"
)

func init() {
}

func TestCalcDiagonal(t *testing.T) {

	cases := []struct {
		Description string
		Raw         []int
		Want        []cell
	}{
		{"Test col 0, 0", []int{0, 0}, []cell{{0, 0}}},
		{"Test col 5, 0", []int{5, 0}, []cell{{0, 5}, {1, 4}, {2, 3}, {3, 2}, {4, 1}, {5, 0}}},
		{"Test col 5, 2", []int{5, 2}, []cell{{2, 5}, {3, 4}, {4, 3}, {5, 2}, {6, 1}, {7, 0}}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := calcDiagonal(test.Raw[0], test.Raw[1])
			if reflect.DeepEqual(got, test.Want) != true {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}

func TestCountInLine(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        int
	}{
		{"Test forwards", "MMMSXXMASM", 1},
		{"Test backwards", "MSAMXMSMSA", 1},
		{"Test both overlap", "XMASAMXMMM", 2},
		{"Test both separate", "XMASMMSAMX", 2},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := countInLine(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
