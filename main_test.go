package main

import (
	"reflect"
	"testing"
)

func TestNextSmaller(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{4, 8, 5, 2, 25}, []int{2, 5, 2, -1, -1}},
		{[]int{1, 2, 3}, []int{-1, -1, -1}},
		{[]int{3, 2, 1}, []int{2, 1, -1}},
		{[]int{2, 2}, []int{-1, -1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := NextSmaller(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("NextSmaller(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
