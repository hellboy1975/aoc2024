package day4

import (
	"testing"
)

func init() {
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
		{"Test short", "ASAMX", 1},
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
