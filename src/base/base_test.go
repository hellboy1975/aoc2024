package base

import (
	"reflect"
	"testing"
)

func init() {
}

func TestRemoveDuplicateInt(t *testing.T) {

	cases := []struct {
		Description string
		Raw         []int
		Want        []int
	}{
		{"1, 1, 1, 1 should be 1", []int{1, 1, 1, 1}, []int{1}},
		{"1, 2, 3, 3, 3 should be 1, 2, 3", []int{1, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"54, 500, 277, 129, 277, 1000, 1, 1000 should be 54, 500, 277, 129, 1000, 1", []int{54, 500, 277, 129, 277, 1000, 1, 1000}, []int{54, 500, 277, 129, 1000, 1}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := RemoveDuplicateInt(test.Raw)
			if reflect.DeepEqual(got, test.Want) != true {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}

func TestStringToIntArray(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        []int
	}{
		{"Basic test", "1 2 3 4", []int{1, 2, 3, 4}},
		{"Multi digit test", "11 24 35 4111", []int{11, 24, 35, 4111}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := StringToIntArray(test.Raw)
			if reflect.DeepEqual(got, test.Want) != true {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
