package main

import "fmt"

// NextSmaller returns, for each element, the first strictly smaller element
// to its right, or -1 — the mirror of next-greater, same monotonic stack
// with the comparison flipped.
func NextSmaller(nums []int) []int {
	out := make([]int, len(nums))
	stack := []int{} // indexes still waiting; their values increase downwards
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > v {
			out[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, i := range stack {
		out[i] = -1
	}
	return out
}

func main() {
	fmt.Println(NextSmaller([]int{4, 8, 5, 2, 25}))
}
